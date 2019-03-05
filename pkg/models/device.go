package models

import (
	"errors"
)

// Device model
// swagger:model Device
type Device struct {
	ID            string `json:"id"`
	VersionID     string `json:"versionId"`
	Version       string `json:"version"`
	AppID         string `json:"appId"`
	DeviceID      string `json:"deviceId"`
	DeviceVersion string `json:"deviceVersion"`
	DeviceType    string `json:"deviceType"`
}

// ValidateInitBody validates the properties of an init
// request and returns an error if any of them are missing
func (d *Device) ValidateInitBody() error {
	if d.Version == "" || d.AppID == "" || d.DeviceID == "" {
		return errors.New("version, appId and deviceId fields can't be empty")
	}

	return nil
}
