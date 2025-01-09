package mysql

import (
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

var DB *sqlx.DB

func Open() {
	var (
		host = os.Getenv("DB_HOST")
		port = os.Getenv("DB_PORT")
		user = os.Getenv("DB_USER")
		pass = os.Getenv("DB_PASS")
		name = os.Getenv("DB_NAME")
	)

	urlConn := fmt.Sprintf("%s:%s@(%s:%s)/%s", user, pass, host, port, name)

	// Initialize a mysql database connection
	db, err := sqlx.Connect("mysql", urlConn)
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	// Verify the connection to the database is still alive
	err = db.Ping()
	if err != nil {
		panic("Failed to ping the database: " + err.Error())
	}

	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	DB = db

	log.Info().Msgf("Connected to database %v", name)
}
