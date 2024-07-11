package auth

import (
	"fmt"
	"time"

	"github.com/v3nooom/st3llar/internal/cobra/command"

	"github.com/spf13/cobra"
)

// login represents the sign-in command
var login = &cobra.Command{
	Use:   "login",
	Short: "Login to the Stellar auto-task.",
	Long: `Login to the Stellar auto-task by the pre-established credentials:
1. Organization
2. Account and password

Available credentials are stored locally in the default path: ./credentials.json, once login successfully.
And you can use --credentials to specify the path of the credentials file.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("login called")
		if time.Now().Second()%2 == 0 {
			return fmt.Errorf("error: invalid credentials")
		}
		return nil
	},
}

func init() {
	command.Root.AddCommand(login)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// aboutCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// aboutCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
