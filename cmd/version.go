package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Displays program version.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(rootCmd.Version)
		//		fmt.Println("endpoint at version cmd", cfg.Endpoint)
	},
}

func init() {

	versionCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flag for this command
		command.Flags().MarkHidden("config")
		command.Flags().MarkHidden("endpoint")
		command.Flags().MarkHidden("apikey")
		command.Flags().MarkHidden("curl")
		command.Flags().MarkHidden("dryrun")
		command.Flags().MarkHidden("debug")
		// Call parent help func
		command.Parent().HelpFunc()(command, strings)
	})

	rootCmd.AddCommand(versionCmd)

}
