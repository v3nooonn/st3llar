package main

import (
	"github.com/v3nooom/stellar-auto-task/internal/cobra/command"
	_ "github.com/v3nooom/stellar-auto-task/internal/cobra/command/auth"
	_ "github.com/v3nooom/stellar-auto-task/internal/cobra/command/overview"
)

func main() {
	command.Execute()
}
