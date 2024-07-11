package overview

import (
	"fmt"

	"github.com/v3nooom/st3llar/internal/cobra/command"

	"github.com/spf13/cobra"
)

// self represents the overview command
var self = &cobra.Command{
	Use:   "self",
	Short: "Shows the current users basic information",
	Long: `Showing the items below:
1. Organization and account info.
2. Session details.
3. Brief info about the available Lambdas.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("self called")
	},
}

func init() {
	command.Root.AddCommand(self)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// overview.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// overview.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
