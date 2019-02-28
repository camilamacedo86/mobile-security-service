package initclient

import (
	"github.com/aerogear/mobile-security-service/pkg/models"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type (
	// Service defines the interface methods to be used
	Service interface {
		InitClientApp(deviceInfo *models.Device) (*models.Init, error)
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
func (a *initService) InitClientApp(deviceInfo *models.Device) (*models.Init, error) {
	// Check if the app exists in the database for the sent app_id
	if _, err := a.repository.GetAppByAppID(deviceInfo.AppID); err != nil {
		return nil, err
	}

	version, err := a.repository.GetVersionByAppIDAndVersion(deviceInfo.AppID, deviceInfo.Version)

	// If any error other Not Found error occurred, return
	if err != nil && err != models.ErrNotFound {
		return nil, err
	}

	// If the version does not exist, create it
	if err == models.ErrNotFound {
		// Create new uuid for our new app version
		versionUUID := uuid.New()

		version = &models.Version{
			ID:               versionUUID.String(),
			Version:          deviceInfo.Version,
			AppID:            deviceInfo.AppID,
			NumOfAppLaunches: 0,
		}
	}

	// Increment the version App Launches
	version.NumOfAppLaunches++

	// Update the existing version or create a new one
	if err := a.repository.InsertVersionOrUpdateNumOfAppLaunches(version); err != nil {
		log.Error(err)
		return nil, err
	}

	// If we can't find the device by version and app ID
	if _, err = a.repository.GetDeviceByVersionAndAppID(version.Version, deviceInfo.AppID); err == models.ErrNotFound {

		// If we can't find the device by device ID and app ID
		if _, err := a.repository.GetDeviceByDeviceIDAndAppID(deviceInfo.DeviceID, deviceInfo.AppID); err == models.ErrNotFound {

			id := uuid.New()

			// Build a new device to save to the database
			device := models.Device{
				ID:            id.String(),
				VersionID:     version.ID,
				Version:       version.Version,
				AppID:         deviceInfo.AppID,
				DeviceID:      deviceInfo.DeviceID,
				DeviceVersion: deviceInfo.DeviceVersion,
				DeviceType:    deviceInfo.DeviceType,
			}

			// Could not insert the device
			if err := a.repository.CreateDevice(&device); err != nil {
				log.Error(err)
			}
		}
	}

	// Build a model for the init data to return
	initData := models.Init{
		ID:              version.ID,
		Version:         version.Version,
		AppID:           version.AppID,
		Disabled:        version.Disabled,
		DisabledMessage: version.DisabledMessage,
	}

	return &initData, nil
}
