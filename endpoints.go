package seaweed

const apiBase = "http://magicseaweed.com/api/"

func spotEp(c *Client, spotId string) string {
	return concat([]string{
		apiBase,
		c.ApiKey,
		"/forecast/?spot_id=",
		spotId,
	})
}
