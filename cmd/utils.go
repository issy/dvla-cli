package cmd

import "os"

func readMOTApiKeyFromEnvVar() string {
	apiKey := os.Getenv("DVLA_MOT_API_KEY")
	if apiKey == "" {
		panic("Must set DVLA_MOT_API_KEY env var")
	}
	return apiKey
}

func readVESApiKeyFromEnvVar() string {
	apiKey := os.Getenv("DVLA_VES_API_KEY")
	if apiKey == "" {
		panic("Must set DVLA_VES_API_KEY env var")
	}
	return apiKey
}
