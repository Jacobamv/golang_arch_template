package service

import (
	"github.com/Jacobamv/golang_arch_template/internal/service/example"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Options(
	example.Module,
)
