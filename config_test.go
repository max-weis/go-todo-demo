// +build unit

package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func tearDown(t *testing.T) {
	err := os.Unsetenv("DB_USER")
	assert.Nil(t, err)
	err = os.Unsetenv("DB_PASS")
	assert.Nil(t, err)
	err = os.Unsetenv("DB_HOST")
	assert.Nil(t, err)
	err = os.Unsetenv("DB_PORT")
	assert.Nil(t, err)
	err = os.Unsetenv("DB_NAME")
	assert.Nil(t, err)
}

func TestGetEnvVar(t *testing.T) {
	defer tearDown(t)

	err := os.Setenv("DB_USER", "test")
	assert.Nil(t, err)
	err = os.Setenv("DB_PASS", "test")
	assert.Nil(t, err)
	err = os.Setenv("DB_HOST", "test")
	assert.Nil(t, err)
	err = os.Setenv("DB_PORT", "test")
	assert.Nil(t, err)
	err = os.Setenv("DB_NAME", "test")
	assert.Nil(t, err)

	assert.Equal(t, "test", getEnvVar("DB_USER", ""))
	assert.Equal(t, "test", getEnvVar("DB_PASS", ""))
	assert.Equal(t, "test", getEnvVar("DB_HOST", ""))
	assert.Equal(t, "test", getEnvVar("DB_PORT", ""))
	assert.Equal(t, "test", getEnvVar("DB_NAME", ""))
}

func TestGetEnvVarDefaultValues(t *testing.T) {
	defer tearDown(t)

	assert.Equal(t, "postgres", getEnvVar("DB_USER", "postgres"))
	assert.Equal(t, "postgres", getEnvVar("DB_PASS", "postgres"))
	assert.Equal(t, "localhost", getEnvVar("DB_HOST", "localhost"))
	assert.Equal(t, "5432", getEnvVar("DB_PORT", "5432"))
	assert.Equal(t, "postgres", getEnvVar("DB_NAME", "postgres"))
}

func TestNewConfig(t *testing.T) {
	defer tearDown(t)

	err := os.Setenv("DB_USER", "test_user")
	assert.Nil(t, err)
	err = os.Setenv("DB_PASS", "test_pass")
	assert.Nil(t, err)
	err = os.Setenv("DB_HOST", "test_host")
	assert.Nil(t, err)
	err = os.Setenv("DB_PORT", "test_port")
	assert.Nil(t, err)
	err = os.Setenv("DB_NAME", "test_name")
	assert.Nil(t, err)

	config := NewConfig()

	assert.Equal(t, "test_user", config.dbUser)
	assert.Equal(t, "test_pass", config.dbPass)
	assert.Equal(t, "test_host", config.dbHost)
	assert.Equal(t, "test_port", config.dbPort)
	assert.Equal(t, "test_name", config.dbName)
}

func TestGetDSN(t *testing.T) {
	defer tearDown(t)

	err := os.Setenv("DB_USER", "test_user")
	assert.Nil(t, err)
	err = os.Setenv("DB_PASS", "test_pass")
	assert.Nil(t, err)
	err = os.Setenv("DB_HOST", "test_host")
	assert.Nil(t, err)
	err = os.Setenv("DB_PORT", "test_port")
	assert.Nil(t, err)
	err = os.Setenv("DB_NAME", "test_name")
	assert.Nil(t, err)

	config := NewConfig()

	dsn := config.getDSN()
	assert.Equal(t, "user=test_user password=test_pass host=test_host port=test_port dbname=test_name sslmode=disable", dsn)
}
