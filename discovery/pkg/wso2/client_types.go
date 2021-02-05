package wso2

import (
	"github.com/Axway/agent-sdk/pkg/api"
	"github.com/Axway/agents-wso2/discovery/pkg/config"
	"github.com/Axway/agents-wso2/discovery/pkg/wso2/models"
)

// AuthResponse - Authentication response data
type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
}

// GatewayClient - client connection to WSO2
type GatewayClient struct {
	config     *config.AgentConfig
	httpClient api.Client
	authData   *AuthResponse
}

type Wso2API struct {
	models.API
	swaggerSpec     []byte
	swaggerSpecType string
}
