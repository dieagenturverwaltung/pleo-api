package webhook

import "time"

type Info struct {
	CreatedAt     time.Time      `json:"createdAt"`
	EndpointUrl   string         `json:"endpointUrl"`
	EventTypes    string         `json:"eventTypes"`
	Id            string         `json:"id"`
	Status        string         `json:"status"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	CustomHeaders map[string]any `json:"customHeaders"`
	DeletedAt     time.Time      `json:"deletedAt"`
	DeletedBy     string         `json:"deletedBy"`
	EndpointAuth  *EndpointAuth  `json:"endpointAuth,omitempty"`
}

type EndpointAuth struct {
	Credentials *EndpointAuthCredentials `json:"credentials,omitempty"`
	Type        string                   `json:"type,omitempty"`
}

type EndpointAuthCredentials struct {
	Password string `json:"password,omitempty"`
	Token    string `json:"token,omitempty"`
	Username string `json:"username,omitempty"`
}
