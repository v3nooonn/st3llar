package overview

import (
	"fmt"

	"github.com/v3nooom/st3llar/internal/cobra/command"

	"github.com/spf13/cobra"
)

// status represents the overview command
var status = &cobra.Command{
	Use:   "status",
	Short: "Shows the current users basic information",
	Long: `Showing the items below:
1. Organization and account info.
2. Session details.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("status called")
	},
}

func init() {
	command.Root.AddCommand(status)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// overview.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// overview.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
