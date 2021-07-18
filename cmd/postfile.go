package cmd

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/marjanbalazs/gofileclient/v2/util"
	"github.com/spf13/cobra"
)

var folderId string
var description string
var password string
var tags string
var expire string

func init() {
	rootCmd.AddCommand(uploadCmd)
	uploadCmd.Flags().StringVarP(&folderId, "folder", "f", "", "ID of the folder to upload to")
	uploadCmd.Flags().StringVarP(&description, "descripion", "d", "", "Description for a newly created folder, ignored if folderId is specified")
	uploadCmd.Flags().StringVarP(&password, "password", "p", "", "Password for a newly created folder, ignored if folderID specified")
	uploadCmd.Flags().StringVarP(&tags, "tags", "t", "", "Tags for a newly created folder, ignored if folderID specified")
	uploadCmd.Flags().StringVarP(&expire, "expire", "e", "", "Expiration timestamp for a newly created folder, ignored if folderID specified")
}

var uploadCmd = &cobra.Command{
	Use:   "upload [server] [filepath] [filename]",
	Short: "Upload a file to the specified server",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		godotenv.Load()
		token := os.Getenv("gofileApiToken")
		server := args[0]
		filePath := args[1]
		fileName := args[2]

		fileHandle, err := os.Open(filePath)

		// Create buffer for upload
		buf := new(bytes.Buffer)
		w := multipart.NewWriter(buf)

		// Load file
		fw, err := w.CreateFormFile("file", fileName)
		_, err = io.Copy(fw, fileHandle)
		if err != nil {
			fmt.Println(err)
			return
		}

		if len(folderId) > 0 {
			folderField, err := w.CreateFormField("folderId")
			folderField.Write([]byte(folderId))
			if err != nil {
				fmt.Println(err)
				return
			}

			tokenField, err := w.CreateFormField("token")
			tokenField.Write([]byte(token))
			if err != nil {
				fmt.Println(err)
				return
			}
		} else {
			if len(description) > 0 {
				descriptionField, err := w.CreateFormField("description")
				descriptionField.Write([]byte(description))
				if err != nil {
					fmt.Println(err)
					return
				}
			}
			if len(password) > 0 {
				passwordField, err := w.CreateFormField("password")
				passwordField.Write([]byte(password))
				if err != nil {
					fmt.Println(err)
					return
				}
			}
			if len(tags) > 0 {
				tagsField, err := w.CreateFormField("tags")
				tagsField.Write([]byte(folderId))
				if err != nil {
					fmt.Println(err)
					return
				}
			}
			if len(expire) > 0 {
				expireField, err := w.CreateFormField("expire")
				expireField.Write([]byte(folderId))
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		}

		w.Close()

		url := fmt.Sprintf("https://%s.gofile.io/uploadFile", server)
		req, err := http.NewRequest("POST", url, buf)
		req.Header.Set("Content-Type", w.FormDataContentType())

		if err != nil {
			fmt.Println(err)
			return
		}

		client := &http.Client{}
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
