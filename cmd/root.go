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
	version    = "1.0.8"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "golerta-cli",
	Short: "Simple alerta.io client for sending alerts and heartbeats.",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.golerta-cli)")
	rootCmd.PersistentFlags().BoolVarP(&debugFlag, "debug", "d", false, "Display info useful for debugging")
	rootCmd.PersistentFlags().BoolVarP(&curlFlag, "curl", "", false, "Generate a curl command representation of gathered parameters for testing")
	rootCmd.PersistentFlags().BoolVarP(&dryrunFlag, "dryrun", "", false, "Display info but don't post the endpoint")
	rootCmd.PersistentFlags().StringVarP(&cfg.Endpoint, "endpoint", "E", "", "Endpoint (Mandatory)")
	viper.BindPFlag("endpoint", rootCmd.PersistentFlags().Lookup("endpoint"))
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
	viper.BindPFlag("curl", rootCmd.PersistentFlags().Lookup("curl"))
	viper.BindPFlag("dryrun", rootCmd.PersistentFlags().Lookup("dryrun"))

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.DisableSuggestions = true
	rootCmd.Version = version

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".golerta-cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("env")
		viper.SetConfigName(".golerta-cli")
	}
	viper.SetEnvPrefix("GOLERTA_CLI")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Info: Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println("Config file error: ", err)
		os.Exit(1)
	}

	if err := viper.Unmarshal(&cfg); err == nil {
	} else {
		fmt.Println("Viper Unmarshal error? ", err)
		os.Exit(1)
	}
}
