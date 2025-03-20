package main

import (
	"context"

	net_http "net/http"

	"github.com/Jacobamv/golang_arch_template/internal/service"
	"github.com/Jacobamv/golang_arch_template/internal/storage"
	"github.com/Jacobamv/golang_arch_template/internal/transport/http/handlers"
	"github.com/Jacobamv/golang_arch_template/internal/transport/http/router"
	"github.com/Jacobamv/golang_arch_template/pkg/bootstrap/http"
	"github.com/Jacobamv/golang_arch_template/pkg/config"
	"github.com/Jacobamv/golang_arch_template/pkg/databases"
	"github.com/Jacobamv/golang_arch_template/pkg/gateways"
	"github.com/Jacobamv/golang_arch_template/pkg/logger"
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
