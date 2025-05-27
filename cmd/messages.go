package cmd

import (
	"fmt"
	"github.com/obanlatomiwa/go-broadcast-server/database"
	"github.com/spf13/cobra"
)

// messagesCmd represents the messages command
var messagesCmd = &cobra.Command{
	Use:   "messages",
	Short: "All the messages from all clients on the server",

	Run: func(cmd *cobra.Command, args []string) {
		messages := database.GetAllMessages()
		fmt.Println("All historical messages from the server...")
		fmt.Println("-------------------------------------------------------------------------------------")
		for _, message := range messages {
			formattedTime := message.Date.Format("02-01-2006")
			text := fmt.Sprintf(" Date: %s ID: %s Text: %s", formattedTime, message.ClientId, message.Text)
			fmt.Println(text)
		}
	},
}

func init() {
	broadcastCmd.AddCommand(messagesCmd)

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// messagesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
