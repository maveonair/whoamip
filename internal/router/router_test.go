package router_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/maveonair/whoamip/internal/router"
)

func TestIndexHandler(t *testing.T) {
	expectedIpAddr := "27.133.16.173"

	testSuite := []struct {
		RemoteAddr    string
		XForwardedFor string
	}{
		{
			RemoteAddr:    fmt.Sprintf("%s:12345", expectedIpAddr),
			XForwardedFor: "",
		},
		{
			RemoteAddr:    "",
			XForwardedFor: expectedIpAddr,
		},
	}

	for _, test := range testSuite {
		req, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		req.RemoteAddr = test.RemoteAddr
		req.Header.Set("X-Forwarded-For", test.XForwardedFor)

		res := httptest.NewRecorder()
		router := router.NewRouter()
		router.ServeHTTP(res, req)

		if res.Code != http.StatusOK {
			t.Errorf("Expected HTTP status 200, found %d", res.Code)
		}

		var result map[string]interface{}
		json.NewDecoder(res.Body).Decode(&result)

		if result["ip"] != expectedIpAddr {
			t.Errorf("Expected IP, %s found %s", expectedIpAddr, result["ip"])
		}
	}
}
