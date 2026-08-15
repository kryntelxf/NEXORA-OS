package nexora

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/router"
)

func TestHealthEndpoint(t *testing.T) {
	app := pocketbase.New()

	// Register plugin
	Register(app)

	// Create test router
	r := router.New()

	// Create serve event with router
	serveEvent := &core.ServeEvent{
		Router: r,
	}

	// Trigger the OnServe hook to register routes
	handlers := app.OnServe().Handlers()
	for _, h := range handlers {
		if err := h.Func(serveEvent); err != nil {
			t.Fatalf("Failed to execute hook: %v", err)
		}
	}

	// Create test request
	req := httptest.NewRequest("GET", "/api/nexora/health", nil)
	w := httptest.NewRecorder()

	// Serve request
	r.ServeHTTP(w, req)

	// Assert response status
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	// Assert response body
	expectedBody := `{"status":"ok","version":"0.1.0","service":"NEXORA-OS"}`
	if w.Body.String() != expectedBody+"\n" {
		t.Errorf("Expected body %s, got %s", expectedBody, w.Body.String())
	}
}
