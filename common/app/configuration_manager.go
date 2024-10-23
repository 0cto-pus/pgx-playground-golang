package app

import "rest-playground/common/postgresql"

type ConfigurationManager struct {
	PostgreSqlConfig postgresql.Config
}

func NewConfigurationManager() *ConfigurationManager {
	postgreSqlConfig := getPostgreSqlConfig()
	return &ConfigurationManager{
		PostgreSqlConfig: postgreSqlConfig,
	}
}

func getPostgreSqlConfig() postgresql.Config {
	return postgresql.Config{
		Host:                  "localhost",
		Port:                  "6432",
		UserName:              "postgres",
		Password:              "postgres",
		DbName:                "productapp",
		MaxConnection:        "10",
		MaxConnectionIdleTime: "30s",
	}
}