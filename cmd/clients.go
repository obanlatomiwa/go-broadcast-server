package cmd

import (
	"fmt"
	"github.com/obanlatomiwa/go-broadcast-server/database"
	"github.com/spf13/cobra"
)

// clientCmd represents the clientList command
var clientCmd = &cobra.Command{
	Use:   "clients",
	Short: "lists all the historical clients",

	Run: func(cmd *cobra.Command, args []string) {
		clients := database.GetAllClients()
		fmt.Println("All registered clients from the server...")
		fmt.Println("-------------------------------------------------------------------------------------")
		for _, client := range clients {
			text := fmt.Sprintf("Client ID: %s Status: %s", client.ClientId, client.Status)
			fmt.Println(text)
		}
	},
}

func init() {
	broadcastCmd.AddCommand(clientCmd)

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clientCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
