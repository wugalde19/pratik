package v1

import (
	"database/sql"
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

// Returns the number of users registered in the DB
func (s *Service) usersCount() (int, error) {
	row := s.dbConnection.QueryRow("SELECT COUNT(email) count FROM users")

	var count int
	err := row.Scan(&count)
	if err != nil {
		msg := "an error occurred when executing usersCount query: %s\n"
		return count, fmt.Errorf(msg, err.Error())
	}

	return count, nil
}
