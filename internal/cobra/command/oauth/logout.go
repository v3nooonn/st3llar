package oauth

import (
	"fmt"
	"time"

	"github.com/v3nooom/st3llar/internal/cobra/command"

	"github.com/spf13/cobra"
)

// logout represents the sign-in command
var logout = &cobra.Command{
	Use:   "logout",
	Short: "Logout from the Stellar auto-task",
	Long:  `Logout will clear the current session and remove the access token from the local credentials.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("logout called")
		if time.Now().Second()%2 == 0 {
			return fmt.Errorf("error: invalid credentials")
		}
		return nil
	},
}

func init() {
	command.Root.AddCommand(logout)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// aboutCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// aboutCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
