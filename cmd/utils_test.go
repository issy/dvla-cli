package cmd

import "testing"

func TestVESAPIKey(t *testing.T) {
	apiKey := readVESApiKeyFromEnvVar()
	if apiKey == "" {
		t.Fatal("API key not set")
	}
}

func TestMOTAPIKey(t *testing.T) {
	apiKey := readMOTApiKeyFromEnvVar()
	if apiKey == "" {
		t.Fatal("API key not set")
	}
}
