package initclient

import (
	"github.com/aerogear/mobile-security-service/pkg/models"
)

// Repository represent the initclient repository contract
type Repository interface {
	GetVersionByAppIDAndVersion(appID string, versionNumber string) (*models.Version, error)
	GetDeviceByDeviceIDAndAppID(deviceID string, appID string) (*models.Device, error)
	GetDeviceByVersionAndAppID(versionID string, appID string) (*models.Device, error)
	GetAppByAppID(appID string) (*models.App, error)
	InsertVersionOrUpdateNumOfAppLaunches(version *models.Version) error
	CreateDevice(device *models.Device) error
}
