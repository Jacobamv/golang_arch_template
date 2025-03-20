package example

import (
	"context"

	"github.com/rs/zerolog"
	"gitlab.humo.tj/orzu/applications/bridge/pkg/config"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewGExample)

type GExample interface {
	Ping(ctx context.Context) (err error)
}

type provider struct {
	// Logger
	logger zerolog.Logger

	// Config
	config *config.Config

	TokenType    string `json:"token_type"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Params struct {
	fx.In

	// Logger
	Logger zerolog.Logger

	// Config
	Config *config.Config
}

// NewGExample ...
func NewGExample(params Params) GExample {

	GExample := &provider{
		// Logger
		logger: params.Logger,

		// Database
		config: params.Config,
	}

	err := GExample.Ping(context.Background())

	if err != nil {
		GExample.logger.Error().Err(err).Msg("An error occurred in Score ping")
	}

	GExample.logger.Info().Err(err).Interface("Example url", params.Config.Example.Url).Send()
	return GExample
}
