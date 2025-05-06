package common

import (
	"currency/internal/config"
	"os"

	"github.com/spf13/cobra"
)

var (
	cfgFile string
	debug   bool

	runCmd = &cobra.Command{
		Use:   "run",
		Short: "Run currency server",
		Run: func(cmd *cobra.Command, args []string) {
			initConfig()
		},
	}
)

func init() {
	runCmd.Flags().StringVarP(&cfgFile, "config", "c", "", "Path to config file")
	runCmd.MarkFlagRequired("config")
	runCmd.Flags().BoolVarP(&debug, "debug", "d", false, "Start in debug mode")
	RootCmd.AddCommand(runCmd)
}

func initConfig() {
	err := config.LoadFromFile(&config.Currency, cfgFile)
	if err != nil {
		errorCmd.Println(err)
		os.Exit(0)
	}
	config.ConfigFile = cfgFile
}
