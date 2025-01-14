package cli

import (
	"fmt"

	"github.com/ARTM2000/parrot/internal/cli/command"
	"github.com/spf13/cobra"
)

func registerCommands(c Composer) {
	// - validate config files
	c.registerCommand(command.ValidateCommand())

	// - serve command
	// 		- watch response files
	// 		- serve response files
	c.registerCommand(command.ServeCommand())
}

// ------------------------------------------

type Composer interface {
	GetCommand() *cobra.Command
	registerCommand(cmd *cobra.Command)
}

type composer struct {
	composerCommand *cobra.Command
}

func (c *composer) GetCommand() *cobra.Command {
	return c.composerCommand
}

func (c *composer) registerCommand(cmd *cobra.Command) {
	c.composerCommand.AddCommand(cmd)
}

func NewComposer() Composer {
	c := &composer{
		composerCommand: &cobra.Command{
			Use:   "",
			Short: "Manager command for parrot project",
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("run with --help, -h for guidance")
			},
		},
	}
	registerCommands(c)

	return c
}
