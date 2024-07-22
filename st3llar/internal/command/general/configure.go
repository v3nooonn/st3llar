package general

import (
	"fmt"
	"log/slog"
	"os"

	rootCMD "github.com/v3nooom/st3llar/internal/command"
	"github.com/v3nooom/st3llar/internal/config"
	"github.com/v3nooom/st3llar/internal/constant"

	"github.com/spf13/cobra"
)

var (
	logLevel     string
	envPrefix    string
	environment  string
	organization string
	credential   string

	// configure represents the configure command
	configure = &cobra.Command{
		Use:   "configure",
		Short: "Configure the Stellar auto-action once for all.",
		Long:  `Configure the Stellar auto-action by passing one of them as flags`,
		Args:  cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if !cmd.Flags().Changed(constant.FLogLevel.ValStr()) &&
				!cmd.Flags().Changed(constant.FEnvPrefix.ValStr()) &&
				!cmd.Flags().Changed(constant.FEnvironment.ValStr()) &&
				!cmd.Flags().Changed(constant.FOrganization.ValStr()) &&
				!cmd.Flags().Changed(constant.FCredential.ValStr()) {
				return fmt.Errorf("at least one of the args must be set")
			}
			return nil
		},
		Run: configureFunc,
	}
)

func init() {
	rootCMD.Root.AddCommand(configure)

	configure.Flags().StringVarP(&logLevel, constant.FLogLevel.ValStr(), "",
		rootCMD.Vp.GetString("log-level"), "to specify the log level")

	configure.Flags().StringVarP(&envPrefix, constant.FEnvPrefix.ValStr(), "",
		rootCMD.Vp.GetString("env-prefix"), "Name prefix of the environment variables")

	configure.Flags().StringVarP(&environment, constant.FEnvironment.ValStr(), "",
		rootCMD.Vp.GetString("environment"), "environment of the CLI work with")

	configure.Flags().StringVarP(&organization, constant.FOrganization.ValStr(), "",
		rootCMD.Vp.GetString("organization"), "organization of the CLI work in")

	configure.Flags().StringVarP(&credential, constant.FCredential.ValStr(), "",
		rootCMD.Vp.GetString("credential"), "credential file path")
}

func configureFunc(cmd *cobra.Command, _ []string) {
	//for key, val := range rootCMD.Vp.AllSettings() {
	//	fmt.Printf("%s: %s\n", key, val)
	//}

	if cmd.Flags().Changed(constant.FLogLevel.ValStr()) {
		logLevel, _ = cmd.Flags().GetString(constant.FLogLevel.ValStr())
	}
	if cmd.Flags().Changed(constant.FEnvPrefix.ValStr()) {
		envPrefix, _ = cmd.Flags().GetString(constant.FEnvPrefix.ValStr())
	}
	if cmd.Flags().Changed(constant.FEnvironment.ValStr()) {
		environment, _ = cmd.Flags().GetString(constant.FEnvironment.ValStr())
	}
	if cmd.Flags().Changed(constant.FOrganization.ValStr()) {
		organization, _ = cmd.Flags().GetString(constant.FOrganization.ValStr())
	}
	if cmd.Flags().Changed(constant.FCredential.ValStr()) {
		credential, _ = cmd.Flags().GetString(constant.FCredential.ValStr())
	}

	cfg := config.Build(
		config.WithLogLevel(logLevel),
		config.WithEnvPrefix(envPrefix),
		config.WithEnvironment(environment),
		config.WithOrganization(organization),
		config.WithCredential(credential),
	)

	if err := config.WriteConfigFile(cfg, rootCMD.Vp.ConfigFileUsed()); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	slog.Info(fmt.Sprintf("updated config file: %s\n", rootCMD.Vp.ConfigFileUsed()))
}
