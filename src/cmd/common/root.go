package common

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var errorCmd = log.New(os.Stderr, "Initialize error: ", 0)

var RootCmd = &cobra.Command{
	Use:   "currency",
	Short: "Currency server",
}

func Execute() {
	RootCmd.CompletionOptions.DisableDefaultCmd = true
	if err := RootCmd.Execute(); err != nil {
		os.Exit(0)
	}
}
