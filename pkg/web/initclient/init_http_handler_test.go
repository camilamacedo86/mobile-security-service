package initclient

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"

	"github.com/aerogear/mobile-security-service/pkg/models"

	"github.com/aerogear/mobile-security-service/pkg/web/apps"
	"github.com/labstack/echo"
)

func TestHTTPHandler_InitClientApp(t *testing.T) {
	type fields struct {
		appsService apps.Service
	}
	type args struct {
		device models.Device
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantStatusCode int
	}{
		{
			name: "A 400 Bad Request should be returned when request body is missing device ID",
			args: args{
				device: models.Device{
					ID:            "d1895cc1-28d7-4283-932d-8bcab9e4a469",
					VersionID:     "1a3fbe59-e69c-42c7-b9c0-d7eb7e65c073",
					AppID:         "com.aerogear.testapp",
					DeviceVersion: "1.1",
					Version:       "9.1.1",
					DeviceType:    "Android",
				},
			},
			wantStatusCode: 400,
		},
		{
			name: "A 400 Bad Request should be returned when request body is missing version",
			args: args{
				device: models.Device{
					ID:            "d1895cc1-28d7-4283-932d-8bcab9e4a469",
					DeviceID:      "1a3fbe59-e69c-42c7-b9c0-d7eb7e65c073",
					AppID:         "com.aerogear.testapp",
					DeviceVersion: "1.1",
					DeviceType:    "Android",
				},
			},
			wantStatusCode: 400,
		},
		{
			name: "A 400 Bad Request should be returned when request body is missing app ID",
			args: args{
				device: models.Device{
					ID:            "d1895cc1-28d7-4283-932d-8bcab9e4a469",
					VersionID:     "1a3fbe59-e69c-42c7-b9c0-d7eb7e65c073",
					DeviceVersion: "1.1",
					Version:       "9.1.1",
					DeviceType:    "Android",
				},
			},
			wantStatusCode: 400,
		},
		{
			name: "Expect init data to be returned when valid device is supplied",
			args: args{
				device: models.Device{
					ID:            "d1895cc1-28d7-4283-932d-8bcab9e4a469",
					VersionID:     "1a3fbe59-e69c-42c7-b9c0-d7eb7e65c073",
					DeviceID:      "db511711-95e8-4da2-8e76-1700465ae8ca",
					AppID:         "com.aerogear.testapp",
					DeviceVersion: "1.1",
					Version:       "9.1.1",
					DeviceType:    "Android",
				},
			},
			wantStatusCode: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()

			deviceJSON, _ := json.Marshal(tt.args.device)

			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(deviceJSON)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)
			c.SetPath("/api/init")

			mockedAppService := &apps.ServiceMock{
				InitClientAppFunc: func(device *models.Device) (*models.InitClient, error) {
					id := uuid.New()

					return &models.InitClient{
						ID:              id.String(),
						Version:         tt.args.device.Version,
						AppID:           tt.args.device.AppID,
						Disabled:        true,
						DisabledMessage: "App is disabled",
					}, nil
				},
			}

			handler := NewHTTPHandler(e, mockedAppService)

			if handler.InitClientApp(c); rec.Code != tt.wantStatusCode {
				t.Errorf("HTTPHandler.InitClientApp() statusCode = %v, wantStatusCode %v", rec.Code, tt.wantStatusCode)
			}
		})
	}
}

func trimBody(body string) string {
	return strings.TrimSpace(body)
}
