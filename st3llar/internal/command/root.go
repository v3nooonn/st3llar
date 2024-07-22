package command

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/v3nooom/st3llar/internal/config"
	"github.com/v3nooom/st3llar/internal/constant"

	"github.com/sagikazarmark/slog-shim"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

var (
	Vp *viper.Viper

	outputPath string
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
	PreRun: func(cmd *cobra.Command, args []string) {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		setupViper()
		checkConfig(home)
		bindViper(home)
		fmt.Printf("Using cfg path: %s\n", Vp.ConfigFileUsed())

	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("st3llar root command...")
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
	//cobra.OnInitialize(preRun) // using pre-run instead

	//# 对于macOS 64位 GOOS=darwin GOARCH=amd64 去建立 -o mycli-macos ./path/to/package
	//# 对于Linux 64位 GOOS=linux GOARCH=amd64 去建立 -o mycli-linux ./path/to/package

	// hide default commands in Cobra
	Root.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})
	Root.SetHelpCommand(&cobra.Command{
		Use:    "completion",
		Hidden: true,
	})

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//Root.PersistentFlags().StringVarP(&configPath, "config", "", "", "to specify the path of the preRun file")
	Root.PersistentFlags().StringP("version", "v", "", "only run when this command is called directly")
	Root.PersistentFlags().StringVarP(&outputPath, "output", "o", "", "the output destination of the current command")

	//Root.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Display more verbose output in console output. (default: false)")
	//Vp.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
	//
	//Root.PersistentFlags().BoolVarP(&Debug, "debug", "d", false, "Display debugging output in the console. (default: false)")
	//Vp.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
}

// setupViper viper initialization
func setupViper() {
	Vp = viper.NewWithOptions(
		viper.EnvKeyReplacer(strings.NewReplacer(".", "_")),
		viper.WithLogger(slog.Default()),
		//viper.KeyDelimiter("::"),
	)

}

// bindViper bind viper
func bindViper(home string) {
	Vp.AddConfigPath(home)
	Vp.SetConfigType(constant.ConfigFileType.ValStr())
	Vp.SetConfigName(constant.ConfigFileName.ValStr())

	//viper.SetDefault("author", "v3nooom@outlook.com")
	//viper.SetDefault("license", "apache 2.0")

	Vp.AutomaticEnv()

	cobra.CheckErr(Vp.ReadInConfig())
}

// checkConfig checks the configuration file
func checkConfig(home string) string {
	path := filepath.Join(home, constant.ConfigFileName.ValStr())

	if isExists(path) {
		return path
	}

	cfg := config.Build(
		config.WithDefault(),
		config.WithCredential(home),
	)

	if err := newConfigFile(cfg, path); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	return path
}

func isExists(path string) bool {
	_, err := os.Stat(path)

	return err == nil
}

func newConfigFile(cfg *config.St3llarConfig, path string) error {
	yamlBytes, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("marshalling default config error: %w", err)
	}

	if err := os.WriteFile(path, yamlBytes, 0666); err != nil {
		return fmt.Errorf("writing config file error, %w: %s\n", err, path)
	}

	return nil
}
