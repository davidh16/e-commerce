package config

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/tkanos/gonfig"
	"golang.org/x/oauth2/google"
	"log"
	"os"
)

// Config struct can be expanded if more env variables are introduced

type Config struct {
	PgUrl              string `env:"PG_URL"`
	Port               string `env:"PORT"`
	SecretKey          string `env:"SECRET_KEY"`
	StorageBucket      string `env:"STORAGE_BUCKET"`
	StorageCredentials *google.Credentials
	SmtpHost           string `env:"SMTP_HOST"`
	SmtpPort           string `env:"SMTP_PORT"`
	GoogleAppPassword  string `env:"GOOGLE_APP_PASSWORD"`
	SmtpFrom           string `env:"SMTP_FROM"`
	BaseUrl            string `env:"BASE_URL"`
}

func GetConfig() Config {

	configuration := Config{}

	if err := godotenv.Load(); err != nil {
		fmt.Println("Failed to load .env file")
		os.Exit(1)
	}

	err := gonfig.GetConf("", &configuration)
	if err != nil {
		fmt.Println(err)
		return Config{}
	}

	storageKey := os.Getenv("STORAGE_BUCKET_SIGNED_URL")
	if storageKey != "" {
		storageCred, err := decodeGoogleKey(storageKey)
		if err != nil {
			log.Panic(err)
		}

		configuration.StorageCredentials = &storageCred
	}

	return configuration
}

func decodeGoogleKey(key string) (google.Credentials, error) {
	if key == "" {
		return google.Credentials{}, errors.New("missing google key")
	}

	dKey, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return google.Credentials{}, errors.New("failed to unmarshal google key")
	}

	ctx := context.Background()
	cred, err := google.CredentialsFromJSON(ctx, dKey)
	if err != nil {
		log.Panic(err)
	}

	return *cred, nil

}
