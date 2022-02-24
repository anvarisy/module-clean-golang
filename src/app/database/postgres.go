package database

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func InitPostgres() *sql.DB {
	dbDriver := viper.GetString("DB_DRIVER")
	dbSource := viper.GetString("DB_SOURCE")

	db, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		panic(err)
	}

	return db
}
