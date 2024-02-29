package database

import "fmt"

type DatabaseConfig struct {
    driver Driver
    user string
    dbname string
    sslmode SSLMode
    password string
    host string
}

type SSLMode string

const (
    SSLModeEnable SSLMode = "enable"
    SSLModeDisable SSLMode = "disable"
)

type Driver string

const (
    PostgresDriver Driver = "postgres"
)

func (src *DatabaseConfig) asDataSource() string {
    return fmt.Sprintf("user=%s dbname=%s sslmode=%s password=%s host=%s", src.user, src.dbname, src.sslmode, src.password, src.host)
}
