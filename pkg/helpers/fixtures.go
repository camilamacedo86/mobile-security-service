package helpers

import (
	"github.com/aerogear/mobile-security-service/pkg/models"
	"github.com/google/uuid"
)

//GetMockUser returns a dummy user
func GetMockUser() *models.User {
	user := &models.User{
		Username: "TestUser",
		Email:    "test@user.com",
	}
	return user
}

// GetMockAppList returns some dummy apps
func GetMockAppList() []models.App {
	apps := []models.App{
		models.App{
			ID:      "7f89ce49-a736-459e-9110-e52d049fc025",
			AppID:   "com.aerogear.mobile_app_one",
			AppName: "Mobile App One",
		},
		models.App{
			ID:      "7f89ce49-a736-459e-9110-e52d049fc026",
			AppID:   "com.aerogear.mobile_app_three",
			AppName: "Mobile App Two",
		},
		models.App{
			ID:      "7f89ce49-a736-459e-9110-e52d049fc027",
			AppID:   "com.aerogear.mobile_app_three",
			AppName: "Mobile App Three",
		},
	}
	return apps
}

//GetMockApp returns a dummy app
func GetMockApp() *models.App {
	versionlist := GetMockAppVersionList()
	app := &models.App{
		ID:               "7f89ce49-a736-459e-9110-e52d049fc027",
		AppID:            "com.aerogear.mobile_app_one",
		AppName:          "Mobile App One",
		DeployedVersions: &versionlist,
	}
	return app
}

// GetMockAppVersionList returns some dummy app versions
func GetMockAppVersionList() []models.Version {
	versions := []models.Version{
		models.Version{
			ID:               "55ebd387-9c68-4137-a367-a12025cc2cdb",
			Version:          "1.0",
			AppID:            "com.aerogear.mobile_app_one",
			DisabledMessage:  "Please contact an administrator",
			Disabled:         false,
			NumOfAppLaunches: 2,
		},
		models.Version{
			ID:               "59ebd387-9c68-4137-a367-a12025cc1cdb",
			Version:          "1.1",
			AppID:            "com.aerogear.mobile_app_one",
			Disabled:         false,
			NumOfAppLaunches: 0,
		},
		models.Version{
			ID:               "59dbd387-9c68-4137-a367-a12025cc2cdb",
			Version:          "1.0",
			AppID:            "com.aerogear.mobile_app_two",
			Disabled:         false,
			NumOfAppLaunches: 0,
		},
	}

	return versions
}

//Get version strut for the disable all.
func GetMockAppVersionForDisableAll() models.Version {
	return models.Version{
		DisabledMessage: "Please contact an administrator",
	}
}

// GetMockDevice returns a mock device
func GetMockDevice() *models.Device {
	return &models.Device{
		ID:            uuid.New().String(),
		VersionID:     uuid.New().String(),
		AppID:         "com.aerogear.testapp",
		DeviceID:      uuid.New().String(),
		DeviceVersion: "8.1",
		DeviceType:    "Android",
	}
}

// GetMockVersion returns a mock version
func GetMockVersion() *models.Version {
	return &models.Version{
		ID:               uuid.New().String(),
		Version:          "1.0",
		AppID:            "com.aerogear.mobile_app_one",
		DisabledMessage:  "Please contact an administrator",
		Disabled:         false,
		NumOfAppLaunches: 10000,
	}
}

// GetMockDevices generates a slice of Devices
func GetMockDevices(number int) []models.Device {
	var devices []models.Device

	for i := 0; i < number; i++ {
		d := GetMockDevice()
		devices = append(devices, *d)
	}

	return devices
}
