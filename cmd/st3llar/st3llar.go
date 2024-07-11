package main

import (
	"github.com/v3nooom/st3llar/internal/cobra/command"
	_ "github.com/v3nooom/st3llar/internal/cobra/command/auth"
	_ "github.com/v3nooom/st3llar/internal/cobra/command/overview"
)

func main() {
	command.Execute()
}
