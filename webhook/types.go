package webhook

import "time"

type Info struct {
	CreatedAt     time.Time      `json:"createdAt"`
	EndpointUrl   string         `json:"endpointUrl"`
	EventTypes    []string       `json:"eventTypes"`
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

type ActivityOperationType string

const (
	ActivityOperationTypeCreate ActivityOperationType = "CREATE"
	ActivityOperationTypeUpdate ActivityOperationType = "UPDATE"
	ActivityOperationTypeDelete ActivityOperationType = "DELETE"
)

type Activity struct {
	ActorUrn       string    `json:"actorUrn"`
	CreatedAt      time.Time `json:"createdAt"`
	Description    string    `json:"description"`
	Id             string    `json:"id"`
	OperationType  string    `json:"operationType"`
	SubscriptionId string    `json:"subscriptionId"`
}

type Secret struct {
	SecretKey string `json:"secretKey"`
}

type IPayload interface {
	PayloadExport | PayloadVendor
}

// Payload is the payload submitted by pleo to the webhook url
type Payload[T IPayload] struct {
	WebhookId        string `json:"webhook-id"`
	WebhookTimestamp int    `json:"webhook-timestamp"`
	WebhookSignature string `json:"webhook-signature"`
	Payload          struct {
		Data      T      `json:"data"`
		EventId   string `json:"eventId"`
		EventType string `json:"eventType"`
	} `json:"payload"`
}

type PayloadExport struct {
	CompanyId     string    `json:"companyId"`
	CreatedAt     time.Time `json:"createdAt"`
	CreatedBy     string    `json:"createdBy"`
	ExpiresIn     int       `json:"expiresIn"`
	Id            string    `json:"id"`
	NumberOfItems int       `json:"numberOfItems"`
	Status        string    `json:"status"`
}

type PayloadVendor struct {
	Id                    string      `json:"id"`
	CompanyId             string      `json:"companyId"`
	Name                  string      `json:"name"`
	Country               string      `json:"country"`
	RegistrationNumber    string      `json:"registrationNumber"`
	TaxRegistrationNumber string      `json:"taxRegistrationNumber"`
	Code                  interface{} `json:"code"`
	State                 string      `json:"state"`
	DefaultCurrency       string      `json:"defaultCurrency"`
	CreatedBy             string      `json:"createdBy"`
	CreatedAt             time.Time   `json:"createdAt"`
	UpdatedAt             time.Time   `json:"updatedAt"`
	DeletedAt             interface{} `json:"deletedAt"`
}
