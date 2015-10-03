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
	today := time.Now().Day()
	forecasts, err := c.Forecast(spot)
	if err != nil {
		return []Forecast{}, err
	}

	return matchDays(forecasts, today), nil
}

func (c *Client) Tomorrow(spot string) ([]Forecast, error) {
	tomorrowDate := time.Now().Day() + 1
	forecasts, err := c.Forecast(spot)
	if err != nil {
		return []Forecast{}, err
	}

	return matchDays(forecasts, tomorrowDate), nil
}
