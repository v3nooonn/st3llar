package main

import (
	"github.com/v3nooom/st3llar/internal/command"
	_ "github.com/v3nooom/st3llar/internal/command/oauth"
	//_ "github.com/v3nooom/st3llar/internal/cobra/command/overview"
)

func main() {
	command.Execute()
}
