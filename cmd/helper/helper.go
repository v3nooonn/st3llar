package main

import (
	"github.com/v3nooom/st3llar-helper/internal/cobra/command"
	_ "github.com/v3nooom/st3llar-helper/internal/cobra/command/auth"
	_ "github.com/v3nooom/st3llar-helper/internal/cobra/command/overview"
)

func main() {
	command.Execute()
}
