package seaweed

func (c *Client) Tomorrow() (Day, error) {
	return Day{
		Date: "foo",
	}, nil
}
