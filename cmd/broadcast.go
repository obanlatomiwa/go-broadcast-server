package cmd

import (
	"github.com/spf13/cobra"
)

// broadcastCmd represents the broadcast command
var broadcastCmd = &cobra.Command{
	Use:   "broadcast",
	Short: "The broadcast server base command",
}

func init() {
	rootCmd.AddCommand(broadcastCmd)
}
