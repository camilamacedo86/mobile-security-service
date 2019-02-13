package initclient

import (
	"github.com/aerogear/mobile-security-service/pkg/models"
	uuid "github.com/satori/go.uuid"
)

type (
	// Service defines the interface methods to be used
	Service interface {
		InitClientApp(deviceInfo *models.Device) (*models.Version, error)
	}

	initService struct {
		repository Repository
	}
)

// NewService instantiates this service
func NewService(repository Repository) Service {
	return &initService{
		repository: repository,
	}
}

// InitClientApp retrieves the list of apps from the repository
func (a *initService) InitClientApp(deviceInfo *models.Device) (*models.Version, error) {
	// Check if the app exists in the database for the sent app_id
	_, err := a.repository.GetAppByAppID(deviceInfo.AppID)

	// If an app doesn't exist, an error is returned and bubbled back to the delivery layer
	if err != nil {
		return nil, err
	}

	version := &models.Version{}
	version, err = a.repository.GetVersionByAppIDAndVersion(deviceInfo.AppID, deviceInfo.Version)

	// If any error other Not Found error occurred, return
	if err != nil && err != models.ErrNotFound {
		return nil, err
	}

	// If the version does not exist, create it
	if err == models.ErrNotFound {
		// Create new uuid for our new app version
		versionUUID, err := uuid.NewV4()
		if err != nil {
			return nil, err
		}
		version = &models.Version{
			ID:               versionUUID.String(),
			Version:          deviceInfo.Version,
			AppID:            deviceInfo.AppID,
			NumOfAppLaunches: 0,
		}
	}
	// Increment the version App Launches
	version.NumOfAppLaunches++

	version, err = a.repository.UpsertVersion(version)

	if err != nil {
		return nil, err
	}

	// device, err := a.repository.GetDeviceByDeviceIDAndAppID(deviceInfo.DeviceID, deviceInfo.AppID)

	// Check app exists first and if not, return a 400
	// 1. Get version from version + appId,
	// 2. Get device from deviceId + versionId + AppId
	// 3. Increment appLaunches
	// 4. Create or Update the device row in the table

	// TODO
	// Increment app launches for this appId and version
	// Get disabled and disabledMessage for this version

	// Check device table to see if this device exists for the current version and appId
	// 		If yes - Do nothing
	// 		If no - Check if it exists for this appId + deviceId, update the versionId to the current version.
	// 						If it doesn't exist at all, create a new row

	return version, nil
}
