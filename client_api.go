package seaweed

import "time"

func (c *Client) Forecast(spot string) ([]Forecast, error) {
	forecasts := []Forecast{}
	err := request(c, spotEp(c, spot), &forecasts)
	if err != nil {
		return forecasts, err
	}

	return forecasts, nil
}

func (c *Client) Today(spot string) ([]Forecast, error) {
	now := time.Now().Day()
	today := []Forecast{}
	forecasts, err := c.Forecast(spot)
	if err != nil {
		return today, err
	}

	for _, each := range forecasts {
		if time.Unix(each.LocalTimestamp, 0).Day() == now {
			today = append(today, each)
		}
	}

	return today, nil
}
