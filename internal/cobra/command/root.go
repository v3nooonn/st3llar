package command

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	//userLicense string

	// Root represents the base command when called without any subcommands
	Root = &cobra.Command{
		Use:   "st3llar",
		Short: "Stellar auto-task supporter CLI: st3llar",
		Long: `st3llar is a command line tool that helps users to quickly run their method functions. 
It provides a set of cobra that can be used to interact with the server endpoints, whose main goal is:
1. Helping users to register their method functions to AWS Lambda.
2. Triggering/Scheduling the Lambda
3. Monitoring the Lambda execution status and results.
4. Lambdas management`,
		//Uncomment the following line if your bare application
		//has an action associated with it:
		//Run: func(cmd *cobra.Command, args []string) {
		//	fmt.Println("st3llar called")
		//},
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("todo error in main CMD")
		},
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
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	Root.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.stellar-auto-task.yaml)")
	//Root.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
	//Root.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
	Root.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	viper.BindPFlag("author", Root.PersistentFlags().Lookup("author"))
	viper.BindPFlag("useViper", Root.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "v3nooom@outlook.com")
	viper.SetDefault("license", "apache 2.0")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	Root.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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
