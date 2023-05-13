package postgres

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewPostgres(t *testing.T) {
	// Initialize test config
	config := &SConfig{
		Host:         "localhost",
		Port:         "5432",
		User:         "admin",
		Password:     "admin",
		DatabaseName: "go_clean_architecture",
		SslMode:      "disable",
		TimeZone:     "Asia/Tehran",
	}

	// Test that NewPostgres initializes and returns an SPostgres object
	sPostgres := NewPostgres(config)
	require.NotNil(t, sPostgres)
}

func TestSPostgres_Close(t *testing.T) {
	// Initialize test config
	config := &SConfig{
		Host:         "localhost",
		Port:         "5432",
		User:         "admin",
		Password:     "admin",
		DatabaseName: "go_clean_architecture",
		SslMode:      "disable",
		TimeZone:     "Asia/Tehran",
	}

	// Initialize SPostgres object
	sPostgres := NewPostgres(config)

	// Test that Close method returns no error
	err := sPostgres.Close()
	assert.NoError(t, err)
}

func TestSPostgres_Transaction(t *testing.T) {
	// Initialize test config
	config := &SConfig{
		Host:         "localhost",
		Port:         "5432",
		User:         "admin",
		Password:     "admin",
		DatabaseName: "go_clean_architecture",
		SslMode:      "disable",
		TimeZone:     "Asia/Tehran",
	}

	// Initialize SPostgres object
	sPostgres := NewPostgres(config)

	// Test that Transaction method runs the passed function and returns no error
	err := sPostgres.Transaction(func(tx *SPostgres) error {
		return nil
	})
	assert.NoError(t, err)
}
