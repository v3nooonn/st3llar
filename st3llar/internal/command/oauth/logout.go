package oauth

import (
	rootCMD "github.com/v3nooom/st3llar/internal/command"
	"github.com/v3nooom/st3llar/internal/config"
	"github.com/v3nooom/st3llar/internal/util"
	"os"

	"github.com/spf13/cobra"
)

// logout represents the logout command
var logout = &cobra.Command{
	Use:   "logout",
	Short: "Logout the Stellar auto-task.",
	Long: `Logout the Stellar auto-action and delete the credentials from config:

If there are multiple credentials are set, the credential that is set 
in the config will be deleted.`,
	Args: cobra.NoArgs,
	//PreRunE: func(cmd *cobra.Command, args []string) error {
	//	util.PreRunBindFlags(cmd, args)
	//	return nil
	//},
	RunE: logoutFunc,
}

func init() {
	rootCMD.Root.AddCommand(logout)
}

func logoutFunc(_ *cobra.Command, _ []string) error {
	cfg, _ := config.FindConfig()

	if util.IsExists(cfg.Credential) {
		err := os.Remove(cfg.Credential)
		if err != nil {
			return err
		}
	}

	return nil
}
