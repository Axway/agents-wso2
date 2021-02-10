package wso2

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/Axway/agent-sdk/pkg/agent"
	"github.com/Axway/agent-sdk/pkg/api"
	"github.com/Axway/agent-sdk/pkg/apic"
	corecfg "github.com/Axway/agent-sdk/pkg/config"
	log "github.com/Axway/agent-sdk/pkg/util/log"
	"github.com/Axway/agents-wso2/discovery/pkg/config"
	"github.com/Axway/agents-wso2/discovery/pkg/wso2/models"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

// NewGatewayClient - builds a new Client using the AgentConfig
func NewGatewayClient(cfg *config.AgentConfig) (*GatewayClient, error) {

	tlsConfig := &corecfg.TLSConfiguration{
		InsecureSkipVerify: true,
		NextProtos:         []string{},
		CipherSuites:       corecfg.TLSDefaultCipherSuites,
		MinVersion:         corecfg.TLSDefaultMinVersion,
		MaxVersion:         0,
	}

	c := &GatewayClient{
		config:     cfg,
		httpClient: api.NewClient(tlsConfig, ""),
	}

	err := c.Authenticate()

	if err != nil {
		return nil, err
	}

	return c, nil
}

// DiscoverAPIs - Discover and publish APIs to AMPLIFY Central
func (c *GatewayClient) DiscoverAPIs() error {
	searchResults, err := c.findAmplifyAPIs()

	for _, api := range *searchResults.List {
		log.Debugf("found: %s", *api.Id)
		apiDetails, err := c.getAPIDetails(*api.Id)
		if err != nil {
			log.Error("Failed to get API details for " + *api.Id)
			continue
		}

		serviceBody, err := apiDetails.buildServiceBody()
		if err != nil {
			log.Error("Failed to get service body for " + apiDetails.Name)
			continue
		}

		err = agent.PublishAPI(serviceBody)
		if err != nil {
			log.Error("Failed to publish api " + apiDetails.Name)
			continue
		}
		log.Infof("Published API '%s' to AMPLIFY Central", serviceBody.NameToPush)
	}

	return err
}

// Authenticate - Authenticate with gateway
func (c *GatewayClient) Authenticate() error {
	authValues := url.Values{}
	authValues.Set("grant_type", "password")
	authValues.Set("username", c.config.GatewayCfg.Username)
	authValues.Set("password", c.config.GatewayCfg.Password)
	authValues.Set("scope", c.config.GatewayCfg.Scope)

	creds := fmt.Sprintf("%s:%s", c.config.GatewayCfg.ClientID, c.config.GatewayCfg.ClientSecret)

	request := api.Request{
		Method: api.POST,
		URL:    c.config.GatewayCfg.TokenEndpoint,

		Headers: map[string]string{
			"Content-Type":  "application/x-www-form-urlencoded",
			"Authorization": fmt.Sprintf("Basic %s", b64.StdEncoding.EncodeToString([]byte(creds))),
		},
		Body: []byte(authValues.Encode()),
	}

	response, err := c.httpClient.Send(request)
	if err != nil {
		log.Info(err.Error())
		return err
	}

	log.Debugf("Status: %s", strconv.Itoa(response.Code))
	log.Debugf("Body: %s", string(response.Body))

	c.authData = &AuthResponse{}
	json.Unmarshal(response.Body, c.authData)
	return nil
}

func (c *GatewayClient) callAPI(endpoint string, method string, queryParams map[string]string, headers map[string]string) (*api.Response, error) {

	if headers == nil {
		headers = make(map[string]string)
	}

	headers["Authorization"] = fmt.Sprintf("Bearer %s", c.authData.AccessToken)

	request := api.Request{
		Method:      method,
		URL:         c.config.GatewayCfg.Basepath + endpoint,
		QueryParams: queryParams,
		Headers:     headers,
	}

	response, err := c.httpClient.Send(request)

	if err != nil {
		log.Info(err.Error())
		return nil, err
	}

	log.Debugf("Status (%s) : %s", endpoint, strconv.Itoa(response.Code))
	log.Debugf("Response body: %s", string(response.Body))

	return response, nil
}

func (c *GatewayClient) findAmplifyAPIs() (*models.APIList, error) {
	queryParmas := map[string]string{
		"query": "tag:" + c.config.GatewayCfg.Tag,
	}
	resp, err := c.callAPI("/apis", api.GET, queryParmas, nil)

	if err != nil {
		return nil, err
	}
	r := &models.APIList{}
	json.Unmarshal(resp.Body, &r)

	log.Infof("Found %d API to be published to AMPLIFY Central", len(*r.List))
	return r, nil
}

func (c *GatewayClient) getAPIDetails(apiID string) (*Wso2API, error) {
	resp, err := c.callAPI(fmt.Sprintf("/apis/%s", apiID), api.GET, nil, nil)

	if err != nil {
		return nil, err
	}

	api := &Wso2API{}
	json.Unmarshal(resp.Body, &api)

	spec, err := c.getAPISpec(*api.Id)

	if err != nil {
		return nil, err
	}

	var jsonSpecMap map[string]interface{}

	err = json.Unmarshal(spec, &jsonSpecMap)
	if jsonSpecMap["swagger"] != nil {
		jsonSpecMap["basePath"] = *&api.Context + "/" + *&api.Version
		// add to config params
		jsonSpecMap["host"] = "gateway.api.cloud.wso2.com"
	} else if jsonSpecMap["openapi"] != nil {
		jsonSpecMap["servers"] = [1]map[string]interface{}{{"url": "gateway.api.cloud.wso2.com" + *&api.Context + "/" + *&api.Version}}
	}

	b, err := json.Marshal(jsonSpecMap)

	api.setSwaggerSpec(b)

	return api, nil
}

func (c *GatewayClient) getAPISpec(apiID string) ([]byte, error) {
	resp, err := c.callAPI(fmt.Sprintf("/apis/%s/swagger", apiID), api.GET, nil, nil)

	if err != nil {
		return nil, err
	}

	spec := resp.Body

	if log.GetLevel() == logrus.DebugLevel {
		log.Debugf("Swagger JSON : \n%s", string(spec))
	}

	return spec, nil
}

func (api *Wso2API) buildServiceBody() (apic.ServiceBody, error) {

	return apic.NewServiceBodyBuilder().
		SetID(*api.Id).
		SetTitle(api.Name).
		SetURL(api.getURL()).
		SetDescription(api.getDescription()).
		SetAPISpec(api.swaggerSpec).
		SetVersion(api.Version).
		SetAuthPolicy(api.getAuthPolicy()).
		SetDocumentation(api.getDocumentation()).
		SetResourceType(apic.Oas2).
		Build()
}

func (api *Wso2API) getURL() string {
	// update
	return ""
}

func (api *Wso2API) getAuthPolicy() string {
	// WS02 Cloud returns only "apiSecurity"
	if api.SecurityScheme == nil {
		return apic.Oauth
	}

	for _, scheme := range *api.SecurityScheme {
		switch scheme {
		case "oauth2":
			return apic.Oauth
		case "api_key":
			return apic.Apikey
		}
	}

	return apic.Passthrough
}

func (api *Wso2API) getDescription() string {
	return fmt.Sprintf("\"%s\"", *api.Description)
}

func (api *Wso2API) getDocumentation() []byte {
	return []byte(fmt.Sprintf("\"%s\"", *api.Description))
}

func (api *Wso2API) setSwaggerSpec(spec []byte) {
	api.swaggerSpec = spec

	if val := gjson.Get(string(spec), "swagger"); val.Exists() {
		api.swaggerSpecType = apic.Oas2
	} else {
		api.swaggerSpecType = apic.Oas3
	}

}
