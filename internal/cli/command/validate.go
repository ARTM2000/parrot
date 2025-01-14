package command

import (
	"github.com/ARTM2000/parrot/internal/core"
	"github.com/spf13/cobra"
	"log/slog"
)

var validateCommand = &cobra.Command{
	Use:   "validate",
	Short: "validate config file",
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("config")
		_ = core.LoadConfig(path)

		slog.Info("config file is valid")
	},
}

func ValidateCommand() *cobra.Command {
	validateCommand.Flags().StringP("config", "c", "config.yml", "path to config file")

	return validateCommand
}
