package schema

type RequestHeaders struct {
	BearerToken string `json:"bearerToken"`
	DeviceId    string `json:"deviceId"`
	Platform    string `json:"platform"`
	Source      string `json:"source"`
}

type RequestPayload struct {
}
