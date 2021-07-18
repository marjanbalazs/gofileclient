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
	rootCmd.AddCommand(setFolderOpsCmd)
}

var setFolderOpsCmd = &cobra.Command{
	Use:   "setfolder [folderid] [option] [value]",
	Short: "Set folder options",
	Long: "Set folder options:\n" +
		"private: true or false\n" +
		"password: password\n" +
		"description: description\n" +
		"expire: Date in unix timestamp\n" +
		"tags: comma separated list of tags\n",
	Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		godotenv.Load()
		token := os.Getenv("gofileApiToken")
		folderId := args[0]
		opt := args[1]
		val := args[2]

		url := "https://api.gofile.io/setFolderOptions"

		payload := strings.NewReader(fmt.Sprintf("folderId=%s&token=%s&option=%s&value=%s", folderId, token, opt, val))

		client := &http.Client{}
		req, err := http.NewRequest("PUT", url, payload)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		if err != nil {
			fmt.Println(err)
			return
		}

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
