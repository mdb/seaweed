package seaweed

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func getForecast(c *Client, spotID string, responseStruct interface{}) error {
	url := fmt.Sprintf("http://magicseaweed.com/api/%s/forecast/?spot_id=%s", c.APIKey, spotID)
	body, err := doRequest(c, url, responseStruct)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, responseStruct)
	if err != nil {
		return err
	}

	return nil
}

func doRequest(c *Client, url string, responseStruct interface{}) (json []byte, er error) {
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}

	bodyContents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("%s returned HTTP status code %d", url, resp.StatusCode)
	}

	c.Log.Debugf("url=%s http_status=%d response_body=%s", url, resp.StatusCode, string(bodyContents))

	return bodyContents, err
}

func matchDays(f []Forecast, match int) []Forecast {
	matched := []Forecast{}

	for _, each := range f {
		if time.Unix(each.LocalTimestamp, 0).UTC().Day() == match {
			matched = append(matched, each)
		}
	}

	return matched
}
