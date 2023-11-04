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

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send an alert to alerta endpoint.",
	Run: func(cmd *cobra.Command, args []string) {
		// Check for mandatory parameters after all values have been gathered
		if cfg.APIKey == "" || cfg.Endpoint == "" || cfg.Event == "" || cfg.Resource == "" {
			fmt.Println("Error: Mandatory parameters (apikey, endpoint, event, resource) must be provided.")
			os.Exit(1)
		} else {
			cfg.Endpoint += "/alert"
			postAlert(&cfg)
		}
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)

	sendCmd.Flags().StringVarP(&cfg.Event, "event", "e", "", "Event string (mandatory)")
	sendCmd.Flags().StringVarP(&cfg.Type, "type", "t", "", "Event type string")
	sendCmd.Flags().StringVarP(&cfg.Group, "group", "g", "", "Group string")
	sendCmd.Flags().StringVarP(&cfg.Resource, "resource", "r", "", "Resource string (mandatory)")
	sendCmd.Flags().StringVarP(&cfg.Environment, "environment", "", "", "Environment string")
	sendCmd.Flags().StringVarP(&cfg.Origin, "origin", "o", "", "Origin string")
	sendCmd.Flags().StringVarP(&cfg.RawData, "raw-data", "", "", "Raw data string")
	sendCmd.Flags().StringVarP(&cfg.Severity, "severity", "s", "normal", "Severity ('ok', 'normal', 'major', 'minor', 'critical')")
	sendCmd.Flags().StringArrayVarP(&cfg.Service, "service", "x", nil, "Service (multiple invokation allowed)")
	sendCmd.Flags().StringArrayVarP(&cfg.Tags, "tag", "", nil, "Tags (multiple invokation allowed)")
	//sendCmd.Flags().StringSetVar(&cfg.Attributes, 0, "attributes", "Attributes like region=eu (multiple invokation allowed)")
	sendCmd.Flags().StringVarP(&cfg.Text, "text", "T", "", "Text string")
	sendCmd.Flags().IntVar(&cfg.Timeout, "timeout", 0, "Timeout (integer)")
	sendCmd.Flags().IntVar(&cfg.Value, "value", 0, "Integer value")
}

func postAlert(c *lib.Config) {
	if debugFlag || dryrunFlag {
		// Print the configuration variables
		fmt.Println("Configuration Variables:")
		fmt.Println("API Key:", c.APIKey)
		fmt.Println("Endpoint:", c.Endpoint)
		fmt.Println("Event:", c.Event)
		fmt.Println("Type:", c.Type)
		fmt.Println("Group:", c.Group)
		fmt.Println("Environment:", c.Environment)
		fmt.Println("Origin:", c.Origin)
		fmt.Println("Raw Data:", c.RawData)
		fmt.Println("Resource:", c.Resource)
		fmt.Println("Severity:", c.Severity)
		fmt.Println("Service:", c.Service)
		fmt.Println("Tags:", c.Tags)
		//fmt.Println("Attributes:", c.Attributes)
		fmt.Println("Text:", c.Text)
		fmt.Println("Timeout:", c.Timeout)
		fmt.Println("Value:", c.Value)
		fmt.Println()
		fmt.Println("Debug: ", debugFlag)
		fmt.Println("Dryrun:", dryrunFlag)
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

	// if --dryrun is specified do not make the final HTTP POST call
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
			fmt.Println("Alerta record was added.")
			if !debugFlag {
				os.Exit(0)
			}
		} else if resp.StatusCode != http.StatusOK {
			// sometimes the server may return an error while still accepting the alert
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
