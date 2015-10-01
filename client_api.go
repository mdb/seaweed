package seaweed

func (c *Client) Forecast(spot string) ([]Forecast, error) {
	forecasts := []Forecast{}
	err := request(c, spotEp(c, spot), &forecasts)
	if err != nil {
		panic(err)
	}

	return forecasts, nil
}
