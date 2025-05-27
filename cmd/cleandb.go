package cmd

import (
	"github.com/obanlatomiwa/go-broadcast-server/database"

	"github.com/spf13/cobra"
)

// cleandbCmd represents the cleandb command
var cleandbCmd = &cobra.Command{
	Use:   "cleandb",
	Short: "deletes all records of the broadcast server from the database",

	Run: func(cmd *cobra.Command, args []string) {
		database.CleanDatabaseData()
	},
}

func init() {
	broadcastCmd.AddCommand(cleandbCmd)
}
