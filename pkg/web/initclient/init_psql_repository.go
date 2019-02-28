package initclient

import (
	"database/sql"

	log "github.com/sirupsen/logrus"

	"github.com/aerogear/mobile-security-service/pkg/models"
)

type (
	initPostgreSQLRepository struct {
		db *sql.DB
	}
)

// NewPostgreSQLRepository creates a new instance of initPostgreSQLRepository
func NewPostgreSQLRepository(db *sql.DB) Repository {
	return &initPostgreSQLRepository{db}
}

// InitClientApp retrieves all apps from the database
func (a *initPostgreSQLRepository) GetDeviceByDeviceID(deviceID string) (*models.Device, error) {
	query := `
	SELECT id,version_id,app_id,device_id,device_type,device_version
	FROM device
	WHERE device_id = $1;`

	row := a.db.QueryRow(query, deviceID)

	var device models.Device
	if err := row.Scan(&device.ID, device.VersionID, device.AppID, device.DeviceID, device.DeviceType, device.DeviceVersion); err != nil {
		log.Error(err)
		switch err {
		case sql.ErrNoRows:
			return nil, models.ErrNotFound
		default:
			return nil, models.ErrDatabaseError
		}
	}

	return &device, nil
}

// InitClientApp retrieves all apps from the database
func (a *initPostgreSQLRepository) GetVersionByAppIDAndVersion(appID string, versionNumber string) (*models.Version, error) {
	version := models.Version{}
	var disabledMessage sql.NullString
	var lastLaunchedAt sql.NullString

	sqlStatement := `
	SELECT v.id,v.version,v.app_id, v.disabled, v.disabled_message, v.num_of_app_launches, v.last_launched_at
	FROM version as v
	WHERE v.app_id = $1 AND v.version = $2;`

	err := a.db.QueryRow(sqlStatement, appID, versionNumber).Scan(&version.ID, &version.Version, &version.AppID, &version.Disabled, &disabledMessage, &version.NumOfAppLaunches, &lastLaunchedAt)

	version.DisabledMessage = disabledMessage.String
	version.LastLaunchedAt = lastLaunchedAt.String

	if err != nil {
		log.Error(err)
		if err == sql.ErrNoRows {
			return nil, models.ErrNotFound
		}
		return nil, models.ErrInternalServerError
	}

	return &version, nil
}

// GetDeviceByDeviceIDAndAppID returns a device by its device ID and app ID
func (a *initPostgreSQLRepository) GetDeviceByDeviceIDAndAppID(deviceID string, appID string) (*models.Device, error) {
	device := models.Device{}

	sqlStatement := `
		SELECT d.id, d.version_id, d.app_id, d.device_id, d.device_type, d.device_version
		FROM device as d
		WHERE d.device_id = $1 AND d.app_id = $2;`

	if err := a.db.QueryRow(sqlStatement, deviceID, appID).
		Scan(&device.ID, &device.VersionID, &device.AppID, &device.DeviceID, &device.DeviceType, &device.DeviceVersion); err != nil {

		log.Error(err)
		if err == sql.ErrNoRows {
			return nil, models.ErrNotFound
		}
		return nil, models.ErrInternalServerError
	}

	return &device, nil
}

// GetDeviceByVersionAndAppID returns a device by its version number and app ID
func (a *initPostgreSQLRepository) GetDeviceByVersionAndAppID(version string, appID string) (*models.Device, error) {
	device := models.Device{}

	sqlStatement := `
		SELECT d.id, d.version_id, d.app_id, d.device_id, d.device_type, d.device_version
		FROM device as d
		WHERE d.app_id = $1 AND d.device_version = $2;`

	if err := a.db.QueryRow(sqlStatement, appID, version).
		Scan(&device.ID, &device.VersionID, &device.AppID, &device.DeviceID, &device.DeviceType, &device.DeviceVersion); err != nil {

		log.Error(err)
		if err == sql.ErrNoRows {
			return nil, models.ErrNotFound
		}
		return nil, models.ErrInternalServerError
	}

	return &device, nil
}

// GetAppByID retrieves an app by id from the database
func (a *initPostgreSQLRepository) GetAppByAppID(appID string) (*models.App, error) {
	app := models.App{}

	sqlStatment := `SELECT id,app_id,app_name FROM app WHERE app_id=$1 AND deleted_at IS NULL;`
	err := a.db.QueryRow(sqlStatment, appID).Scan(&app.ID, &app.AppID, &app.AppName)

	if err != nil {
		log.Error(err)
		if err == sql.ErrNoRows {
			return nil, models.ErrNotFound
		}
		return nil, models.ErrInternalServerError
	}

	return &app, nil

}

// InsertVersionOrUpdateNumOfAppLaunches creates a new version row
// or increments the num_of_app_launches counter if the version already exists
func (a *initPostgreSQLRepository) InsertVersionOrUpdateNumOfAppLaunches(version *models.Version) error {
	sqlStatement := `
		INSERT INTO version(id, version, app_id, disabled, disabled_message, num_of_app_launches, last_launched_at)
		VALUES($1, $2, $3, $4, $5, $6, NOW())
		ON CONFLICT (id)
		DO UPDATE
		SET num_of_app_launches = $6,
		last_launched_at = NOW();`

	_, err := a.db.Exec(sqlStatement, version.ID, version.Version, version.AppID, version.Disabled, version.DisabledMessage, version.NumOfAppLaunches)

	if err != nil {
		log.Error(err)
		return models.ErrDatabaseError
	}

	return nil
}

// CreateDevice creates a new device row in the device table
func (a *initPostgreSQLRepository) CreateDevice(device *models.Device) error {
	sqlStatement := `
		INSERT INTO device(id,version_id,app_id,device_id,device_type,device_version)
		VALUES($1, $2, $3, $4, $5, $6);`

	_, err := a.db.Exec(sqlStatement, device.ID, device.VersionID, device.AppID, device.DeviceID, device.DeviceType, device.DeviceVersion)

	if err != nil {
		log.Error(err)
		return models.ErrDatabaseError
	}

	return nil
}
