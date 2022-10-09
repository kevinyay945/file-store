package pcloud

type Client struct {
	accessToken string
}

func (c *Client) SetAccessToken(s string) {
	c.accessToken = s
}

func (c *Client) CheckWriteAuthorization() (err error) {
	return nil
}

func New() Client {
	return Client{}
}
