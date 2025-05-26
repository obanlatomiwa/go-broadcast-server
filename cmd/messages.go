package cmd

import (
	broadCastServer "github.com/obanlatomiwa/go-broadcast-server/websocket"

	"github.com/spf13/cobra"
)

// messagesCmd represents the messages command
var messagesCmd = &cobra.Command{
	Use:   "messages",
	Short: "All the messages from all clients on the server",

	Run: func(cmd *cobra.Command, args []string) {
		broadCastServer.GetAllClients()
	},
}

func init() {
	broadcastCmd.AddCommand(messagesCmd)

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// messagesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
