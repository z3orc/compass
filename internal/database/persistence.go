package database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
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
    
    log.Printf("connect to postgress, with host: %s, user: %s, dbname: %s ", config.host, config.user, config.dbname)
    fmt.Println(config.asDataSource())
    db, err := sqlx.Connect(string(config.driver), config.asDataSource())
    if(err != nil){
        log.Fatalln(fmt.Sprint("error while connecting to database: ", err.Error()))
    }

    defer db.Close()

    if err := db.Ping(); err != nil {
        log.Fatalln(fmt.Sprint("error while pinging database: ", err.Error()))
        return nil
    } else {
        log.Println("successfully connected to database")
        return db
    }
}
