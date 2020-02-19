package v1

import (
	"database/sql"
	"fmt"
)

// RegistrationModel is used to decode registration request data
type RegistrationModel struct {
	Email        string `json:"email"`
	MobileNumber string `json:"mobile_number"`
	Name         string `json:"name"`
	Password     string `json:"password"`
}

// Creates a new registration in the DB
func (rm *RegistrationModel) Create(dbConnection *sql.DB) {
	_, err := dbConnection.Exec(
		"CALL registerUser($1, $2, $3, $4);",
		rm.Name, rm.MobileNumber, rm.Password, rm.Email,
	)
	if err != nil {
		fmt.Printf("an error occurred when executing procedure: %s\n", err.Error())
		return
	}

	fmt.Printf("%s has been successfully registered!\n", rm.Name)

}

// Checks if user alredy registered with
func (rm *RegistrationModel) FindUserByEmail(dbConnection *sql.DB, value string) *sql.Row {
	return dbConnection.QueryRow("SELECT name FROM users WHERE email = $1", value)
}
