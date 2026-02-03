package broker

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadBrokerConfig(path string) (*BrokerConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("cannot read config file: %w", err)
	}

	var cfg BrokerConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("cannot parse yaml: %w", err)
	}

	
	if cfg.Broker.Service.Name == "" {
		return nil, fmt.Errorf("broker.service.name is required")
	}
	if cfg.Broker.Service.DataDir == "" {
		return nil, fmt.Errorf("broker.service.dataDir is required")
	}
	if cfg.Broker.Service.Host == "" {
		return nil, fmt.Errorf("broker.service.host is required")
	}
	if cfg.Broker.Service.Port == 0 {
		return nil, fmt.Errorf("broker.service.port is required")
	}
	if cfg.Broker.Service.BrokerID < 0 {
		return nil, fmt.Errorf("broker.service.brokerId must be >= 0")
	}

	return &cfg, nil
}
