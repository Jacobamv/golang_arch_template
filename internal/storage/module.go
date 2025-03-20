package storage

import (
	"github.com/Jacobamv/golang_arch_template/internal/storage/example"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Options(
	example.Module,
)
