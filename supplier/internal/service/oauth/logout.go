package oauth

import (
	"fmt"

	repr "github.com/v3nooom/st3llar/supplier/internal/api/representation"
)

func Logout(input repr.Login) {
	fmt.Printf("---> OAuth service Layer, Logout %#v\n", input)
}
