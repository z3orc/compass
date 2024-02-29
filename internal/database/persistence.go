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
        sslmode: SSLModeEnable,
        password: env.PGPassword(),
        host: env.PGDatabase(),
    }
    
    fmt.Println(config.asDataSource())
    db, err := sqlx.Connect(string(config.driver), config.asDataSource())
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
