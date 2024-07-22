package general

import (
	rootCMD "github.com/v3nooom/st3llar/internal/command"
	"github.com/v3nooom/st3llar/internal/constant"
	"os"

	"github.com/spf13/cobra"
)

var (
	workDir      string
	envPrefix    string
	environment  string
	organization string

	// configure represents the configure command
	configure = &cobra.Command{
		Use:   "configure",
		Short: "Configure the Stellar auto-action once for all.",
		Long: `Configure the Stellar auto-action by passing one of them below:

1. Configuration file path
2. Key-Value pairs`,
		Args:      cobra.OnlyValidArgs,
		ValidArgs: []string{"workdir", "env-prefix", "environment"},
		Run:       configureFunc,
	}
)

func init() {
	rootCMD.Root.AddCommand(configure)

	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	configure.Flags().StringVarP(&workDir, "workdir", "",
		home,
		"to specify the path of the configuration file")
	configure.Flags().StringVarP(&envPrefix, "env-prefix", "",
		constant.EnvPrefix.ValStr(),
		"Name prefix of the environment variables")
	configure.Flags().StringVarP(&environment, "environment", "",
		constant.Environment.ValStr(),
		"environment of the CLI work with")
	configure.Flags().StringVarP(&organization, "organization", "",
		constant.Organization.ValStr(),
		"organization of the CLI work in")

	rootCMD.Vp.BindPFlag("env", rootCMD.Root.PersistentFlags().Lookup("debug"))
}

func configureFunc(cmd *cobra.Command, args []string) {
	// TODO: error handling
}
