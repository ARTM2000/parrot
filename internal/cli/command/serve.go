package command

import (
	"os"

	"github.com/ARTM2000/parrot/internal/core"
	"github.com/spf13/cobra"
)

var serveCommand = &cobra.Command{
	Use:   "run",
	Short: "run server within config file",
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("config")
		config := core.LoadConfig(path)

		if !config.IsValid() {
			os.Exit(1)
		}

		core.RunServer(config)
	},
}

func ServeCommand() *cobra.Command {
	serveCommand.Flags().StringP("config", "c", "", "path to config file")
	serveCommand.MarkFlagRequired("config")

	return serveCommand
}
