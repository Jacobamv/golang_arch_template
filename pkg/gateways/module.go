package gateways

import (
	"github.com/Jacobamv/golang_arch_template/pkg/gateways/example"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Options(
	example.Module,
)
