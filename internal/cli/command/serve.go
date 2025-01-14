package command

import (
	"github.com/ARTM2000/parrot/internal/core"
	"github.com/spf13/cobra"
)

var serveCommand = &cobra.Command{
	Use:   "run",
	Short: "run server within config file",
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("config")

		_ = core.RunServer(path)
	},
}

func ServeCommand() *cobra.Command {
	serveCommand.Flags().StringP("config", "c", "config.yml", "path to config file")

	return serveCommand
}
