package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/lara-go/app/app/models"
)

// CreateUsersTable migration.
type CreateUsersTable struct{}

// Migrate runs migrations.
func (m *CreateUsersTable) Migrate(tx *gorm.DB) error {
	return tx.AutoMigrate(&models.User{}).Error
}

// Rollback changes.
func (m *CreateUsersTable) Rollback(tx *gorm.DB) error {
	return tx.DropTable(&models.User{}).Error
}
