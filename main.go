package spotifyprivateapi

type Client struct {
	Debug bool // false
}

func NewClient(options Client) *Client {

	client := &Client{
		Debug: options.Debug || false,
	}
	LogInfo("Successfully created client")
	return client
}
