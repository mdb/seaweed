package seaweed

const apiBase = "http://magicseaweed.com/api/"

func spotEP(c *Client, spotID string) string {
	return concat([]string{
		apiBase,
		c.APIKey,
		"/forecast/?spot_id=",
		spotID,
	})
}
