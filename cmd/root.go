package cmd

import (
	"fmt"
	"golerta-cli/lib"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfg lib.Config

var (
	cfgFile    = ""
	curlFlag   = false
	debugFlag  = false
	dryrunFlag = false
	version    = "1.0.10"
)

// rootCmd does not have a function. All action is inside subcommands.
var rootCmd = &cobra.Command{
	Use:     "golerta-cli",
	Short:   "Simple alerta.io client for sending alerts and heartbeats.",
	Version: version,
	//Run:   func(cmd *cobra.Command, args []string) { fmt.Println("root cmd mock") },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	// Print all the configuration parameters
	if debugFlag {
		fmt.Println("\nAll viper configuration parameters:")
		for key, value := range viper.AllSettings() {
			fmt.Printf("%s: %v\n", key, value)
		}
	}
	if dryrunFlag {
		fmt.Println("Dry run!")
		os.Exit(1)
	} else {
		err := rootCmd.Execute()
		if err != nil {
			os.Exit(1)
		}

	}
}

func init() {

	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is ./.golerta-cli)")
	rootCmd.PersistentFlags().BoolVarP(&debugFlag, "debug", "d", false, "Display info useful for debugging")
	rootCmd.PersistentFlags().BoolVarP(&curlFlag, "curl", "", false, "Generate a curl command representation of gathered parameters for testing")
	rootCmd.PersistentFlags().BoolVarP(&dryrunFlag, "dryrun", "", false, "Display info but don't post the endpoint")
	rootCmd.PersistentFlags().StringVarP(&cfg.Endpoint, "endpoint", "E", "", "Endpoint (Mandatory)")
	rootCmd.PersistentFlags().StringVarP(&cfg.APIKey, "apikey", "a", "", "Apikey (Mandatory)")
	viper.BindPFlag("endpoint", rootCmd.PersistentFlags().Lookup("endpoint"))
	viper.BindPFlag("apikey", rootCmd.PersistentFlags().Lookup("apikey"))
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
	viper.BindPFlag("curl", rootCmd.PersistentFlags().Lookup("curl"))
	viper.BindPFlag("dryrun", rootCmd.PersistentFlags().Lookup("dryrun"))

	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.DisableSuggestions = true
	//rootCmd.Version = version
	rootCmd.SetVersionTemplate(version)

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetEnvPrefix("GOLERTA_CLI")
	viper.AutomaticEnv() // read in environment variables that match
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
		viper.SetConfigType("env")
	} else {
		// Search config in home directory with name ".golerta-cli" (without extension).
		viper.AddConfigPath(".")
		viper.SetConfigType("env")
		viper.SetConfigName(".golerta-cli")
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		// good
	} else if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
		//panic(err)
		fmt.Fprintln(os.Stderr, "Provided Config file not found", viper.ConfigFileUsed())
		os.Exit(1)
	}

	// unmarshall viper config flags to our Config struct
	if err := viper.Unmarshal(&cfg); err == nil {
	} else {
		fmt.Println("Viper Unmarshal error? ", err)
		os.Exit(1)
	}
}
