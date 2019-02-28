package models

// InitClient model
// swagger:model Init
type InitClient struct {
	ID              string `json:"id"`
	Version         string `json:"version"`
	AppID           string `json:"appId"`
	Disabled        bool   `json:"disabled"`
	DisabledMessage string `json:"disabledMessage,omitempty"`
}
