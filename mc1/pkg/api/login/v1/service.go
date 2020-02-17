package v1

import (
	"database/sql"
	"errors"
	"fmt"
)

// Service holds data to connect registration requests with DB
type Service struct {
	dbConnection *sql.DB
}

// Creates a new Service
func NewService(dbConnection *sql.DB) *Service {
	return &Service{dbConnection: dbConnection}
}

// Checks if user exists in DB
func (s *Service) executeLogin(data LoginData) (string, error) {
	rows, err := s.dbConnection.Query("SELECT loginUser($1, $2);", data.MobileNumber, data.Password)
	if err != nil {
		msg := "an error occurred when executing login function: %s\n"
		return "", fmt.Errorf(msg, err.Error())
	} else if err == sql.ErrNoRows {
		msg := "Invalid credentials. Not user found with those credentials"
		return "", fmt.Errorf(msg, err.Error())
	}

	var name *string

	for rows.Next() {
		err = rows.Scan(&name)
		if err != nil {
			msg := "an error occurred when scanning rows in executeLogin: %s %#v\n"
			return "", fmt.Errorf(msg, err.Error(), rows)
		}
	}

	err = rows.Err()
	if err != nil {
		msg := "an error occurred during iteration of rows in executeLogin: %s\n"
		return "", fmt.Errorf(msg, err.Error())
	}

	if name == nil {
		return "", errors.New("Invalid credentials. Not user found with those credentials")
	}

	return fmt.Sprintf("valid credentials for %s", *name), nil
}
