package nexora

import (
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/hook"
)

// Register registers the Nexora plugin with the PocketBase app
func Register(app core.App) {
	// Register health check endpoint
	app.OnServe().Bind(&hook.Handler[*core.ServeEvent]{
		Func: func(e *core.ServeEvent) error {
			e.Router.GET("/api/nexora/health", func(c *core.RequestEvent) error {
				return c.JSON(200, map[string]interface{}{
					"status":  "ok",
					"version": "0.1.0",
					"service": "NEXORA-OS",
				})
			})
			return e.Next()
		},
		Priority: 100, // Run before static file handler
	})
}
