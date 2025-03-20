package misc

import (
	"github.com/Jacobamv/golang_arch_template/pkg/bootstrap/http/misc/response"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Options(
	response.Module,
)
