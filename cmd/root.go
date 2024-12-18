package cmd

import (
	"fmt"
	"goscope/internal/ui"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goscope",
	Short: "GoScope is a lightweight system monitoring tool.",
	Long:  `GoScope monitors CPU, memory, disk, and network usage with a TUI.`,
	Run: func(cmd *cobra.Command, args []string) {
		ui.StartUI()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
