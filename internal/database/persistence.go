package database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/z3orc/compass/internal/env"
)

func GetPostgressClient() *sqlx.DB {

    src := DatabaseConfig {
        driver: PostgresDriver,
        user: env.PGUser(),
        dbname: env.PGDatabase(),
        sslmode: SSLModeDisable,
        password: env.PGPassword(),
        host: env.PGDatabase(),
    }
    
    db, err := sqlx.Connect(string(src.driver), src.asDataSource())
    if(err != nil){
        log.Fatalln(fmt.Sprint("error while connecting to database: ", err.Error()))
    }

    defer db.Close()

    if err := db.Ping(); err != nil {
        log.Fatalln(fmt.Sprint("error while connecting to database: ", err.Error()))
        return nil
    } else {
        log.Println("Database connected")
        return db
    }
}