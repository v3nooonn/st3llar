package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string

	configPath string
	outputPath string
)

// Root represents the base command when called without any subcommands
var Root = &cobra.Command{
	Use: "st3llar",
	//Args: cobra.NoArgs,
	Args:      cobra.OnlyValidArgs,
	ValidArgs: []string{"login", "logout"},
	// func(cmd *Command, args []string) err
	// NoArgs(cmd *Command, args []string) err
	Short: "Stellar AutoTask CLI: st3llar",
	Long:  `st3llar is a CLI tool that helps users to quickly run their method functions.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("configPath: %v\n", configPath)
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

	// Customized HelpFunc and UsageFunc
	//Root.SetHelpCommand(&cobra.Command{})
	//Root.SetUsageFunc(func(cmd *cobra.Command) err {
	//	fmt.Println("This is the help message for the root command")
	//	return nil
	//})
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".stellar-auto-task" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".st3llar")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
