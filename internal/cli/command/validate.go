package command

import (
	"log/slog"
	"os"

	"github.com/ARTM2000/parrot/internal/core"
	"github.com/spf13/cobra"
)

var validateCommand = &cobra.Command{
	Use:   "validate",
	Short: "validate config file",
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("config")
		config := core.LoadConfig(path)

		if !config.IsValid() {
			os.Exit(1)
		}

		slog.Info("config file is valid")
	},
}

func ValidateCommand() *cobra.Command {
	validateCommand.Flags().StringP("config", "c", "", "path to config file")
	validateCommand.MarkFlagRequired("config")

	return validateCommand
}
