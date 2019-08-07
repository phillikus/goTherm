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

	if config.Database.ConnectionString != "" {
		t.Errorf("Opening empty config file should not provide any data")
	}
}

func TestLoadConfigurationValidJson(t *testing.T) {
	config := LoadConfiguration("test_data/valid_config.json")

	expected := "host=localhost port=5432 user=testuser password=Test123 dbname=testdb"

	if config.Database.ConnectionString != expected {
		t.Errorf("ConnectionString read incorrectly: expected: %v but was %v", expected, config.Database.ConnectionString)
	}
}
