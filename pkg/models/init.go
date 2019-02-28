package models

// Init model
// swagger:model Init
type Init struct {
	ID              string `json:"id"`
	Version         string `json:"version"`
	AppID           string `json:"appId"`
	Disabled        bool   `json:"disabled"`
	DisabledMessage string `json:"disabledMessage,omitempty"`
}
