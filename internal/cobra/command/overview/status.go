package overview

import (
	"fmt"

	"github.com/v3nooom/st3llar-helper/internal/cobra/command"

	"github.com/spf13/cobra"
)

// status represents the overview command
var status = &cobra.Command{
	Use:   "status",
	Short: "A brief description of overview",
	Long: `A longer description. For example:

status

Cobra is a CLI library for Go that empowers applications.`,
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
