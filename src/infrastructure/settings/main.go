package settings

import (
	"fmt"
	"os"
	"strconv"
)

type SettingsSchema struct {
	APP_DATABASE_DSN  string
	APP_ALLOW_ORIGINS string
	APP_SECRET_KEY    string
	APP_PORT          int
}

func InitSettings() SettingsSchema {
	databaseDsn := os.Getenv("APP_DATABASE_DSN")
	if databaseDsn == "" {
		panic(fmt.Errorf("you need to specify `APP_DATABASE_DSN` environment variable"))
	}

	secretKey := os.Getenv("APP_SECRET_KEY")
	if secretKey == "" {
		panic(fmt.Errorf("you need to specify `APP_SECRET_KEY` environment variable"))
	}

	origins := os.Getenv("APP_ALLOW_ORIGINS")
	if origins == "" {
		origins = "*"
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8000"
	}
	portInt, err := strconv.Atoi(port)
	if err != nil {
		panic(err)
	}

	return SettingsSchema{
		APP_DATABASE_DSN:  databaseDsn,
		APP_ALLOW_ORIGINS: origins,
		APP_PORT:          portInt,
		APP_SECRET_KEY:    secretKey,
	}
}

var Settings SettingsSchema = InitSettings()
