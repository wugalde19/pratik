package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/wugalde19/pratik/mc1/config"
)

const (
	dbclient = "postgres"
)

// Database holds connection data
type Database struct {
	Host     string
	Name     string
	Password string
	Port     int
	User     string
}

// Creates a new DB
func NewDatabase(cfg *config.DB) (*Database, error) {
	if cfg == nil {
		return nil, errors.New("unable to init Database with no configuration data")
	}

	return &Database{
		Host:     cfg.Host,
		Name:     cfg.DBName,
		Password: cfg.Password,
		Port:     cfg.Port,
		User:     cfg.User,
	}, nil
}

// Returns the connection url string
func (db *Database) ConnectionURL() (string, error) {
	password := os.Getenv(db.Password)
	if password == "" {
		return "", fmt.Errorf("%v env variable could not be found", db.Password)
	}

	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		db.Host, db.Port, db.User, password, db.Name,
	), nil
}

// Connect open the connection with the DB
func (db *Database) Connect() (*sql.DB, error) {
	connectionURL, err := db.ConnectionURL()
	if err != nil {
		return nil, err
	}

	dbConnection, err := sql.Open(dbclient, connectionURL)
	if err != nil {
		return nil, err
	}

	if err = dbConnection.Ping(); err != nil {
		return nil, err
	}

	return dbConnection, nil
}
