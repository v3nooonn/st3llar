package oauth

import (
	"fmt"
	"os"
	"syscall"

	rootCMD "github.com/v3nooom/st3llar/internal/command"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var (
	organization string
	username     string
	password     string

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
	Run:       loginFunc,
}

func init() {
	rootCMD.Root.AddCommand(login)

	login.Flags().StringVarP(&organization, "organization", "O", "", "Name of the organization")
	login.Flags().StringVarP(&username, "username", "U", "", "Name of the account")
	login.Flags().StringVarP(&username, "env", "", "", "Environment of the CLI")
	rootCMD.Vp.BindPFlag("env", rootCMD.Root.PersistentFlags().Lookup("debug"))

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

	fmt.Println("----------------- loginFunc() ------------------")
	//fmt.Printf("in viper: %v\n", rootCMD.Vp.GetString("environment"))
	//fmt.Printf("login args: %v\n", args)
	//fmt.Printf("login output: %v\n", credentialPath)
	//fmt.Printf("login env: %v\n", env)
	//fmt.Printf("login env from viper: %v\n", rootCMD.Vp.GetString("env"))

	//cred := Credential{
	//	Organization: organization,
	//	Username:     username,
	//	Environment:  env,
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
