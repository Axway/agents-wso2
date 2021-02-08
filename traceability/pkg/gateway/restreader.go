package gateway

import (
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	// CHANGE_HERE - Change the import path(s) below to reference packages correctly
	log "github.com/Axway/agent-sdk/pkg/util/log"
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
	log.Info(http.ListenAndServe(":"+r.cfg.RestPort, myRouter))
}

func (r RestReader) getTrace(w http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	log.Debugf("trace message: %s", string(body))

	r.eventChannel <- string(body)
}
