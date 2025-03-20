package main

import (
	"context"

	net_http "net/http"

	"gitlab.humo.tj/orzu/applications/bridge/internal/service"
	"gitlab.humo.tj/orzu/applications/bridge/internal/storage"
	"gitlab.humo.tj/orzu/applications/bridge/internal/transport/http/handlers"
	"gitlab.humo.tj/orzu/applications/bridge/internal/transport/http/router"
	"gitlab.humo.tj/orzu/applications/bridge/pkg/bootstrap/http"
	"gitlab.humo.tj/orzu/applications/bridge/pkg/config"
	"gitlab.humo.tj/orzu/applications/bridge/pkg/databases"
	"gitlab.humo.tj/orzu/applications/bridge/pkg/gateways"
	"gitlab.humo.tj/orzu/applications/bridge/pkg/logger"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		http.Module,
		config.Module,
		logger.Module,
		databases.PostgresModule,
		service.Module,
		storage.Module,
		handlers.Module,
		router.Module,
		gateways.Module,
		ModuleLifecycleHooks,
	)

	app.Run()
}

var ModuleLifecycleHooks = fx.Invoke(RegisterHooks)

// RegisterHooks ...
func RegisterHooks(lifecycle fx.Lifecycle, server *net_http.Server) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {

			go server.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
}
