package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aknt",
	Short: "anki-notion-syncer is a simple CLI for syncing Notion tables as decks in Anki",
	Long:  "anki-notion-syncer is a simple CLI for syncing Notion tables as decks in Anki",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
