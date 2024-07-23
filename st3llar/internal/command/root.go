package command

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/v3nooom/st3llar/internal/config"
	"github.com/v3nooom/st3llar/internal/constant"
	"github.com/v3nooom/st3llar/internal/util"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var version = "v0.0.1"

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
	PreRun: util.PreRunBindFlags,
	Run:    rootFunc,
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
	initConfig()

	//# 对于macOS 64位 GOOS=darwin GOARCH=amd64 去建立 -o mycli-macos ./path/to/package
	//# 对于Linux 64位 GOOS=linux GOARCH=amd64 去建立 -o mycli-linux ./path/to/package
	Root.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})
	Root.SetHelpCommand(&cobra.Command{
		Use:    "completion",
		Hidden: true,
	})

	Root.SetVersionTemplate(`Version: {{.Version}}`)
	Root.Version = version

	flagCred := constant.FlagCredential.ValStr()
	Root.PersistentFlags().StringP(
		flagCred,
		"c",
		viper.GetString(flagCred),
		"the credential file for the command about to be executed")

	flagEnv := constant.FlagEnvironment.ValStr()
	Root.PersistentFlags().StringP(
		flagEnv,
		"e",
		viper.GetString(flagEnv),
		"the execution environment")
}

func initConfig() {
	cfg, _ := findConfig()
	setupViper(cfg)
	slog.Info(fmt.Sprintf("using config path: %s", viper.ConfigFileUsed()))
}

// findConfig checks the configuration file
func findConfig() (*config.St3llarConfig, string) {
	home := config.Home()

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

// setupViper viper setup
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

	viper.NewWithOptions(
		viper.EnvKeyReplacer(strings.NewReplacer(".", "_")),
		viper.WithLogger(logger),
		//viper.KeyDelimiter("::"),
	)

	viper.AddConfigPath(config.Home())
	viper.SetConfigType(constant.ConfigType.ValStr())
	viper.SetConfigName(constant.ConfigName.ValStr())

	//viper.SetDefault("author", "v3nooom@outlook.com")
	//viper.SetDefault("license", "apache 2.0")

	viper.AutomaticEnv()

	cobra.CheckErr(viper.ReadInConfig())
}

func rootFunc(cmd *cobra.Command, args []string) {
	fmt.Println()
	fmt.Println("Root Func:")
	fmt.Println("----> viper settings:")
	for k, v := range viper.AllSettings() {
		fmt.Printf("%v: %v\n", k, v)
	}
	fmt.Println("----> args:")
	for _, v := range args {
		fmt.Printf("%v\n", v)
	}

	fmt.Println("----> flags:")
	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		fmt.Printf("flag.Name: %v, flag.Value: %v\n", flag.Name, flag.Value)
	})

	cmd.Usage()
}
