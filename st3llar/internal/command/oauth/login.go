package oauth

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"syscall"

	rootCMD "github.com/v3nooom/st3llar/internal/command"
	"github.com/v3nooom/st3llar/internal/config"
	"github.com/v3nooom/st3llar/internal/constant"
	"github.com/v3nooom/st3llar/internal/util"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"golang.org/x/term"
	"gopkg.in/yaml.v3"
)

type Credential struct {
	EndPoint     string `yaml:"endpoint"`
	Account      string `yaml:"account"`
	Token        string `yaml:"token"`
	RefreshToken string `yaml:"refresh-token"`
}

// login represents the login command
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

	login.Flags().StringP("organization", "O", "", "Name of the organization")
	login.Flags().StringP("account", "A", "", "Name of the account")

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

	// TODO: tidy the code, move the request to a separate function/package
	// 	Currently, the request is in the loginFunc, which is not a good practice
	// 	It's a verification/experimentation code
	// make a request to the supplier server
	// if the request is successful, store the token and refresh token in the credential file
	var endpoint = "http://st3llar-alb-365211.us-east-2.elb.amazonaws.com"

	client := &http.Client{}
	req, err := http.NewRequest(
		"POST",
		endpoint+"/lambda/register/v3nooom@outlook.com?k1=v1v1v1v1&k2=v2v2v2v2",
		bytes.NewBuffer([]byte(`{"account":"v3nooom.account","organization":"v3nooom.org","password":"111.pwd"}`)))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("---> Response from supplier: \n%#v\n", string(body))

	// new credential file
	cred := Credential{
		EndPoint:     endpoint,
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
	err = os.WriteFile(viper.GetString(constant.FlagCredential.ValStr()), credYaml, 0666)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error writing credentials to file")
		os.Exit(1)
	}

	// sync to configuration file
	cfg := config.Build(
		config.WithLogLevel(viper.GetString(constant.FlagLogLevel.ValStr())),
		config.WithEnvPrefix(viper.GetString(constant.FlagEnvPrefix.ValStr())),
		config.WithEnvironment(viper.GetString(constant.FlagEnvironment.ValStr())),
		config.WithOrganization(viper.GetString(constant.FlagOrganization.ValStr())),
		config.WithCredential(viper.GetString(constant.FlagCredential.ValStr())),
	)

	if err := config.WriteConfig(cfg, config.DefaultPath()); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
