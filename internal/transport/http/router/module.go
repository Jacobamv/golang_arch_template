package router

import "go.uber.org/fx"

// Module ...
var Module = fx.Provide(NewRouter)
