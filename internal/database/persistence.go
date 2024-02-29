package database

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/z3orc/compass/internal/env"
)

func GetPostgressClient() *sqlx.DB {

    config := DatabaseConfig {
        driver: PostgresDriver,
        user: env.PGUser(),
        dbname: env.PGDatabase(),
        sslmode: SSLModeDisable,
        password: env.PGPassword(),
        host: env.PGHost(),
    }

    log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

    log.Info().Msgf("connect to postgress, with host: %s, user: %s, dbname: %s ", config.host, config.user, config.dbname)
    db, err := sqlx.Connect(string(config.driver), config.asDataSource())
    if(err != nil){
        log.Error().AnErr("failed to connect to database", err)
    }

    defer db.Close()

    if err := db.Ping(); err != nil {
        log.Error().AnErr("failed to ping database", err)
        return nil
    } else {
        log.Info().Msg("successfully connected to database")
        return db
    }
}
