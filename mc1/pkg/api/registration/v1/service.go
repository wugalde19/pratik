package v1

import (
	"database/sql"
)

// Service holds data to connect registration requests with DB
type Service struct {
	dbConnection *sql.DB
}

// Creates a new Service
func NewService(dbConnection *sql.DB) *Service {
	return &Service{dbConnection: dbConnection}
}

func (s *Service) createRegistration(model RegistrationModel) {
	model.Create(s.dbConnection)
}
