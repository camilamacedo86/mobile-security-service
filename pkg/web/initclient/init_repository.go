package initclient

import (
	"github.com/aerogear/mobile-security-service/pkg/models"
)

// Repository represent the initclient repository contract
type Repository interface {
	GetVersionByAppIDAndVersion(appID string, versionNumber string) (*models.Version, error)
	GetDeviceByDeviceIDAndAppID(deviceID string, appID string) (*models.Device, error)
	GetAppByAppID(appID string) (*models.App, error)
	UpsertVersion(version *models.Version) (*models.Version, error)
}
