package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string

	// Root represents the base command when called without any subcommands
	Root = &cobra.Command{
		Use:   "st3llar",
		Short: "Stellar AutoTask CLI: st3llar",
		Long:  `st3llar is a CLI tool that helps users to quickly run their method functions.`,
		//It provides a set of commands to interact with the server, whose main features are:
		//1. Register functions to AWS Lambda.
		//2. Triggering/Scheduling the Lambda.
		//3. Monitoring the Lambda execution status and results.
		//4. Lambdas management.`,
	}
)

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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// aboutCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// aboutCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	//viper.SetDefault("author", "v3nooom@outlook.com")
	//viper.SetDefault("license", "apache 2.0")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	Root.PersistentFlags().StringP("gfs", "g", "", "used for this command and all subcommands")
	Root.Flags().StringP("lfs", "l", "", "only run when this command is called directly")

	// Customized HelpFunc and UsageFunc
	//Root.SetHelpCommand(&cobra.Command{})
	//Root.SetUsageFunc(func(cmd *cobra.Command) error {
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
