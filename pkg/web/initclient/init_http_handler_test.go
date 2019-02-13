package initclient

import (
	"strings"
	"testing"
)

// var mockPostgresRepository = NewPostgreSQLRepository()
// var mockService = NewService(mockPostgresRepository)

func TestHTTPHandler_InitClientApp(t *testing.T) {
	// // set up mock context
	// e := echo.New()
	// req := httptest.NewRequest(http.MethodGet, "/", nil)
	// rec := httptest.NewRecorder()
	// c := e.NewContext(req, rec)

	// c.SetPath("/init")
	// h := NewHTTPHandler(e, mockService)

	// _ = h.InitClientApp(c)

	// if rec.Code != http.StatusOK {
	// 	t.Errorf("HTTPHandler.InitClientApp() statusCode = %v, wantCode = %v", rec.Code, http.StatusOK)
	// }

	// expected := `[{"id":1,"appId":"com.aerogear.app1","appName":"app1","numOfDeployedVersions":1,"numOfClients":1,"numOfAppLaunches":1}]`

	// resBody := trimBody(rec.Body.String())

	// if resBody != expected {
	// 	t.Errorf("HTTPHandler.InitClientApp() want = %v, wantCode = %v", expected, resBody)
	// }
}

func trimBody(body string) string {
	return strings.TrimSpace(body)
}
