package server

import (
	"fmt"
	"net"
	"net/http"

	"gitlab.humo.tj/orzu/applications/bridge/pkg/bootstrap/http/router"
	"gitlab.humo.tj/orzu/applications/bridge/pkg/config"
	"go.uber.org/fx"
)

// ServerModule ...
var ServerModule = fx.Provide(NewServer)

// Dependecies ...
type Dependecies struct {
	fx.In

	Config *config.Config
	Router *router.HTTPRouter
}

// NewServer ...
func NewServer(params Dependecies) *http.Server {
	url := net.JoinHostPort(params.Config.Server.Host, fmt.Sprint(params.Config.Server.Port))

	return &http.Server{
		MaxHeaderBytes: 32 << 20, // 32 Mb
		Addr:           url,
		Handler:        params.Router,
	}
}
