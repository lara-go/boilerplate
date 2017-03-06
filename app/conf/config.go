package conf

import (
	"fmt"

	"github.com/lara-go/larago"
	dotaccess "github.com/maxwellhealth/go-dotaccess"
)

// Config of the app.
type Config struct {
	App      *App
	Database *Database
	HTTP     *HTTP
	Logs     *Logs
}

// ConfigLoader creates and returns a new config instance.
func ConfigLoader() larago.Config {
	return &Config{
		App:      MakeAppConfig(),
		Database: MakeDatabaseConfig(),
		HTTP:     MakeHTTPConfig(),
		Logs:     MakeLogsConfig(),
	}
}

// Env returns current environment name.
func (c *Config) Env() string {
	return c.App.Env
}

// Debug returs debug mode state.
func (c *Config) Debug() bool {
	return c.App.Debug
}

// Get value from config.
func (c *Config) Get(key string) interface{} {
	value, err := dotaccess.Get(c, key)

	if err != nil {
		panic(fmt.Sprintf("Can not resolve config value: %s", key))
	}

	return value
}

// Set value from config.
func (c *Config) Set(key string, value interface{}) {
	err := dotaccess.Set(c, key, value)

	if err != nil {
		panic(fmt.Sprintf("Can not resolve config value: %s", key))
	}
}
