package cmd

import (
	"fmt"
	"github.com/obanlatomiwa/go-broadcast-server/database"
	"github.com/spf13/cobra"
)

var online bool

// clientCmd represents the clientList command
var clientCmd = &cobra.Command{
	Use:   "clients",
	Short: "lists all the historical clients",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		clients := database.GetAllClients()

		if online {
			fmt.Println("All online registered clients from the server...")
			fmt.Println("-------------------------------------------------------------------------------------")
			for _, client := range clients {
				if client.Status == "ONLINE" {
					text := fmt.Sprintf("Client ID: %s Status: %s", client.ClientId, client.Status)
					fmt.Println(text)
				}
			}
		} else {
			fmt.Println("All registered clients from the server...")
			fmt.Println("-------------------------------------------------------------------------------------")
			for _, client := range clients {
				text := fmt.Sprintf("Client ID: %s Status: %s", client.ClientId, client.Status)
				fmt.Println(text)
			}
		}

	},
}

func init() {
	clientCmd.Flags().BoolVarP(&online, "online", "o", false, "lists all the online clients")
	broadcastCmd.AddCommand(clientCmd)

}
