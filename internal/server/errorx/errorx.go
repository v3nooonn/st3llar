package errorx

import (
	"fmt"
	"os"
)

func ErrorMapping(err error) {
	if err != nil {
		n, _ := fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(n)
	}
}
