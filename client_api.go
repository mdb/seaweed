package seaweed

import "time"

// Forecast fetches the full, multi-day forecast for a given spot.
func (c *Client) Forecast(spot string) ([]Forecast, error) {
	forecasts := []Forecast{}
	err := getForecast(c, spot, &forecasts)
	if err != nil {
		return forecasts, err
	}

	return forecasts, nil
}

// Today fetches the today's forecast for a given spot.
func (c *Client) Today(spot string) ([]Forecast, error) {
	today := time.Now().Day()
	forecasts, err := c.Forecast(spot)
	if err != nil {
		return []Forecast{}, err
	}

	return matchDays(forecasts, today), nil
}

// Tomorrow fetches tomorrow's forecast for a given spot.
func (c *Client) Tomorrow(spot string) ([]Forecast, error) {
	tomorrowDate := time.Now().Day() + 1
	forecasts, err := c.Forecast(spot)
	if err != nil {
		return []Forecast{}, err
	}

	return matchDays(forecasts, tomorrowDate), nil
}

// Weekend fetches the weekend's forecast for a given spot.
func (c *Client) Weekend(spot string) ([]Forecast, error) {
	forecasts, err := c.Forecast(spot)
	if err != nil {
		return []Forecast{}, err
	}

	return matchWeekendDays(forecasts), nil
}
