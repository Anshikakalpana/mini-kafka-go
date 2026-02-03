package broker

type BrokerConfig struct {
	Broker BrokerSection `yaml:"broker"`
}

type BrokerSection struct {
	Service   ServiceSection   `yaml:"service"`
	Segment   SegmentSection   `yaml:"segment"`
	Retention RetentionSection `yaml:"retention"`
	Logging   LoggingSection   `yaml:"logging"`
}

type ServiceSection struct {
	Name    string `yaml:"name"`
	DataDir string `yaml:"dataDir"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
}

type SegmentSection struct {
	MaxMessageBytes  int64 `yaml:"maxMessageBytes"`
	MaxSegmentBytes  int64 `yaml:"maxSegmentBytes"`
	FlushIntervalMs  int   `yaml:"flushIntervalMs"`
	FlushMessages    int   `yaml:"flushMessages"`
}

type RetentionSection struct {
	MaxAgeSeconds    int64 `yaml:"maxAgeSeconds"`
	MaxBytes         int64 `yaml:"maxBytes"`
	CheckIntervalMs  int   `yaml:"checkIntervalMs"`
}

type LoggingSection struct {
	Level           string `yaml:"level"`
	LogDir          string `yaml:"logDir"`
	MaxLogFiles     int    `yaml:"maxLogFiles"`
	MaxLogSizeBytes int64  `yaml:"maxLogSizeBytes"`
}
