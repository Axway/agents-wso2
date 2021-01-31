package cmd

import (
	corecmd "github.com/Axway/agent-sdk/pkg/cmd"
	corecfg "github.com/Axway/agent-sdk/pkg/config"
	"github.com/Axway/agent-sdk/pkg/util/log"

	// CHANGE_HERE - Change the import path(s) below to reference packages correctly
	"github.com/Axway/agents-wso2/discovery/pkg/config"
	"github.com/Axway/agents-wso2/discovery/pkg/wso2"
)

// RootCmd - Agent root command
var RootCmd corecmd.AgentRootCmd
var agentConfig *config.AgentConfig

func init() {
	// Create new root command with callbacks to initialize the agent config and command execution.
	// The first parameter identifies the name of the yaml file that agent will look for to load the config
	RootCmd = corecmd.NewRootCmd(
		"apic_discovery_agent",        // Name of the yaml file
		"WSO2 Sample Discovery Agent", // Agent description
		initConfig,                    // Callback for initializing the agent config
		run,                           // Callback for executing the agent
		corecfg.DiscoveryAgent,        // Agent Type (Discovery or Traceability)
	)

	// Get the root command properties and bind the config property in YAML definition
	rootProps := RootCmd.GetProperties()
	rootProps.AddStringProperty("wso2.basepath", "https://localhost:9443/api/am/publisher/v1", "WS02 API Manger basepath")
	rootProps.AddStringProperty("wso2.tokenEndpoint", "", "")
	rootProps.AddStringProperty("wso2.username", "", "")
	rootProps.AddStringProperty("wso2.password", "", "")
	rootProps.AddStringProperty("wso2.clientId", "", "")
	rootProps.AddStringProperty("wso2.clientSecret", "", "")
	rootProps.AddStringProperty("wso2.tag", "ampc", "")
	rootProps.AddStringProperty("wso2.scope", "", "")

}

// Callback that agent will call to process the execution
func run() error {

	wso2Client, err := wso2.NewGatewayClient(agentConfig)
	if err != nil {
		log.Info(err.Error())
		return err
	}
	err = wso2Client.DiscoverAPIs()
	return err
}

// Callback that agent will call to initialize the config. CentralConfig is parsed by Agent SDK
// and passed to the callback allowing the agent code to access the central config
func initConfig(centralConfig corecfg.CentralConfig) (interface{}, error) {
	rootProps := RootCmd.GetProperties()
	// Parse the config from bound properties and setup gateway config
	wso2Config := &config.WSO2Config{
		Basepath:      rootProps.StringPropertyValue("wso2.basepath"),
		TokenEndpoint: rootProps.StringPropertyValue("wso2.tokenEndpoint"),
		Username:      rootProps.StringPropertyValue("wso2.username"),
		Password:      rootProps.StringPropertyValue("wso2.password"),
		ClientID:      rootProps.StringPropertyValue("wso2.clientId"),
		ClientSecret:  rootProps.StringPropertyValue("wso2.clientSecret"),
		Tag:           rootProps.StringPropertyValue("wso2.tag"),
		Scope:         rootProps.StringPropertyValue("wso2.scope"),
	}

	agentConfig = &config.AgentConfig{
		CentralCfg: centralConfig,
		GatewayCfg: wso2Config,
	}
	return agentConfig, nil
}

// GetAgentConfig - Returns the agent config
func GetAgentConfig() *config.AgentConfig {
	return agentConfig
}
