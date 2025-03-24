package repository

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConnection struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

type Repository struct {
	Orm    *gorm.DB
	RawSql *sql.DB
}

func NewRepo() (*Repository, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	config := DBConnection{
		Host:     host,
		Port:     5432,
		User:     user,
		Password: password,
		DBName:   dbName,
	}
	// Recieve a structure with the connection parameters
	// Return a pointer and an error

	// Create a Data Source Name (dsn) wich is the connection string for the db
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBName)

	// Configure GORM logs for show INFO level logs.
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	// Opening the Database connection
	dbOrm, err := gorm.Open(postgres.Open(dsn), gormConfig)

	// If the connection fails return the error.
	if err != nil {
		return nil, fmt.Errorf("Error to connect database with GORM: %w", err)
	}

	// Extracts the sql connection from GORM
	// For Raw queries
	dbRaw, err := dbOrm.DB()

	if err != nil {
		return nil, fmt.Errorf("Error getting sql for raw queries: %w", err)
	}

	// Allows 25 simultaneous connections
	dbRaw.SetMaxOpenConns(25)
	// Keep 5 inactive connections ready for use
	dbRaw.SetMaxIdleConns(5)
	// Close the connections after of 5 minutes
	dbRaw.SetConnMaxLifetime(5 * time.Minute)

	// Verify the accessibility of the connection
	// If fails the returns an error
	if err := dbRaw.Ping(); err != nil {
		return nil, fmt.Errorf("Error to ping sql database: %w", err)
	}

	// Return a repository instance with both connections (GORM and raw sql)
	// nil There are no errors.
	return &Repository{
		Orm:    dbOrm,
		RawSql: dbRaw,
	}, nil
}
