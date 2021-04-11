package storage

import (
	"context"
	"fmt"
	"os"
	"time"

	"../logger"
	"github.com/google/uuid"
)

// CreateDatabase creates database for application on server
func CreateDatabase(dbname string) error {
	slog := logger.InitSugarLogger()
	defer slog.Sync() // Flushes buffer, if any

	db, err := createClient("")
	if err != nil {
		slog.Errorw("Unable to create database client",
			"error", err)
		return err
	}
	defer db.Close()

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	query := "CREATE DATABASE IF NOT EXISTS " + os.Getenv("DB_NAME")
	res, err := db.ExecContext(ctx, query)
	if err != nil {
		slog.Errorw("Unable to execute query",
			"query", query,
			"error", err)
		return err
	}

	no, err := res.RowsAffected()
	if err != nil {
		slog.Errorw("Unable to fetch rows affected",
			"query", query,
			"error", err)
		return err
	}

	msg := fmt.Sprintf("%d rows affected on running query", no)
	slog.Infow(msg,
		"query", query)

	return err
}

// CreateTables creates table(s) required in application database
func CreateTables() error {
	slog := logger.InitSugarLogger()
	defer slog.Sync() // Flushes buffer, if any

	db, err := createClient(os.Getenv("DB_NAME"))
	defer db.Close()
	if err != nil {
		slog.Errorw("Unable to create client for database "+os.Getenv("DB_NAME"),
			"error", err)
		return err
	}

	query := "CREATE TABLE IF NOT EXISTS account(id VARCHAR(36) PRIMARY KEY, created DATETIME DEFAULT CURRENT_TIMESTAMP, modified DATETIME DEFAULT CURRENT_TIMESTAMP, payment_processor_id VARCHAR(50), payment_processor VARCHAR(20), wallet_id VARCHAR(36) NOT NULL UNIQUE, user_id VARCHAR(50) NOT NULL UNIQUE, status ENUM('active','blocked'), active BOOLEAN DEFAULT true)"

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	res, err := db.ExecContext(ctx, query)
	if err != nil {
		slog.Errorw("Unable to create account table in database "+os.Getenv("DB_NAME"),
			"query", query,
			"error", err)
		return err
	}

	no, err := res.RowsAffected()
	if err != nil {
		slog.Errorw("Unable to fetch rows affected",
			"query", query,
			"error", err)
		return err
	}

	msg := fmt.Sprintf("%d rows affected on running query", no)
	slog.Infow(msg,
		"query", query)

	return err
}

// CreateAccount ne user
func CreateUser(user *User) error {
	slog := logger.InitSugarLogger()
	defer slog.Sync() // Flushes buffer, if any

	db, err := createClient(os.Getenv("DB_NAME"))
	defer db.Close()
	if err != nil {
		slog.Errorw("Unable to create client for database "+os.Getenv("DB_NAME"),
			"error", err)
		return err
	}

	user.ID = uuid.New().String()

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	// Begin database transaction
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		slog.Errorw("Unable to begin database transaction for creating new account",
			"error", err)
		return err
	}

	query := "INSERT INTO User (id,mobile, active) VALUES (?, ?, true)"
	_, err = tx.ExecContext(ctx, query, user.ID, user.Mobile, user.Active)
	if err != nil {
		slog.Errorw("Unable to execute query in database transaction",
			"query", query,
			"error", err)
		slog.Infow("Rolling back transaction to create account as the database query could not be executed")
		tx.Rollback()
		return err
	}

	// Commit the change if all queries ran successfully
	err = tx.Commit()
	if err != nil {
		slog.Errorw("Unable to commit transaction to database",
			"error", err)
	}

	return err
}

// Creat Car
func CreateCar(car *Car) error {
	slog := logger.InitSugarLogger()
	defer slog.Sync() // Flushes buffer, if any

	db, err := createClient(os.Getenv("DB_NAME"))
	defer db.Close()
	if err != nil {
		slog.Errorw("Unable to create client for database "+os.Getenv("DB_NAME"),
			"error", err)
		return err
	}

	car.ID = uuid.New().String()

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	// Begin database transaction
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		slog.Errorw("Unable to begin database transaction for creating new account",
			"error", err)
		return err
	}

	query := "INSERT INTO Car (id,model,manufacturer,carLicenseNumber,basePrice,securitydeposit,PPH,available) VALUES (?,?,?,?,?,?,?, true)"
	_, err = tx.ExecContext(ctx, query, car.ID, car.Model, car.Manufacturer, car.CarLicenseNumber, car.BasePrice, car.Securitydeposit, car.PPH, car.Available)
	if err != nil {
		slog.Errorw("Unable to execute query in database transaction",
			"query", query,
			"error", err)
		slog.Infow("Rolling back transaction to create account as the database query could not be executed")
		tx.Rollback()
		return err
	}

	// Commit the change if all queries ran successfully
	err = tx.Commit()
	if err != nil {
		slog.Errorw("Unable to commit transaction to database",
			"error", err)
	}

	return err
}

func CreateBooking(carbooking *carBooking) error {
	slog := logger.InitSugarLogger()
	defer slog.Sync() // Flushes buffer, if any

	db, err := createClient(os.Getenv("DB_NAME"))
	defer db.Close()
	if err != nil {
		slog.Errorw("Unable to create client for database "+os.Getenv("DB_NAME"),
			"error", err)
		return err
	}

	carbooking.BookingId = uuid.New().String()

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	// Begin database transaction
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		slog.Errorw("Unable to begin database transaction for creating new account",
			"error", err)
		return err
	}

	// check car avaibality

	isAvaliable := true
	if isAvaliable == true {
		query := "INSERT INTO carBooking (BookingId,CarID,UserID,StartDateTime,EndDateTime) VALUES (?,?,?,?,?)"
		_, err = tx.ExecContext(ctx, query, carbooking.BookingId, carbooking.CarID, carbooking.UserID, carbooking.StartDateTime, carbooking.EndDateTime)
		if err != nil {
			slog.Errorw("Unable to execute query in database transaction",
				"query", query,
				"error", err)
			slog.Infow("Rolling back transaction to create account as the database query could not be executed")
			tx.Rollback()
			return err
		}

		// Commit the change if all queries ran successfully
		err = tx.Commit()
		if err != nil {
			slog.Errorw("Unable to commit transaction to database",
				"error", err)
		}
	}
	return err
}

// GetAccount fetches account details from database
func GetAccount(id string) (*User, error) {
	slog := logger.InitSugarLogger()
	var account User
	defer slog.Sync() // Flushes buffer, if any

	db, err := createClient(os.Getenv("DB_NAME"))
	defer db.Close()
	if err != nil {
		slog.Errorw("Unable to create client for database "+os.Getenv("DB_NAME"),
			"error", err)
		return nil, err
	}

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	query := "SELECT id, created, modified, payment_processor_id, payment_processor, wallet_id, status, user_id FROM account WHERE id = ? AND active = true"
	row := db.QueryRowContext(ctx, query, id)

	if row == nil {
		slog.Errorw("Unable to fetch account with id "+id,
			"query", query)
		return nil, nil
	}

	if err != nil {
		slog.Errorw("Unable to map data fetched from database "+os.Getenv("DB_NAME")+" for account with id "+id,
			"error", err)
		return nil, err
	}

	return &account, nil
}

// DeleteAccount account deletes account from database
func DeleteAccount(id string) (int64, error) {
	slog := logger.InitSugarLogger()
	defer slog.Sync() // Flushes buffer, if any

	db, err := createClient(os.Getenv("DB_NAME"))
	defer db.Close()
	if err != nil {
		slog.Errorw("Unable to create client for database "+os.Getenv("DB_NAME"),
			"error", err)
		return 0, err
	}

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	query := "UPDATE account SET active = false WHERE id = '" + id + "'"
	result, err := db.ExecContext(ctx, query)

	if err != nil {
		slog.Errorw("Unable to delete account with id "+id,
			"error", err)
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	return rowsAffected, err
}

// GetAccountList fetches list of accounts from database
func GetAccountList(pageNumber int, pageSize int) (int, []User, error) {
	slog := logger.InitSugarLogger()
	var accounts []User
	defer slog.Sync() // Flushes buffer, if any

	db, err := createClient(os.Getenv("DB_NAME"))
	defer db.Close()
	if err != nil {
		slog.Errorw("Unable to create client for database "+os.Getenv("DB_NAME"),
			"error", err)
		return 0, nil, err
	}

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	query := "SELECT id, created, modified, payment_processor_id, payment_processor, wallet_id, status, user_id FROM account WHERE active = true"

	results, err := db.QueryContext(ctx, query)

	if err != nil {
		slog.Errorw("Unable to create client for database "+os.Getenv("DB_NAME"),
			"error", err)
		return 0, nil, err
	}

	itemCount := 0
	for results.Next() {
		var account User

		if err != nil {
			slog.Errorw("Unable to map fields to object",
				"error", err)
			return 0, nil, err
		}
		accounts = append(accounts, account)
		itemCount++
	}

	return itemCount, accounts, nil
}
