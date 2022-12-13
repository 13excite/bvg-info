package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

const DefaultConfigPath = "/etc/bvg_config.yaml"

// C is the global configuration
var C = Config{}

// Default returns config's object with default values
func Default() *Config {
	return &Config{
		LogLevel:                "info",
		LogEncoding:             "console",
		LoggerColor:             true,
		LoggerDisableStacktrace: true,
		LoggerDevMode:           true,
		LoggerDisableCaller:     false,
		LoggerDisabledHttp:      []string{"/version"},
		ServerHost:              "127.0.0.1",
		ServerPort:              "8081",
		VBB: VBB{
			ScanSecInterval: 90,
			API:             "https://v5.vbb.transport.rest",
			APIMaxResults:   10,
		},
	}
}

// ReadConfigFile reading and parsing configuration yaml file
func (conf *Config) ReadConfigFile(configPath string) {
	if configPath == "" {
		configPath = DefaultConfigPath
	}
	yamlConfig, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(yamlConfig, &conf)
	if err != nil {
		fmt.Printf("Could not unmarshal config %v", conf)
		log.Fatal(err)
	}
}
