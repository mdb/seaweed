package seaweed

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

func getForecast(c *Client, url string, responseStruct interface{}) error {
	file := cacheFile(url, c)
	var err error
	var body []byte

	if isCacheStale(file, c) {
		body, err = doRequest(c, url, responseStruct)
		if err != nil {
			return err
		}
		if !DisableCache() {
			if err := writeCache(file, body); err != nil {
				return err
			}
		}
	} else {
		if LogRequests() {
			fmt.Printf("Reading cached forecast file: \n\t%s\n", file)
		}
		body, err = ioutil.ReadFile(file)
	}
	if err := json.Unmarshal(body, responseStruct); err != nil {
		return err
	}
	return nil
}

func doRequest(c *Client, url string, responseStruct interface{}) (json []byte, er error) {
	if LogRequests() {
		fmt.Printf("Request url: \n\t%s\n", url)
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		defer resp.Body.Close()
		return nil, err
	}

	bodyContents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if LogRequests() {
		fmt.Printf("Response status: \n\t%d\nresponse body: \n\t%s \n\n", resp.StatusCode, bodyContents)
	}

	return bodyContents, nil
}

func concat(arr []string) string {
	var buff bytes.Buffer

	for _, elem := range arr {
		buff.WriteString(elem)
	}

	return buff.String()
}

func matchDays(f []Forecast, match int) []Forecast {
	matched := []Forecast{}

	for _, each := range f {
		if time.Unix(each.LocalTimestamp, 0).Day() == match {
			matched = append(matched, each)
		}
	}

	return matched
}

func matchWeekendDays(f []Forecast) []Forecast {
	matched := []Forecast{}

	for _, each := range f {
		if isWeekend(each) {
			matched = append(matched, each)
		}
	}

	return matched
}

func isWeekend(f Forecast) bool {
	day := time.Unix(f.LocalTimestamp, 0).Weekday().String()

	if day == "Saturday" || day == "Sunday" {
		return true
	}
	return false
}

func cacheFile(url string, c *Client) string {
	file := fmt.Sprintf("seaweed_%s", strings.Split(url, "=")[1])
	return path.Join(c.CacheDir, file)
}

func isCacheStale(cacheFile string, c *Client) bool {
	stat, err := os.Stat(cacheFile)

	return os.IsNotExist(err) || time.Since(stat.ModTime()) > c.CacheAge || DisableCache()
}

func writeCache(cacheFile string, json []byte) (err error) {
	return ioutil.WriteFile(cacheFile, json, 0644)
}
