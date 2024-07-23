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
	// configure represents the configure command
	configure = &cobra.Command{
		Use:   "configure",
		Short: "Configure the Stellar auto-action once for all.",
		Long:  `Configure the Stellar auto-action by passing one of them as flags`,
		Args:  cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
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
	configure.Flags().StringP(flagLogLevel, "",
		rootCMD.Vp.GetString(flagLogLevel), "to specify the log level")
	rootCMD.Vp.BindPFlag(flagLogLevel, configure.Flags().Lookup(flagLogLevel))

	flagEnvPrefix := constant.FlagEnvPrefix.ValStr()
	configure.Flags().StringP(flagEnvPrefix, "",
		rootCMD.Vp.GetString(flagEnvPrefix), "Name prefix of the environment variables")
	rootCMD.Vp.BindPFlag(flagEnvPrefix, configure.Flags().Lookup(flagEnvPrefix))

	flagEnv := constant.FlagEnvironment.ValStr()
	configure.Flags().StringP(flagEnv, "",
		rootCMD.Vp.GetString(flagEnv), "environment of the CLI work with")
	rootCMD.Vp.BindPFlag(flagEnv, configure.Flags().Lookup(flagEnv))

	flagOrg := constant.FlagOrganization.ValStr()
	configure.Flags().StringP(flagOrg, "",
		rootCMD.Vp.GetString(flagOrg), "organization of the CLI work in")
	rootCMD.Vp.BindPFlag(flagOrg, configure.Flags().Lookup(flagOrg))

	flagCred := constant.FlagCredential.ValStr()
	configure.Flags().StringP(flagCred, "",
		rootCMD.Vp.GetString(flagCred), "credential file path")
	rootCMD.Vp.BindPFlag(flagCred, configure.Flags().Lookup(flagCred))
}

func configureFunc(cmd *cobra.Command, _ []string) {
	cfg := config.Build(
		config.WithLogLevel(rootCMD.Vp.GetString(constant.FlagLogLevel.ValStr())),
		config.WithEnvPrefix(rootCMD.Vp.GetString(constant.FlagEnvPrefix.ValStr())),
		config.WithEnvironment(rootCMD.Vp.GetString(constant.FlagEnvironment.ValStr())),
		config.WithOrganization(rootCMD.Vp.GetString(constant.FlagOrganization.ValStr())),
		config.WithCredential(rootCMD.Vp.GetString(constant.FlagCredential.ValStr())),
	)

	if err := config.WriteConfig(cfg, rootCMD.Vp.ConfigFileUsed()); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	slog.Info(fmt.Sprintf("updated config file: %s\n", rootCMD.Vp.ConfigFileUsed()))
}
