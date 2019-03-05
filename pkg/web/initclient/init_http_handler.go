package initclient

import (
	"errors"
	"github.com/aerogear/mobile-security-service/pkg/web/apps"
	"net/http"

	"github.com/aerogear/mobile-security-service/pkg/httperrors"

	"github.com/aerogear/mobile-security-service/pkg/models"
	"github.com/labstack/echo"
)

type (
	// HTTPHandler instance
	HTTPHandler struct {
		appsService apps.Service
	}
)

// NewHTTPHandler returns a new instance of app.Handler
func NewHTTPHandler(e *echo.Echo, a apps.Service) *HTTPHandler {
	return &HTTPHandler{
		appsService: a,
	}
}

// InitClientApp stores device information and returns if the app version is disabled
func (h *HTTPHandler) InitClientApp(c echo.Context) error {
	deviceInfo := new(models.Device)

	if err := c.Bind(deviceInfo); err != nil {
		return err
	}

	// Check the request body is valid
	if err := validateInitRequestBody(deviceInfo); err != nil {
		return httperrors.BadRequest(c, err.Error())
	}

	initResponse, err := h.appsService.InitClientApp(deviceInfo)

	// If no app has been found in the database, return a bad request to the client
	if err == models.ErrNotFound {
		return httperrors.BadRequest(c, "No bound app found for the sent App ID")
	}

	if err != nil {
		return httperrors.GetHTTPResponseFromErr(c, err)
	}

	return c.JSON(http.StatusOK, initResponse)
}

// Request body must have version, appId and deviceId
func validateInitRequestBody(deviceInfo *models.Device) error {
	if deviceInfo.Version == "" || deviceInfo.AppID == "" || deviceInfo.DeviceID == "" {
		return errors.New("version, appId and deviceId fields can't be empty")
	}
	return nil
}
