package models

import (
	"errors"

	"github.com/google/uuid"
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

// NewDevice returns a new Device model
func NewDevice(versionID string, version string, appID string, deviceID string, deviceVersion string, deviceType string) *Device {
	return &Device{
		ID:            uuid.New().String(),
		VersionID:     versionID,
		Version:       version,
		AppID:         appID,
		DeviceID:      deviceID,
		DeviceVersion: deviceVersion,
		DeviceType:    deviceType,
	}
}

// ValidateInitBody validates the properties of an init
// request and returns an error if any of them are missing
func (d *Device) ValidateInitBody() error {
	if d.Version == "" || d.AppID == "" || d.DeviceID == "" {
		return errors.New("version, appId and deviceId fields can't be empty")
	}

	return nil
}
