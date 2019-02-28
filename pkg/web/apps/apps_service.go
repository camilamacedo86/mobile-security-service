package apps

import (
	"github.com/aerogear/mobile-security-service/pkg/helpers"
	"github.com/aerogear/mobile-security-service/pkg/models"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type (
	// Service defines the interface methods to be used
	Service interface {
		GetApps() (*[]models.App, error)
		GetActiveAppByID(ID string) (*models.App, error)
		UpdateAppVersions(versions []models.Version) error
		DisableAllAppVersionsByAppID(id string, message string) error
		UnbindingAppByAppID(appID string) error
		BindingAppByApp(appId, name string) error
		InitClientApp(deviceInfo *models.Device) (*models.InitClient, error)
	}

	appsService struct {
		repository Repository
	}
)

// NewService instantiates this service
func NewService(repository Repository) Service {
	return &appsService{
		repository: repository,
	}
}

// GetApps retrieves the list of apps from the repository
func (a *appsService) GetApps() (*[]models.App, error) {
	apps, err := a.repository.GetApps()

	// Check for errors and return the appropriate error to the handler
	if err != nil {
		return nil, err
	}

	return apps, nil
}

// GetActiveAppByID retrieves app by id from the repository
func (a *appsService) GetActiveAppByID(id string) (*models.App, error) {

	app, err := a.repository.GetActiveAppByID(id)

	if err != nil {
		return nil, err
	}

	deployedVersions, err := a.repository.GetAppVersionsByAppID(app.AppID)
	if err != nil {
		return nil, err
	}
	app.DeployedVersions = deployedVersions

	return app, nil
}

// GetApps retrieves the list of apps from the repository
func (a *appsService) UpdateAppVersions(versions []models.Version) error {
	err := a.repository.UpdateAppVersions(versions)

	// Check for errors and return the appropriate error to the handler
	if err != nil {
		return err
	}

	return nil
}

// Update all versions
func (a *appsService) DisableAllAppVersionsByAppID(id string, message string) error {

	// get the app id to send it to the re
	app, err := a.repository.GetActiveAppByID(id)

	if err != nil {
		return err
	}

	return a.repository.DisableAllAppVersionsByAppID(app.AppID, message)
}

func (a *appsService) UnbindingAppByAppID(appID string) error {
	err := a.repository.DeleteAppByAppID(appID)
	if err != nil {
		return err
	}
	return nil
}

func (a *appsService) BindingAppByApp(appId, name string) error {

	// Check if it exist
	app, err := a.repository.GetAppByAppID(appId)

	// If it is new then create an app
	if err != nil && err == models.ErrNotFound {
		id := helpers.GetUUID()
		return a.repository.CreateApp(id, appId, name)
	}

	if err != nil {
		return err
	}

	// if is deleted so just reactive the existent app
	return a.repository.UnDeleteAppByAppID(app.AppID)
}

// InitClientApp retrieves the list of apps from the repository
func (a *appsService) InitClientApp(deviceInfo *models.Device) (*models.InitClient, error) {
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
	initData := models.InitClient{
		ID:              version.ID,
		Version:         version.Version,
		AppID:           version.AppID,
		Disabled:        version.Disabled,
		DisabledMessage: version.DisabledMessage,
	}

	return &initData, nil
}
