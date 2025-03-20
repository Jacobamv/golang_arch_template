package databases

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Jacobamv/golang_arch_template/pkg/config"
	_ "github.com/godror/godror"
	"github.com/rs/zerolog"
)

// Oracle ...
func Oracle(params Dependencies) (odb *sql.DB) {
	connStr := fmt.Sprintf(
		`user=%s password=%s connectString=%s poolMaxSessions=%v poolMinSessions=%v poolSessionMaxLifetime=%v timezone="Asia/Dushanbe"`,

		params.Config.Oracle.Username,
		params.Config.Oracle.Password,
		params.Config.Oracle.ConnectString,
		params.Config.Oracle.PoolMaxSessions,
		params.Config.Oracle.PoolMinSessions,
		params.Config.Oracle.PoolSessionMaxLifetime,
	)

	odb, err := sql.Open("godror", connStr)
	if err != nil {
		log.Fatal(err)
	}

	odb.SetMaxOpenConns(0)
	odb.SetMaxIdleConns(0)
	odb.SetConnMaxLifetime(time.Minute)

	if err = odb.Ping(); err != nil {
		log.Fatal(err)

		return
	}

	go runScheduler(odb, params.Logger, params.Config)

	return
}

func runScheduler(oracle *sql.DB, log zerolog.Logger, config *config.Config) {
	timer := time.NewTicker(time.Duration(config.Oracle.PingEvery) * time.Minute)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			ping(oracle, log)
		}
	}
}

func ping(conn *sql.DB, log zerolog.Logger) {
	log.Info().Msg("PING")

	err := conn.Ping()
	if err != nil {

		return
	}

}
