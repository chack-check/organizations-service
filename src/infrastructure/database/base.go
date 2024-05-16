package database

import (
	"errors"
	"fmt"

	"github.com/chack-check/organizations-service/infrastructure/settings"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnection() *gorm.DB {
	db, err := gorm.Open(postgres.Open(settings.Settings.APP_DATABASE_DSN), &gorm.Config{})
	if err != nil {
		panic(errors.Join(fmt.Errorf("error when connecting to database"), err))
	}

	return db
}

func SetupMigrations() {
	DatabaseConnection.AutoMigrate(&DBSavedFile{}, &DBRole{}, &DBMember{}, &DBOrganization{})
}

var DatabaseConnection *gorm.DB = GetConnection()
