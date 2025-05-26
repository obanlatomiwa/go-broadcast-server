package cmd

import (
	broadCastServer "github.com/obanlatomiwa/go-broadcast-server/websocket"
	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop the broadcast server",

	Run: func(cmd *cobra.Command, args []string) {
		broadCastServer.StopBroadCast()
	},
}

func init() {
	broadcastCmd.AddCommand(stopCmd)
}
