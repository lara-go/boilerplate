package conf

import "github.com/caarlos0/env"

// Database config.
type Database struct {
	Driver    string `env:"DB_DRIVER" envDefault:"sqlite3"`
	DSN       string `env:"DB_DSN" envDefault:"database.sqlite"`
	Charset   string
	Collation string
}

// MakeDatabaseConfig values.
func MakeDatabaseConfig() *Database {
	config := &Database{
		Charset:   "utf8",
		Collation: "utf8_unicode_ci",
	}
	env.Parse(config)

	return config
}
