package config

import (
	"errors"

	corecfg "github.com/Axway/agent-sdk/pkg/config"
)

// AgentConfig - represents the config for agent
type AgentConfig struct {
	CentralCfg corecfg.CentralConfig `config:"central"`
	GatewayCfg *GatewayConfig        `config:"gateway-section"`
}

// GatewayConfig - represents the config for gateway
type GatewayConfig struct {
	corecfg.IConfigValidator
	LogFile        string `config:"logFile"`
	ProcessOnInput bool   `config:"processOnInput"`
	RestPort       string `config:"restPort"`
	ConfigKey1     string `config:"config_key_1"`
	ConfigKey2     string `config:"config_key_2"`
	ConfigKey3     string `config:"config_key_3"`
}

// ValidateCfg - Validates the gateway config
func (c *GatewayConfig) ValidateCfg() (err error) {
	if c.LogFile == "" {
		return errors.New("Invalid gateway configuration: logFile is not configured")
	}
	if c.RestPort == "" {
		return errors.New("Invalid gateway configuration: restPort is not configured")
	}

	return
}
