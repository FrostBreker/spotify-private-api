package responseTypes

type GrantedToken struct {
	Token               string `json:"token"`
	ExpiresAfterSeconds int    `json:"expires_after_seconds"`
	RefreshAfterSeconds int    `json:"refresh_after_seconds"`
}

type ResponseClientId struct {
	ResponseType string       `json:"response_type"`
	GrantedToken GrantedToken `json:"granted_token"`
}

type ResponseClientToken struct {
	AccessToken             string `json:"accessToken"`
	AccessTokenExpirationMs int64  `json:"accessTokenExpirationTimestampMs"`
	IsAnonymous             bool   `json:"isAnonymous"`
	ClientId                string `json:"clientId"`
}
