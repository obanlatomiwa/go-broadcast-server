package cmd

import (
	broadCastServer "github.com/obanlatomiwa/go-broadcast-server/websocket"
	"github.com/spf13/cobra"
)

// clientListCmd represents the clientList command
var clientListCmd = &cobra.Command{
	Use:   "clientList",
	Short: "A brief description of your command",

	Run: func(cmd *cobra.Command, args []string) {
		broadCastServer.GetAllClients()
	},
}

func init() {
	broadcastCmd.AddCommand(clientListCmd)

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clientListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
