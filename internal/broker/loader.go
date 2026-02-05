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
		cfg.Broker.Service.Host = "localhost"
	}


	if cfg.Broker.Service.Port == 0 {
		cfg.Broker.Service.Port = 8080
	}



	if cfg.Segment.MaxMessageBytes == 0 {
		cfg.Segment.MaxMessageBytes = 1 * 1024 * 1024 // 1MB
	}



	if cfg.Segment.MaxSegmentBytes == 0 {
		cfg.Segment.MaxSegmentBytes = 1 * 1024 * 1024 * 1024 // 1GB
	}



	if cfg.Segment.FlushIntervalMs == 0 {
		cfg.Segment.FlushIntervalMs = 1000
	}



	if cfg.Segment.FlushMessages == 0 {
		cfg.Segment.FlushMessages = 1000
	}



	if cfg.Retention.MaxAgeSeconds == 0 {
		cfg.Retention.MaxAgeSeconds = 7 * 24 * 60 * 60 // 7 days
	}



	if cfg.Retention.MaxBytes == 0 {
		cfg.Retention.MaxBytes = 1 * 1024 * 1024 * 1024 // 1GB
	}



	if cfg.Retention.CheckIntervalMs == 0 {
		cfg.Retention.CheckIntervalMs = 1000
	}



	return &cfg, nil


}
