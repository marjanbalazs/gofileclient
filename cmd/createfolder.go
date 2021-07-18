package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/marjanbalazs/gofileclient/v2/util"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createFolderCmd)
}

var createFolderCmd = &cobra.Command{
	Use:   "create [parent id] [folder name]",
	Short: "Create a folder",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		godotenv.Load()

		token := os.Getenv("gofileApiToken")
		parentFolder := args[0]
		folderName := args[1]

		url := "https://api.gofile.io/createFolder/"

		payload := strings.NewReader(fmt.Sprintf("parentFolderId=%s&token=%s&folderName=%s", parentFolder, token, folderName))

		client := &http.Client{}
		req, err := http.NewRequest("PUT", url, payload)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		res, err := client.Do(req)

		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		defer res.Body.Close()

		util.PrintResponse(body)
	},
}
