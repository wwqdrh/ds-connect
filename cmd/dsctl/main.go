package main

import (
	"github.com/spf13/cobra"
	"github.com/wwqdrh/ds-connect/pkg/ds/command"
	"github.com/wwqdrh/ds-connect/pkg/ds/command/general"
	"github.com/wwqdrh/logger"
)

var (
	version = "dev"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:     "dsctl",
		Version: version,
		Short:   "A utility tool to help you work with docker swarm dev environment more efficiently",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
		Example: "dsctl <command> [command options]",
	}

	rootCmd.AddCommand(command.NewConnectCommand())
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
	rootCmd.SilenceUsage = true
	rootCmd.SilenceErrors = true

	// process will hang here
	if err := rootCmd.Execute(); err != nil {
		logger.DefaultLogger.Error("Exit: " + err.Error())
	}
	general.CleanupWorkspace()
}
