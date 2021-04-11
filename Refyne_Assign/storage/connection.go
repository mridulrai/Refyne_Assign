package storage

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	// mysql driver
	"../logger"
	_ "github.com/go-sql-driver/mysql"
)

const connectionMaxLifetime = time.Minute * 3
const maxOpenConnections = 10
const maxIdleConnections = 10

// Creates a new client for database
func createClient(dbName string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn(dbName))
	if err == nil {
		db.SetConnMaxLifetime(connectionMaxLifetime)
		db.SetMaxOpenConns(maxOpenConnections)
		db.SetMaxIdleConns(maxIdleConnections)
	}
	return db, err
}

// Generates data source name
func dsn(dbName string) string {
	if len(os.Getenv("DB_USERNAME")) == 0 {
		panic("Environment variable for database username is not set.")
	}

	if len(os.Getenv("DB_PASSWORD")) == 0 {
		panic("Environment variable for database password is not set.")
	}

	if len(os.Getenv("DB_HOSTNAME")) == 0 {
		panic("Environment variable for database hostname is not set.")
	}

	return fmt.Sprintf("%s:%s@tcp(%s)/%s", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOSTNAME"), dbName)
}

// Health checks health of database
func Health() error {
	slog := logger.InitSugarLogger()
	defer slog.Sync() // Flushes buffer, if any

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	db, err := createClient(os.Getenv("DB_NAME"))
	if err != nil {
		slog.Errorw("Unable to create client for database "+os.Getenv("DB_NAME"),
			"error", err)
		return err
	}
	defer db.Close()
	err = db.PingContext(ctx)
	if err != nil {
		slog.Errorw("Unable to ping database "+os.Getenv("DB_NAME"),
			"error", err)
	}
	return err
}
