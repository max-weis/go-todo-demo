package main

import (
	"fmt"
	"os"
)

type Config struct {
	dbUser string
	dbPass string
	dbHost string
	dbPort string
	dbName string
}

func NewConfig() Config {
	return Config{
		dbUser: getEnvVar("DB_USER", "postgres"),
		dbPass: getEnvVar("DB_PASS", "postgres"),
		dbHost: getEnvVar("DB_HOST", "localhost"),
		dbPort: getEnvVar("DB_PORT", "5432"),
		dbName: getEnvVar("DB_NAME", "postgres"),
	}
}

func getEnvVar(key, def string) string {
	val := os.Getenv(key)
	if val == "" {
		val = def
	}
	return val
}

func (c Config) getDSN() string {
	return fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		c.dbUser, c.dbPass, c.dbHost, c.dbPort, c.dbName,
	)
}
