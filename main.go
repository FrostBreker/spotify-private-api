package spotifyprivateapi

// The above code defines a struct type called "Client" with a boolean field called "Debug".
// @property {bool} Debug - The `Debug` property is a boolean value that indicates whether debugging
// mode is enabled or not. If `Debug` is set to `true`, it means that debugging mode is enabled. If
// `Debug` is set to `false`, it means that debugging mode is disabled.
type Client struct {
	Debug bool // false
}

// The NewClient function creates a new client object with the specified options.
func NewClient(options Client) *Client {

	client := &Client{
		Debug: options.Debug || false,
	}
	LogInfo("Successfully created client")
	return client
}
