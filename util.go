package seaweed

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func request(c *Client, url string, responseStruct interface{}) error {
	if LogRequests() {
		fmt.Printf("Request url: \n\t%s\n", url)
	}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return err
	}

	resp, err := c.HttpClient.Do(req)

	bodyContents, err := ioutil.ReadAll(resp.Body)
	if LogRequests() {
		fmt.Printf("Response status: \n\t%d\nresponse body: \n\t%s \n\n", resp.StatusCode, bodyContents)
	}
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bodyContents, responseStruct); err != nil {
		return err
	}
	return nil
}

func concat(arr []string) string {
	var buff bytes.Buffer

	for _, elem := range arr {
		buff.WriteString(elem)
	}

	return buff.String()
}
