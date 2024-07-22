package main

import (
	"github.com/v3nooom/st3llar/internal/command"
	_ "github.com/v3nooom/st3llar/internal/command/general"
	_ "github.com/v3nooom/st3llar/internal/command/oauth"
)

func main() {
	command.Execute()
}
