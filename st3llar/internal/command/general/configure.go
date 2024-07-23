package general

import (
	"fmt"
	"github.com/v3nooom/st3llar/internal/util"
	"log/slog"
	"os"

	rootCMD "github.com/v3nooom/st3llar/internal/command"
	"github.com/v3nooom/st3llar/internal/config"
	"github.com/v3nooom/st3llar/internal/constant"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// configure represents the configure command
	configure = &cobra.Command{
		Use:   "configure",
		Short: "Configure the Stellar auto-action once for all.",
		Long:  `Configure the Stellar auto-action by passing one of them as flags`,
		Args:  cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			util.PreRunBindFlags(cmd, args)

			if !cmd.Flags().Changed(constant.FlagLogLevel.ValStr()) &&
				!cmd.Flags().Changed(constant.FlagEnvPrefix.ValStr()) &&
				!cmd.Flags().Changed(constant.FlagEnvironment.ValStr()) &&
				!cmd.Flags().Changed(constant.FlagOrganization.ValStr()) &&
				!cmd.Flags().Changed(constant.FlagCredential.ValStr()) {
				return fmt.Errorf("at least one of the args must be set")
			}
			return nil
		},
		Run: configureFunc,
	}
)

func init() {
	rootCMD.Root.AddCommand(configure)

	flagLogLevel := constant.FlagLogLevel.ValStr()
	configure.Flags().StringP(
		flagLogLevel,
		"",
		viper.GetString(flagLogLevel),
		"configure the log level")

	flagEnvPrefix := constant.FlagEnvPrefix.ValStr()
	configure.Flags().StringP(
		flagEnvPrefix,
		"",
		viper.GetString(flagEnvPrefix),
		"configure the name prefix of the environment variables")

	flagEnv := constant.FlagEnvironment.ValStr()
	configure.Flags().StringP(
		flagEnv,
		"",
		viper.GetString(flagEnv),
		"configure the environment of the CLI work with")

	flagOrg := constant.FlagOrganization.ValStr()
	configure.Flags().StringP(
		flagOrg,
		"",
		viper.GetString(flagOrg),
		"configure the organization of the CLI work in")

	flagCred := constant.FlagCredential.ValStr()
	configure.Flags().StringP(
		flagCred,
		"",
		viper.GetString(flagCred),
		"configure the credential file path")
}

func configureFunc(_ *cobra.Command, _ []string) {
	cfg := config.Build(
		config.WithLogLevel(viper.GetString(constant.FlagLogLevel.ValStr())),
		config.WithEnvPrefix(viper.GetString(constant.FlagEnvPrefix.ValStr())),
		config.WithEnvironment(viper.GetString(constant.FlagEnvironment.ValStr())),
		config.WithOrganization(viper.GetString(constant.FlagOrganization.ValStr())),
		config.WithCredential(viper.GetString(constant.FlagCredential.ValStr())),
	)
	fmt.Printf("latest configuration: %#v\n", cfg)

	if err := config.WriteConfig(cfg, viper.ConfigFileUsed()); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	slog.Info(fmt.Sprintf("updated config file: %s\n", viper.ConfigFileUsed()))
}
