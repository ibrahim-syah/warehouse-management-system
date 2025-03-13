package databaseutils

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/spf13/viper"
)

func ConnectDB() (*sql.DB, error) {
	host := viper.GetString("DB_HOST")
	port := viper.GetString("DB_PORT")
	username := viper.GetString("DB_USER")
	password := viper.GetString("DB_PASS")
	databaseName := viper.GetString("DB_NAME")
	db, err := sql.Open("pgx", fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, databaseName))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
