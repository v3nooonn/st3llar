package oauth

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
2. Account and Password

Available credentials are stored locally in the default path: ./credentials.json, once login successfully.
And you can use --credentials to specify the path of the credentials file.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		//# 对于Windows 64位 GOOS=windows GOARCH=amd64 去建立 -o mycli-windows.exe ./path/to/package
		//# 对于macOS 64位 GOOS=darwin GOARCH=amd64 去建立 -o mycli-macos ./path/to/package
		//# 对于Linux 64位 GOOS=linux GOARCH=amd64 去建立 -o mycli-linux ./path/to/package
		//survey.AskOne() //TODO: pwd interactiveness

		fmt.Println("login called")
		fmt.Printf("args: %v\n", args)
		if time.Now().Second()%2 == 0 {
			return fmt.Errorf("error: invalid credentials")
		}

		fmt.Println("login successful")
		return nil
	},
}

func init() {
	command.Root.AddCommand(login)

	login.Flags().StringP("username", "U", "", "Name of the account")
	login.Flags().StringP("organization", "O", "", "Name of the organization")
}
