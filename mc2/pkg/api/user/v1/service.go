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

// Updates user password
func (s *Service) updatePassword(
	userId int, newPass string, oldPass string,
) error {
	sqlStatement := `
		SELECT COUNT(email) count
		FROM users
		WHERE user_id = $1 AND password = crypt($2, password);`
	row := s.dbConnection.QueryRow(sqlStatement, userId, oldPass)

	var count int
	err := row.Scan(&count)
	if err != nil {
		msg := "an error occurred when executing updatePassword query: %s\n"
		return fmt.Errorf(msg, err.Error())
	}

	if count < 1 {
		msg := "password cannot be updated.The id and password provided don't match."
		return errors.New(msg)
	}

	_, err = s.dbConnection.Exec("CALL UpdatePassword($1, $2)", userId, newPass)
	if err != nil {
		msg := "an error occurred when executing updatePassword query: %s\n"
		return fmt.Errorf(msg, err.Error())
	}

	return nil
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
