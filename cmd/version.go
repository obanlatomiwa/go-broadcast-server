package cmd

import (
	"fmt"
	"github.com/obanlatomiwa/go-broadcast-server/utils"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "The latest version of the broadcast server",

	Run: func(cmd *cobra.Command, args []string) {
		version := fmt.Sprintf("Broadcast Server Version %s", utils.GetValueFromConfigFile("APP_VERSION"))
		fmt.Println(version)
	},
}

func init() {
	broadcastCmd.AddCommand(versionCmd)
}
