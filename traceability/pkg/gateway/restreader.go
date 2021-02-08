package gateway

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/cast"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	// CHANGE_HERE - Change the import path(s) below to reference packages correctly
	"github.com/Axway/agents-wso2/traceability/pkg/config"
)

// RestReader - Represents the Gateway client
type RestReader struct {
	cfg          *config.GatewayConfig
	eventChannel chan string
}

// NewRestReader - Creates a new Gateway Client
func NewRestReader(gatewayCfg *config.GatewayConfig, eventChannel chan string) (*RestReader, error) {
	return &RestReader{
		cfg:          gatewayCfg,
		eventChannel: eventChannel,
	}, nil
}

// Start - Starts reading log file
func (r *RestReader) Start() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/trace", r.getTrace).Methods("POST")
	log.Fatal(http.ListenAndServe(":"+r.cfg.RestPort, myRouter))
}

func (r RestReader) getTrace(w http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		// Try again manually
	}

	body, err = fixBody(body)

	r.eventChannel <- string(body)
}

func fixBody(body []byte) ([]byte, error) {
	var jsonBody map[string]interface{}

	err := json.Unmarshal(body, &jsonBody)
	if err != nil {
		return nil, errors.New("Failed to unmarshal string " + string(body) + " into json: " + err.Error())
	}

	inBound := fixDirection(cast.ToStringMap(jsonBody["inbound"]))
	outBound := fixDirection(cast.ToStringMap(jsonBody["outbound"]))

	jsonBody["inbound"] = inBound
	jsonBody["outbound"] = outBound

	return json.Marshal(jsonBody)
}

// WSO2 returns some string values that we want to convert to integer values
func fixDirection(stringMap map[string]interface{}) map[string]interface{} {
	returnMap := stringMap
	for key, value := range returnMap {
		if key == "srcPort" || key == "destPort" || key == "statusCode" {
			stringValue := fmt.Sprintf("%v", value)
			// If we get the reminder of the url remove it so that we only have the port number
			if strings.Contains(stringValue, "/") {
				stringValue = strings.Split(stringValue, "/")[0]
			}

			// If we fail to convert to int we will return 0
			intValue, _ := strconv.Atoi(stringValue)
			var intInterface interface{} = intValue
			returnMap[key] = intInterface
		}
	}

	return returnMap
}
