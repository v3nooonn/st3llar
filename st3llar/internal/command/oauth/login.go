package oauth

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	rootCMD "github.com/v3nooom/st3llar/internal/command"
	"github.com/v3nooom/st3llar/internal/config"
	"github.com/v3nooom/st3llar/internal/constant"
	"github.com/v3nooom/st3llar/internal/util"
	"golang.org/x/term"
	"gopkg.in/yaml.v3"
	"os"
	"syscall"
)

var (
	organization string
	account      string
	//password     string

	//env          = "dev"
	//credentialPath = "./credential.json"
)

type Credential struct {
	Account      string `json:"account"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

// login represents the sign-in command
var login = &cobra.Command{
	Use:   "login",
	Short: "Login to the Stellar auto-task.",
	Long: `Login to the Stellar auto-action by the pre-established credentials:
Organization, Account and Password

Available credential file is stored locally in the default path: $HOME/.st3llar-credentials.
- using --credential to specify the custom credential file path
- using --environment to specify the environment stored in the configuration file`,
	Args:      cobra.OnlyValidArgs,
	ValidArgs: []string{"organization", "account", "env"},
	PreRunE: func(cmd *cobra.Command, args []string) error {

		util.PreRunBindFlags(cmd, args)
		return nil
	},
	Run: loginFunc,
}

func init() {
	rootCMD.Root.AddCommand(login)

	login.Flags().StringVarP(&organization, "organization", "O", "", "Name of the organization")
	login.Flags().StringVarP(&account, "account", "A", "", "Name of the account")

	if err := login.MarkFlagRequired("organization"); err != nil {
		return
	}
	if err := login.MarkFlagRequired("account"); err != nil {
		return
	}
}

func loginFunc(cmd *cobra.Command, args []string) {
	// TODO: error handling
	fmt.Print("Password: ")
	passwordBytes, err := term.ReadPassword(syscall.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading password")
		os.Exit(1)
		return
	}

	password := string(passwordBytes)
	fmt.Printf("input password: %v\n", password)

	fmt.Println()
	fmt.Println("Login Func:")
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

	// new credential file
	cred := Credential{
		Account:      viper.GetString("account"),
		Token:        "JWT_TOKEN",
		RefreshToken: "JWT_REFRESH_TOKEN",
	}

	credYaml, err := yaml.Marshal(cred)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error marshaling credentials to YAML")
		os.Exit(1)
	}

	// Write YAML to file
	err = os.WriteFile(viper.GetString(constant.FlagCredential.ValStr()), credYaml, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error writing credentials to file")
		os.Exit(1)
	}

	// sync to configuration file
	cfgPath := config.Home() + "/.st3llar"
	cfg := config.Build(
		config.WithLogLevel(viper.GetString(constant.FlagLogLevel.ValStr())),
		config.WithEnvPrefix(viper.GetString(constant.FlagEnvPrefix.ValStr())),
		config.WithEnvironment(viper.GetString(constant.FlagEnvironment.ValStr())),
		config.WithOrganization(viper.GetString(constant.FlagOrganization.ValStr())),
		config.WithCredential(viper.GetString(constant.FlagCredential.ValStr())),
	)

	if err := config.WriteConfig(cfg, cfgPath); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
