package auth

import (
	"fmt"
	"time"

	"github.com/v3nooom/st3llar-helper/internal/cobra/command"

	"github.com/spf13/cobra"
)

// signIn represents the sign-in command
var signIn = &cobra.Command{
	Use:   "sign-in",
	Short: "A brief description of sign-in",
	Long: `A longer description. For example:

sign-in

Cobra is a CLI library for Go that empowers applications.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("sign-in called")
		if time.Now().Second()%2 == 0 {
			return fmt.Errorf("error: invalid credentials")
		}
		return nil
	},
}

func init() {
	command.Root.AddCommand(signIn)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// aboutCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// aboutCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
