package models

import (
	"github.com/aerogear/mobile-security-service/pkg/helpers"
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
	AppName              string   `json:"appName,omitempty"`
}

// NewDevice returns a new Device model
func NewDevice(sdkInfo *Device, app *App, version *Version) *Device {
	dev := new(Device)
	dev.ID = helpers.GetUUID()
	dev.VersionID = version.ID
	dev.Version = version.Version
	dev.AppID = app.ID
	dev.DeviceVersion = sdkInfo.DeviceVersion
	dev.DeviceType = sdkInfo.DeviceType
	return dev
}

//TODO it is not part of model, it should be in the router and or handler since it is from this layer

//// ValidateInitBody validates the properties of an init
//// request and returns an error if any of them are missing
//func (d *Device) ValidateInitBody() error {
//	if d.Version == "" || d.AppID == "" || d.DeviceID == "" {
//		return errors.New("version, appId and deviceId fields can't be empty")
//	}
//
//	return nil
//}
