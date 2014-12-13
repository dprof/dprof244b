
type Config struct {
	Aggregator AggregatorConfig `json:aggregator`
	LogFiles   []LogFileConfig  `json:logfiles`
}

type AggregatorConfig struct {
	Servers       []string `json:servers`
	timeout        time.Duration
}

type LogFileConfig struct {
	Path    string          `json:path`
	Type    string          `json:type`
}


