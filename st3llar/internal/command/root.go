package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/v3nooom/st3llar/internal/config"
	"github.com/v3nooom/st3llar/internal/constant"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

var (
	Viper *viper.Viper

	configPath string
	outputPath string
)

// Root represents the base command when called without any subcommands
var Root = &cobra.Command{
	Use:       "st3llar",
	Args:      cobra.OnlyValidArgs,
	ValidArgs: []string{"login", "logout"},
	Short:     "Stellar AutoTask CLI: st3llar",
	Long:      `st3llar is a CLI tool that helps users to quickly run their method functions.`,
	//Run: func(cmd *cobra.Command, args []string) {
	//	//fmt.Printf("configPath: %v\n", configPath)
	//	fmt.Println("calls st3llar root command")
	//},
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
	//initConfig()
	cobra.OnInitialize(initConfig)

	//# 对于macOS 64位 GOOS=darwin GOARCH=amd64 去建立 -o mycli-macos ./path/to/package
	//# 对于Linux 64位 GOOS=linux GOARCH=amd64 去建立 -o mycli-linux ./path/to/package
	//viper.SetDefault("author", "v3nooom@outlook.com")
	//viper.SetDefault("license", "apache 2.0")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	Root.PersistentFlags().StringVarP(&configPath, "config", "", "", "to specify the path of the configuration file")
	Root.PersistentFlags().StringP("version", "v", "", "only run when this command is called directly")
	Root.PersistentFlags().StringVarP(&outputPath, "output", "o", "", "the output destination of the current command")

	//Root.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Display more verbose output in console output. (default: false)")
	//Viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
	//
	//Root.PersistentFlags().BoolVarP(&Debug, "debug", "d", false, "Display debugging output in the console. (default: false)")
	//Viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))

	// Customized HelpFunc and UsageFunc
	//Root.SetHelpCommand(&cobra.Command{})
	//Root.SetUsageFunc(func(cmd *cobra.Command) err {
	//	fmt.Println("This is the help message for the root command")
	//	return nil
	//})
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	fmt.Println("----------------- initConfig() ------------------")
	// Find home/wd directory.
	//currentDir, err := os.Getwd()
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	//cfgFilePath := filepath.Join(home, fmt.Sprintf(".%s.%s", constant.ConfigFileName.ValStr(), constant.ConfigFileType.ValStr()))
	cfgFilePath := filepath.Join(home, constant.ConfigFileName.ValStr())
	fmt.Println("config_file_path", cfgFilePath)

	if isExists(cfgFilePath) {
		fmt.Printf("Using: %s as the configuration file.\n", cfgFilePath)
	} else {
		fmt.Println("File does not exist, here comes the creation")

		sc := config.St3llarConfig{
			Environment:  "dev",
			EnvVarPrefix: "st3llar_",
		}
		yamlBytes, err := yaml.Marshal(&sc)
		if err != nil {
			fmt.Println("Marshalling configuration error: ", err.Error())
		}
		err = os.WriteFile(home+"/.st3llar", yamlBytes, 0666)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error occurred at writing configuration")
			os.Exit(1)
		}
	}

	//v := viper.NewWithOptions(
	//	viper.KeyDelimiter("::"),
	//	viper.EnvKeyReplacer(strings.NewReplacer(".", "_")),
	//	viper.WithLogger(slog.Default()),
	//)
	//
	//v.SetConfigName(constant.ConfigFileName.ValStr())
	//v.SetConfigType(constant.ConfigFileType.ValStr())

	Viper = viper.NewWithOptions()
	Viper.SetConfigFile(cfgFilePath)

	//// Search config in home directory with name ".stellar-auto-task" (without extension).
	//viper.AddConfigPath(currentDir)
	//viper.SetConfigType("yaml")
	//viper.SetConfigName(".st3llar")

	Viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := Viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", Viper.ConfigFileUsed())
	}

	//fmt.Printf("viper found env: %s\n", Viper.GetString("environment"))
	fmt.Println("Config file used:", Viper.ConfigFileUsed())
	fmt.Println("All settings loaded by Viper:")
	for key, value := range Viper.AllSettings() {
		fmt.Printf("%s: %v\n", key, value)
	}
}

func isExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		// Print any error, not just os.IsNotExist
		// Enhanced error logging
		fmt.Printf("Error checking file: %v\n", err)
		return !os.IsNotExist(err)
	}
	return true
}
