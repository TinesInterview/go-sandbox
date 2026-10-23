package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	cases := []struct {
		name       string
		method     string
		path       string
		wantStatus int
	}{
		{name: "greets on the documented route", method: http.MethodGet, path: "/api/hello", wantStatus: http.StatusOK},
		{name: "preflight is allowed", method: http.MethodOptions, path: "/api/hello", wantStatus: http.StatusNoContent},
		{name: "unknown route is not found", method: http.MethodGet, path: "/api/nope", wantStatus: http.StatusNotFound},
	}
	handler := withCORS(newMux())
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			handler.ServeHTTP(rec, httptest.NewRequest(tc.method, tc.path, nil))

			if rec.Code != tc.wantStatus {
				t.Fatalf("status = %d, want %d", rec.Code, tc.wantStatus)
			}
			if got := rec.Header().Get("Access-Control-Allow-Origin"); got != "*" {
				t.Errorf("Access-Control-Allow-Origin = %q, want %q", got, "*")
			}
			if tc.wantStatus != http.StatusOK {
				return
			}

			var got helloResponse
			if err := json.NewDecoder(rec.Body).Decode(&got); err != nil {
				t.Fatalf("decode body: %v", err)
			}
			if got.Message != "Welcome" {
				t.Errorf("message = %q, want %q", got.Message, "Welcome")
			}
		})
	}
}
