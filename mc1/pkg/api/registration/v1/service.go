package v1

import (
	"database/sql"
	"errors"
	"fmt"
	"regexp"
)

// Service holds data to connect registration requests with DB
type Service struct {
	dbConnection *sql.DB
}

// Creates a new Service
func NewService(dbConnection *sql.DB) *Service {
	return &Service{dbConnection: dbConnection}
}

func (s *Service) createRegistration(model RegistrationModel) error {
	isEmailValid, err := validateEmail(model.Email)
	if err != nil {
		return fmt.Errorf("email not valid. %s", err.Error())
	}

	if !isEmailValid {
		return fmt.Errorf("email '%s' not valid. Make sure it has email format.", model.Email)
	}

	err = s.validateIfEmailAlreadyInDB(model)
	if err != nil {
		return fmt.Errorf("email not valid, %s", err.Error())
	}

	isPassValid, err := validatePassword(model.Password)
	if err != nil {
		return fmt.Errorf("password not valid, %s", err.Error())
	}

	if !isPassValid {
		return errors.New("password not valid. Please check if length=6, at least 1 digit, at least 1 alphabet")
	}

	isNumberValid, err := validateMobileNumber(model.MobileNumber)
	if err != nil {
		return fmt.Errorf("mobile number not valid, %s", err.Error())
	}

	if !isNumberValid {
		return fmt.Errorf("mobile number '%s' not valid. Make sure it has 10 digits.", model.MobileNumber)
	}

	model.Create(s.dbConnection)

	return nil
}

func (s *Service) validateIfEmailAlreadyInDB(model RegistrationModel) error {
	var name string
	err := model.FindUserByEmail(s.dbConnection, model.Email).Scan(&name)
	if name == "" && (err == nil || err == sql.ErrNoRows) {
		return nil
	}

	return fmt.Errorf("email %s has alredy been taken by another user", model.Email)
}

func validateEmail(email string) (bool, error) {
	if email == "" {
		return false, errors.New("email can't be blank")
	}

	match, err := regexp.MatchString("^[a-z0-9._%+\\-]+@[a-z0-9.\\-]+\\.[a-z]{2,4}$", email)
	if err != nil {
		return false, fmt.Errorf("something went wrong when validating email: %s", err.Error())
	}

	return match, nil
}

func validateMobileNumber(number string) (bool, error) {
	if number == "" {
		return false, errors.New("mobile number can't be blank")
	}

	match, err := regexp.MatchString("^\\d{10}$", number)
	if err != nil {
		return false, fmt.Errorf("something went wrong when validating mobile number: %s", err.Error())
	}

	return match, nil
}

func validatePassword(pass string) (bool, error) {
	if pass == "" {
		return false, errors.New("password can't be blank")
	}
	if len(pass) < 6 {
		return false, errors.New("password must contain more that 6 characters")
	}
	containsANumber, err := regexp.MatchString("\\d", pass)
	if err != nil {
		return false, fmt.Errorf("something went wrong when validating password: %s", err.Error())
	}
	if !containsANumber {
		return false, errors.New("password must contain at least 1 digit")
	}
	containsALetter, err := regexp.MatchString("[a-zA-Z]", pass)
	if err != nil {
		return false, fmt.Errorf("something went wrong when validating password: %s", err.Error())
	}
	if !containsALetter {
		return false, errors.New("password must contain at least 1 alphabet")
	}

	return true, nil
}
