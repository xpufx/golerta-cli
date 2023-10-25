package cmd

import (
	"fmt"

	"golerta-cli/lib"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Displays program version.",
	Long: `Displays program version
	such that i can be used to send as alerta data`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(lib.Version)
		//		fmt.Println("endpoint at version cmd", cfg.Endpoint)
	},
}

func init() {

	versionCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flag for this command
		command.Flags().MarkHidden("config")
		command.Flags().MarkHidden("endpoint")
		// Call parent help func
		command.Parent().HelpFunc()(command, strings)
	})

	rootCmd.AddCommand(versionCmd)

}
