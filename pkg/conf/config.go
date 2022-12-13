package conf

type Config struct {
	ServerHost              string   `yaml:"server_host"`
	ServerPort              string   `yaml:"server_port"`
	LogLevel                string   `yaml:"log_level"`
	PidFile                 string   `yaml:"pid_file"`
	LogEncoding             string   `yaml:"log_encoding"`
	LoggerColor             bool     `yaml:"logger_color"`
	LoggerDisableStacktrace bool     `yaml:"logger_disable_stacktrace"`
	LoggerDevMode           bool     `yaml:"logger_dev_mode"`
	LoggerDisableCaller     bool     `yaml:"logger_disable_caller"`
	LoggerDisabledHttp      []string `yaml:"log_disabled_http"`
	VBB                     VBB      `yaml:"vbb"`
}

type VBB struct {
	ScanSecInterval int    `yaml:"scan_interval"`
	API             string `yaml:"api_endpoint"`
	APIMaxResults   int    `yaml:"api_max_result"`
}
