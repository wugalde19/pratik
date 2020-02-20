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

// Search user in DB using the email
func (s *Service) findUserByEmail(email string) (*UserModel, error) {
	row := s.dbConnection.QueryRow(
		"SELECT name, mobile_number FROM users WHERE email = $1",
		email,
	)

	model := &UserModel{}
	err := row.Scan(&model.Name, &model.MobileNumber)
	if err != nil {
		msg := "an error occurred when executing findUserByEmail query: %s\n"
		return model, fmt.Errorf(msg, err.Error())
	} else if err == sql.ErrNoRows {
		msg := "Invalid credentials. Not user found with those credentials"
		return model, fmt.Errorf(msg, err.Error())
	}

	return model, nil
}