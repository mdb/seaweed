package seaweed

func (c *Client) Tomorrow(spot string) (Day, error) {
	forecast := []Day{}
	err := request(c, spotEp(c, spot), &forecast)
	if err != nil {
		panic(err)
	}

	return forecast[0], nil
}
