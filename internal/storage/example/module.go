package example

import (
	"context"
	"database/sql"

	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewSExample)

type SExample interface {
	Example(ctx context.Context) (err error)
}

type provider struct {
	// Logger

	logger zerolog.Logger
	// Database
	oracleDB *sql.DB
}

type Params struct {
	fx.In

	// Logger
	Logger zerolog.Logger
	// Database
}

// NewSExample ...
func NewSExample(params Params) SExample {
	return &provider{
		// Logger
		logger: params.Logger,
		// Database
	}
}
