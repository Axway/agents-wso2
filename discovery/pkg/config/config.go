package config

import (
	"errors"

	corecfg "github.com/Axway/agent-sdk/pkg/config"
)

// AgentConfig - represents the config for agent
type AgentConfig struct {
	CentralCfg corecfg.CentralConfig `config:"central"`
	GatewayCfg *WSO2Config           `config:"wso2"`
}

// WSO2Config - represents the config for WSO2gateway
type WSO2Config struct {
	corecfg.IConfigValidator
	Basepath      string `config:"basepath"`
	TokenEndpoint string `config:"tokenEndpoint"`
	Username      string `config:"username"`
	Password      string `config:"password"`
	ClientSecret  string `config:"clientSecret"`
	ClientID      string `config:"clientId"`
	Tag           string `config:"tag"`
	Scope         string `config:"scope"`
}

// ValidateCfg - Validates the gateway config
func (c *WSO2Config) ValidateCfg() (err error) {
	if c.Basepath == "" || c.TokenEndpoint == "" {
		return errors.New("Invalid ws02 configuration: is not configured")
	}

	return
}
