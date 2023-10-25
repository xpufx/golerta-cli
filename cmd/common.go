package cmd

import "fmt"

func generateCurlCommand(endpoint, apiKey string, jsonData []byte) string {
	// Generate the curl command with the JSON data and authorization header
	curlCommand := fmt.Sprintf(
		"curl -X POST -H 'Authorization: Key %s' -H 'Content-Type: application/json' -d '%s' %s",
		apiKey, jsonData, endpoint)

	return curlCommand
}
