package databases

import (
	"database/sql"

	"github.com/Jacobamv/golang_arch_template/pkg/config"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

// OracleModule ...
var OracleModule = fx.Provide(NewOracleConn)

// PostgresModule ...
var PostgresModule = fx.Provide(NewPostgresConn)

// Dependencies ...
type Dependencies struct {
	fx.In

	Logger zerolog.Logger
	Config *config.Config
}

// NewOracleConn ...
func NewOracleConn(params Dependencies) *sql.DB {
	return Oracle(params)
}

// NewPostgresConn ...
func NewPostgresConn(params Dependencies) *pgxpool.Pool {
	return Postgres(params)
}
