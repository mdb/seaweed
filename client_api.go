package seaweed

func (c *Client) Forecast(spot string) ([]Forecast, error) {
	forecasts := []Forecast{}
	err := request(c, spotEp(c, spot), &forecasts)
	if err != nil {
		return forecasts, err
	}

	return forecasts, nil
}
