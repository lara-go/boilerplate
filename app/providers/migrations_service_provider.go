package providers

import (
	"github.com/lara-go/app/app/database/migrations"
	"github.com/lara-go/larago"
	"github.com/lara-go/larago/database"
)

// MigrationsServiceProvider struct.
type MigrationsServiceProvider struct{}

// Boot provider
func (p *MigrationsServiceProvider) Boot(migrator *database.Migrator) {
	migrator.SetMigrations(
		&migrations.CreateUsersTable{},
	)
}

// Register service.
func (p *MigrationsServiceProvider) Register(application *larago.Application) {

}
