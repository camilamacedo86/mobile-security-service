package initclient

import (
	log "github.com/sirupsen/logrus"
	"database/sql"

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
	device1 := models.Device{
		ID:      "a0874c82-2b7f-11e9-b210-d663bd873d93",
		Version: "1.2.1",
		AppID:   "com.aerogear.app1",
	}

	return &device1, nil
}

// InitClientApp retrieves all apps from the database
func (a *initPostgreSQLRepository) GetVersionByAppIDAndVersion(appID string, versionNumber string) (*models.Version, error) {
	version := models.Version{}
	var disabledMessage sql.NullString

	sqlStatment := `
	SELECT v.id,v.version,v.app_id, v.disabled, v.disabled_message, v.num_of_app_launches
	FROM version as v
	WHERE v.app_id = $1 AND v.version = $2;`
	err := a.db.QueryRow(sqlStatment, appID, versionNumber).Scan(&version.ID, &version.Version, &version.AppID, &version.Disabled, &disabledMessage, &version.NumOfAppLaunches)

	version.DisabledMessage = disabledMessage.String

	if err != nil {
		log.Error(err)
		if err == sql.ErrNoRows {
			return nil, models.ErrNotFound
		}
		return nil, models.ErrInternalServerError
	}

	return &version, nil
}

func (a *initPostgreSQLRepository) GetDeviceByDeviceIDAndAppID(deviceID string, appID string) (*models.Device, error) {
	device := models.Device{}

	sqlStatment := `
	SELECT d.id, d.version_id, d.app_id, d.device_id, d.device_type, d.device_version
	FROM device as d
	WHERE d.device_id = $1 AND d.app_id = $2;`

	err := a.db.QueryRow(sqlStatment, deviceID, appID).Scan(&device.ID, &device.VersionID, &device.AppID, &device.DeviceID, &device.DeviceType, &device.DeviceVersion)

	if err != nil {
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

// UpsertVersion updates an existing version 
// or creates a new one if it does not exist
func (a *initPostgreSQLRepository) UpsertVersion(version *models.Version) (error) {
	// TODO: Add updated_at column to version table
	sqlStatement := `
	INSERT INTO version(id, version, app_id, disabled, disabled_message, num_of_app_launches)
	VALUES($1, $2, $3, $4, $5, $6)
	ON CONFLICT (id)
	DO UPDATE
	SET num_of_app_launches = $6;
	`

	_, err := a.db.Exec(sqlStatement, version.ID, version.Version, version.AppID, version.Disabled, version.DisabledMessage, version.NumOfAppLaunches)

	if err != nil {
		return models.ErrDatabaseError
	}

	return nil
}
