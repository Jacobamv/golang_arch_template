package http

import (
	"github.com/Jacobamv/golang_arch_template/pkg/bootstrap/http/middlewares"
	"github.com/Jacobamv/golang_arch_template/pkg/bootstrap/http/misc"
	"github.com/Jacobamv/golang_arch_template/pkg/bootstrap/http/server"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Options(
	middlewares.Module,
	misc.Module,

	server.ServerModule,
)
