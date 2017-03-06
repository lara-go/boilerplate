package conf

import "github.com/caarlos0/env"

// App config.
type App struct {
	Env   string `env:"APP_ENV" envDefault:"production"`
	Debug bool   `env:"APP_DEBUG"`
}

// MakeAppConfig values.
func MakeAppConfig() *App {
	config := &App{}
	env.Parse(config)

	return config
}
