package marketplace

type InstallationStatus string

const (
	InstallationStatusActivated          InstallationStatus = "ACTIVATED"
	InstallationStatusInactive           InstallationStatus = "INACTIVE"
	InstallationStatusAuthorized         InstallationStatus = "AUTHORIZED"
	InstallationStatusPendingIntegration InstallationStatus = "PENDING_INTEGRATION"
	InstallationStatusError              InstallationStatus = "ERROR"
)

type InstallationErrorCode string

const (
	InstallationErrorCodeNotEntitled InstallationErrorCode = "NOT_ENTITLED"
	InstallationErrorCodeAuthError   InstallationErrorCode = "AUTH_ERROR"
)

type ExternalClientAuthMethod string

const (
	ExternalClientAuthMethodOAuth2 ExternalClientAuthMethod = "OAUTH2"
	ExternalClientAuthMethodAPIKey ExternalClientAuthMethod = "API_KEY"
	ExternalClientAuthMethodNone   ExternalClientAuthMethod = "NONE"
)

type CreateClientInstallationRequest struct {
	Metadata map[string]any     `json:"metadata,omitempty"`
	Status   InstallationStatus `json:"status,omitempty"`
}

type UpdateInstallationRequest struct {
	ErrorCode *InstallationErrorCode `json:"errorCode,omitempty"`
	Metadata  map[string]any         `json:"metadata,omitempty"`
	Status    InstallationStatus     `json:"status"`
}

type InstallationResponse struct {
	ApplicationID string             `json:"applicationId"`
	CreatedAt     string             `json:"createdAt"`
	ErrorCode     string             `json:"errorCode,omitempty"`
	ID            string             `json:"id"`
	Metadata      map[string]any     `json:"metadata"`
	Resource      string             `json:"resource"`
	Status        InstallationStatus `json:"status"`
	UpdatedAt     string             `json:"updatedAt"`
}

type ExternalClientModel struct {
	ApplicationName string                   `json:"applicationName,omitempty"`
	AuthMethod      ExternalClientAuthMethod `json:"authMethod"`
	ClientID        string                   `json:"clientId"`
	Scope           string                   `json:"scope"`
	SubjectURN      string                   `json:"subjectURN,omitempty"`
}

type ExternalClientModelDetails struct {
	Data ExternalClientModel `json:"data"`
}
