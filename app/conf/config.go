package conf

import "github.com/lara-go/larago"

// Config of the app.
type Config struct {
	App      *App
	Database *Database
	HTTP     *HTTP
	Logs     *Logs
}

// NewConfig creates and returns a new config instance.
func NewConfig() larago.Config {
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
