package lambda

import (
	"fmt"

	repr "github.com/v3nooom/st3llar/supplier/internal/api/representation"
)

func Register(input repr.Login) {
	fmt.Printf("---> Lambda service Layer: %#v\n", input)
}
