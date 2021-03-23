package config

import (
	"log"

	"github.com/spf13/viper"
)

const (
	bankService  = "BANK_SERVICE"
	fraudService = "FRAUD_SERVICE"
)

type Config struct {
	httpServerConfig HTTPServerConfig
	logConfig        LogConfig
	logFileConfig    LogFileConfig
	dbConfig         DBConfig
}

func (config Config) GetDBConfig() DBConfig {
	return config.dbConfig
}

func (config Config) GetHTTPServerConfig() HTTPServerConfig {
	return config.httpServerConfig
}

func (config Config) GetLogConfig() LogConfig {
	return config.logConfig
}

func (config Config) GetLogFileConfig() LogFileConfig {
	return config.logFileConfig
}

func NewConfig(configFile string) Config {
	viper.AutomaticEnv()

	if configFile != "" {
		viper.SetConfigFile(configFile)
		if err := viper.ReadInConfig(); err != nil {
			log.Fatal(err)
		}
	}

	return Config{
		httpServerConfig: newHTTPServerConfig(),
		logConfig:        newLogConfig(),
		logFileConfig:    newLogFileConfig(),
		dbConfig:         NewDBConfig(),
	}
}
