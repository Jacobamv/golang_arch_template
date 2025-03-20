package config

import (
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Provide(ImportConfigs)

// Postgres configuration
type Postgres struct {
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	Host         string `mapstructure:"host"`
	Port         uint16 `mapstructure:"port"`
	SSLMode      string `mapstructure:"sslMode"`
	DatabaseName string `mapstructure:"databaseName"`
}

type Logger struct {
	Path  string `mapstructure:"path"`
	Type  string `mapstructure:"type"`
	Level string `mapstructure:"level"`
}

// Server configuration
type Server struct {
	Host      string `mapstructure:"host"`
	Port      uint16 `mapstructure:"port"`
	SecretKey string `mapstructure:"secretKey"`
}

type Oracle struct {
	Username               string `mapstructure:"username"`
	Password               string `mapstructure:"password"`
	ConnectString          string `mapstructure:"connectString"`
	PoolMaxSessions        string `mapstructure:"poolMaxSessions"`
	PoolMinSessions        string `mapstructure:"poolMinSessions"`
	PoolSessionMaxLifetime string `mapstructure:"poolSessionMaxLifetime"`
	PingEvery              int    `mapstructure:"pingEvery"`
}

type Example struct {
	Url      string `mapstructure:"url"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type RabbitMQ struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

// Config — struct is designed to combine the
// structures used
//
// Add or remove structures that are contained
// in your `config.json` file to this structure
type Config struct {
	Oracle   Oracle   `mapstructure:"oracle"`
	Postgres Postgres `mapstructure:"postgres"`
	RabbitMQ RabbitMQ `mapstructure:"rabbitmq"`
	Server   Server   `mapstructure:"server"`
	Example  Example  `mapstructure:"Example"`
	Logger   Logger   `mapstructure:"logger"`
}
