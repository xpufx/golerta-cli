package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"golerta-cli/lib"

	"github.com/spf13/cobra"
)

// heartbeatCmd represents the heartbeat command
var heartbeatCmd = &cobra.Command{
	Use:   "heartbeat",
	Short: "Send a heartbeat to alerta endpoint",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if cfg.APIKey == "" || cfg.Endpoint == "" {
			fmt.Println("Error: Mandatory parameters (apikey, endpoint) must be provided.")
			os.Exit(1)
		} else {
			cfg.Endpoint += "/heartbeat"
			postHeartbeat(&cfg)
		}
	},
}

func init() {
	rootCmd.AddCommand(heartbeatCmd)

	heartbeatCmd.Flags().StringVarP(&cfg.Origin, "origin", "o", "", "Origin string")
	heartbeatCmd.Flags().StringArrayVarP(&cfg.Tags, "tag", "", nil, "Tags")
	heartbeatCmd.Flags().IntVar(&cfg.Timeout, "timeout", 0, "Timeout (integer)")
	heartbeatCmd.Flags().MarkHidden("config")
	// service
	// tag
	// customer
	// delete?
}

func postHeartbeat(c *lib.Config) {
	if debugFlag || dryrunFlag {
		// Print the configuration variables
		fmt.Println("Configuration Variables:")
		fmt.Println("API Key:", c.APIKey)
		fmt.Println("Endpoint:", c.Endpoint)
		fmt.Println("Group:", c.Group)
		fmt.Println("Environment:", c.Environment)
		fmt.Println("Origin:", c.Origin)
		fmt.Println("Severity:", c.Severity)
		fmt.Println("Service:", c.Service)
		fmt.Println("Tags:", c.Tags)
		fmt.Println("Timeout:", c.Timeout)
		fmt.Println("Value:", c.Value)
	}

	// Convert the Config struct to JSON
	jsonData, err := json.Marshal(c)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Send an HTTP POST request with the JSON data and "Authorization: Key" header
	client := &http.Client{}
	req, err := http.NewRequest("POST", c.Endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Set the authorization header as "Authorization: Key"
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Key "+c.APIKey) // Set the API key as the authorization header

	// Print the request headers for debugging
	if debugFlag || dryrunFlag {
		// Print the JSON data
		fmt.Println("JSON Data:")
		fmt.Println(string(jsonData))
		// Print the equivalent curl command
		fmt.Println("Request Headers:")
		for key, values := range req.Header {
			for _, value := range values {
				fmt.Printf("%s: %s\n", key, value)
			}
		}
	}
	if debugFlag || curlFlag || dryrunFlag {
		fmt.Println("\nEquivalent curl command:")
		curlCommand := generateCurlCommand(c.Endpoint, c.APIKey, jsonData)

		fmt.Println(curlCommand)
	}

	if !dryrunFlag {
		// network api call
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		// Check the response status
		if resp.StatusCode == http.StatusCreated {
			fmt.Println("Heartbeat was sent.")
			if !debugFlag {
				os.Exit(0)
			}
		} else if resp.StatusCode != http.StatusOK {
			fmt.Printf("Error: Unexpected response status code %d\n", resp.StatusCode)
			os.Exit(1)
		}

		// Read and print the response body
		if debugFlag {
			responseData := make([]byte, 1024)
			n, _ := io.ReadFull(resp.Body, responseData)
			fmt.Println("HTTP POST Response:")
			fmt.Println(string(responseData[:n]))
		}
	}
}
