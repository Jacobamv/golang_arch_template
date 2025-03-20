package databases

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Postgres ...
func Postgres(params Dependencies) (con *pgxpool.Pool) {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		
		params.Config.Postgres.Username,
		params.Config.Postgres.Password,
		params.Config.Postgres.Host,
		fmt.Sprint(params.Config.Postgres.Port),
		params.Config.Postgres.DatabaseName,
	)

	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		log.Println("Unable to parse DATABASE_URL:", err)
		// Надо?
		return
	}

	con, err = pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		log.Println(err)
		// Надо?
		return
	}

	err = con.Ping(context.Background())

	if err != nil {
		log.Println(err)
		// Надо?
		return
	}
	return
}
