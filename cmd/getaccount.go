package cmd

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/marjanbalazs/gofileclient/v2/util"
	"github.com/spf13/cobra"
)

var FullDetail bool

func init() {
	rootCmd.AddCommand(accountCmd)
	accountCmd.Flags().BoolVarP(&FullDetail, "full", "f", false, "full account detail")
}

var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Get accound details",
	Run: func(cmd *cobra.Command, args []string) {
		godotenv.Load()
		key := os.Getenv("gofileApiToken")

		base, err := url.Parse("https://api.gofile.io/")
		if err != nil {
			return
		}
		base.Path += "getAccountDetails"

		params := url.Values{}
		if FullDetail {
			params.Add("allDetails", strconv.FormatBool(FullDetail))
		}
		params.Add("token", key)

		base.RawQuery = params.Encode()

		res, err := http.Get(base.String())
		if err != nil {
			fmt.Println("Failed to get account detail")
			return
		}
		body, err := io.ReadAll(res.Body)
		defer res.Body.Close()

		util.PrintResponse(body)
	},
}
