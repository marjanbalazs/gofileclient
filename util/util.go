package util

import (
	"encoding/json"
	"fmt"
)

func PrintResponse(body []byte) {
	var resp map[string]interface{}
	jsonErr := json.Unmarshal(body, &resp)
	if jsonErr != nil {
		fmt.Println("Failed to decode json:", jsonErr)
	} else {
		status, err := json.MarshalIndent(resp["status"], "", "  ")
		if err != nil {
			fmt.Println("Reading status failed", err)
		}
		fmt.Println("Response:", string(status))
		data, err := json.MarshalIndent(resp["data"], "", "  ")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(data))
	}
}
