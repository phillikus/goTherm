package config

import (
	"testing"
)

func TestLoadConfigurationInvalidPath(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Opening config file with invalid path should panic.")
		}
	}()

	LoadConfiguration("/invalidFilePath")
}

func TestLoadConfigurationEmptyJson(t *testing.T) {
	config := LoadConfiguration("test_data/empty_config.json")

	if config.Database.Host != "" || config.Database.User != "" || config.Database.Password != "" {
		t.Errorf("Opening empty config file should not provide any data")
	}
}

func TestLoadConfigurationValidJson(t *testing.T) {
	config := LoadConfiguration("test_data/valid_config.json")

	if config.Database.Host != "localhost" {
		t.Errorf("Host read incorrectly: expected: %v but was %v", "localhost", config.Database.Host)
	}
	if config.Database.User != "user" {
		t.Errorf("User read incorrectly: expected: %v but was %v", "user", config.Database.User)
	}
	if config.Database.Password != "Test123" {
		t.Errorf("Password read incorrectly: expected: %v but was %v", "Test123", config.Database.Password)
	}
}
