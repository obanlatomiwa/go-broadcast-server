package cmd

import (
	"github.com/obanlatomiwa/go-broadcast-server/database"
	broadCastServer "github.com/obanlatomiwa/go-broadcast-server/websocket"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start the broadcast server",

	Run: func(cmd *cobra.Command, args []string) {
		// set up DB
		database.InitialiseDatabase()
		broadCastServer.InitiateBroadCast()
	},
}

func init() {
	broadcastCmd.AddCommand(startCmd)
}
