package oauth

import "fmt"

var HandlerLogin = &Login{}

func init() {
	if HandlerLogin == nil {
		HandlerLogin = &Login{}
	}
}

type Login struct{}

func (l *Login) Login(o, u, p string) error {
	fmt.Printf("organization: %s, username: %s, password: %s\n", o, u, p)

	if o == "57b" && u == "v3nooom" && p == "123123" {
		fmt.Println("Login successfully.")
		return nil
	}

	return fmt.Errorf("invalid credentials")
}
