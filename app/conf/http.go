package conf

import "github.com/caarlos0/env"

// HTTP config.
type HTTP struct {
	Secure      bool   `env:"HTTP_SECURE"`
	Listen      string `env:"HTTP_LISTEN" envDefault:"0.0.0.0:8080"`
	LogRequests bool   `env:"HTTP_LOG"`
}

// MakeHTTPConfig values.
func MakeHTTPConfig() *HTTP {
	config := &HTTP{}
	env.Parse(config)

	return config
}
