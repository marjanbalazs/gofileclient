package cmd

import (
	"fmt"
	"io"
	"net/http"

	"github.com/marjanbalazs/gofileclient/v2/util"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Get the nearest server",
	Run: func(cmd *cobra.Command, args []string) {
		res, err := http.Get("https://api.gofile.io/getServer")
		if err != nil {
			fmt.Println("Failed to get nearest server")
			return
		}
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Failed to get nearest server")
			return
		}
		util.PrintResponse(body)
	},
}
