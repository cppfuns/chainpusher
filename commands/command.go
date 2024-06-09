package commands

import (
	"fmt"
	"log"
	"os"

	"github.com/chainpusher/chainpusher/config"
	"github.com/spf13/cobra"
)

func NewMonitorCobraCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "monitor",
		Short: "Monitor blockchain data",
		Run: func(cmd *cobra.Command, args []string) {

			var cfg *config.Config

			defer func() {
				if r := recover(); r != nil {
					log.Println("Recovered in f", r)
				}
			}()

			p, err := cmd.Flags().GetString("config")

			if err != nil {
				cfg = &config.Config{}
			} else {
				cfg, err = config.ParseConfigFromYaml(p)
				if err != nil {
					log.Fatalf("failed to parse config: %v", err)
					cfg = &config.Config{}
				}
			}

			cfg.BlockLoggingFile, _ = cmd.Flags().GetString("block-file")

			SetupLogger(cfg)

			monitor := NewMonitorCommand(cfg)
			monitor.Execute()
		},
	}

	cmd.PersistentFlags().StringP("block-file", "b", "", "File to write raw blockchain data to")
	cmd.PersistentFlags().StringP("trx-file", "t", "", "File to write transactions to")

	return cmd
}

func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "chainpusher",
		Short: "A CLI tool for pushing blockchain data",
		Long: "Chainpusher is a CLI tool for pushing blockchain data to a remote server. " +
			"Chainpusher can also monitor blockchain data and push it to a remote server.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.PersistentFlags().String(
		"config",
		"c",
		"config file (default is $HOME/.chainpusher.yaml)",
	)

	return cmd
}

func RunCommand() {

	rootCmd := NewRootCommand()
	monitorCmd := NewMonitorCobraCommand()

	rootCmd.AddCommand(monitorCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
