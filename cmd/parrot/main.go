package main

import (
	"log/slog"
	"os"

	"github.com/ARTM2000/parrot/internal/cli"
)

func main() {
	c := cli.NewComposer().GetCommand()

	if err := c.Execute(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
