package command

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/v3nooom/st3llar/internal/config"
	"github.com/v3nooom/st3llar/internal/constant"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Vp *viper.Viper
)

// Root represents the base command when called without any subcommands
var Root = &cobra.Command{
	Use:       "st3llar",
	Args:      cobra.OnlyValidArgs,
	ValidArgs: []string{"login", "logout"},
	Short:     "Stellar AutoTask CLI: st3llar",
	Long:      `st3llar is a CLI tool that helps users to quickly run their method functions.`,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
		HiddenDefaultCmd:  true,
	},
	PreRun: func(cmd *cobra.Command, args []string) {},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("viper, credential file: ", Vp.GetString(constant.FlagCredential.ValStr()))
		cmd.Usage()
	},
}

// Execute adds all child cobra to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the Root.
func Execute() {
	err := Root.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	//# 对于macOS 64位 GOOS=darwin GOARCH=amd64 去建立 -o mycli-macos ./path/to/package
	//# 对于Linux 64位 GOOS=linux GOARCH=amd64 去建立 -o mycli-linux ./path/to/package
	// This part below is configured for the root and all of its subcommands,
	// so, should not be put in the `PreRun`
	Root.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})
	Root.SetHelpCommand(&cobra.Command{
		Use:    "completion",
		Hidden: true,
	})

	home := config.Home()

	cfg, _ := findConfig(home)
	setupViper(cfg)
	bindViper(home)
	slog.Info(fmt.Sprintf("using config path: %s", Vp.ConfigFileUsed()))

	// TODO: remove the input type of version command
	//Root.PersistentFlags().StringP("version", "v", "", "only run when this command is called directly")

	flagCred := constant.FlagCredential.ValStr()
	Root.PersistentFlags().StringP(flagCred, "c",
		Vp.GetString(flagCred), "the credential file for the command about to be executed")
	Vp.BindPFlag(flagCred, Root.PersistentFlags().Lookup(flagCred))
}

// setupViper viper initialization
func setupViper(cfg *config.St3llarConfig) {
	logger := slog.Default()
	switch cfg.LogLevel {
	case constant.Debug.ValStr():
		slog.SetLogLoggerLevel(slog.LevelDebug)
	case constant.Warn.ValStr():
		slog.SetLogLoggerLevel(slog.LevelWarn)
	case constant.Error.ValStr():
		slog.SetLogLoggerLevel(slog.LevelError)
	default:
		slog.SetLogLoggerLevel(slog.LevelInfo)
	}

	Vp = viper.NewWithOptions(
		viper.EnvKeyReplacer(strings.NewReplacer(".", "_")),
		viper.WithLogger(logger),
		//viper.KeyDelimiter("::"),
	)
}

// bindViper bind viper
func bindViper(home string) {
	Vp.AddConfigPath(home)
	Vp.SetConfigType(constant.ConfigType.ValStr())
	Vp.SetConfigName(constant.ConfigName.ValStr())

	//viper.SetDefault("author", "v3nooom@outlook.com")
	//viper.SetDefault("license", "apache 2.0")

	Vp.AutomaticEnv()

	cobra.CheckErr(Vp.ReadInConfig())
}

// findConfig checks the configuration file
func findConfig(home string) (*config.St3llarConfig, string) {
	cfgPath := filepath.Join(home, constant.ConfigName.ValStr())

	if isExists(cfgPath) {
		cfg, err := config.ReadConfig(cfgPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			os.Exit(1)
		}

		return cfg, cfgPath
	}

	cfg := config.Build(
		config.WithDefault(),
		config.WithCredential(filepath.Join(home, constant.CredentialName.ValStr())),
	)

	if err := config.WriteConfig(cfg, cfgPath); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	return cfg, cfgPath
}

func isExists(path string) bool {
	_, err := os.Stat(path)

	return err == nil
}
