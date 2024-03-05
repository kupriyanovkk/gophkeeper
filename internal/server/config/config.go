package config

import (
	"github.com/caarlos0/env/v10"
	"github.com/kupriyanovkk/gophkeeper/pkg/logger"
)

type Config struct {
	Address     string `env:"ADDRESS" envDefault:"localhost"`
	Port        string `env:"PORT" envDefault:"8080"`
	SSLCertPath string `env:"SSL_CERT_PATH" envDefault:"./certs/localhost.crt"`
	SSLKeyPath  string `env:"SSL_KEY_PATH" envDefault:"./certs/localhost.key"`
	JWTSecret   string `env:"JWT_SECRET" envDefault:"jwt_secret"`
	JWTExp      string `env:"JWT_EXP" envDefault:"24"`
	DatabaseDSN string `env:"DATABASE_DSN" envDefault:"mongodb://localhost:27017/"`
}

var config Config

// NewConfig returns a Config and parses it.
//
// No parameters.
// Returns a Config.
func NewConfig() Config {
	config.parse()

	return config
}

// parse is a method of Config that parses the configuration.
//
// No parameters.
// No return type.
func (c *Config) parse() {
	if err := env.Parse(&config); err != nil {
		logger.NewLogger().Error(err.Error())
	}
}
