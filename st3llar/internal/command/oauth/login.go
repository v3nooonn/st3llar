package oauth

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	rootCMD "github.com/v3nooom/st3llar/internal/command"
	"github.com/v3nooom/st3llar/internal/util"
	"golang.org/x/term"
	"os"
	"syscall"
)

var (
	organization string
	username     string
	password     string
	outputPath   string

	//env          = "dev"
	//credentialPath = "./credential.json"
)

type Credential struct {
	Organization string `json:"organization"`
	Username     string `json:"username"`
	Environment  string `json:"env"`
	Token        string `json:"token"`
}

// login represents the sign-in command
var login = &cobra.Command{
	Use:   "login",
	Short: "Login to the Stellar auto-task.",
	Long: `Login to the Stellar auto-action by the pre-established credentials:
Organization, Account and Password

Available credentials are stored locally in the default path: ./credentials.json, once login successfully.
And you can use --credentials to specify the path of the credentials file.`,
	Args:      cobra.OnlyValidArgs,
	ValidArgs: []string{"organization", "username", "env"},
	PreRun:    util.PreRunBindFlags,
	Run:       loginFunc,
}

func init() {
	rootCMD.Root.AddCommand(login)

	login.Flags().StringVarP(&outputPath, "output", "o", "", "the output destination of the current command")
	login.Flags().StringVarP(&organization, "organization", "O", "", "Name of the organization")
	login.Flags().StringVarP(&username, "username", "U", "", "Name of the account")

	if err := login.MarkFlagRequired("organization"); err != nil {
		return
	}
	if err := login.MarkFlagRequired("username"); err != nil {
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

	password = string(passwordBytes)

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

	//cred := Credential{
	//	Organization: viper.GetString(constant.Organization.ValStr()),
	//	Username:     username,
	//	Environment:  viper.GetString(constant.Environment.ValStr()),
	//}
	//
	//// Marshal struct to JSON
	////credJSON, err := json.Marshal(cred)
	//credJSON, err := json.MarshalIndent(cred, "", "  ")
	//if err != nil {
	//	fmt.Fprintln(os.Stderr, "Error marshaling credentials to JSON")
	//	os.Exit(1)
	//}
	//
	//// Write JSON to file
	//err = os.WriteFile(credentialPath, credJSON, 0644)
	//if err != nil {
	//	fmt.Fprintln(os.Stderr, "Error writing credentials to file")
	//	os.Exit(1)
	//}
}
