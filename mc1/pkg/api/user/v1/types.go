package v1

// UserModel is used store user data coming from the DB
type UserModel struct {
	MobileNumber string	`sql:"mobile_number"`
	Name         string	`sql:"name"`
}

// UserData is used store user data
type UserData struct {
	MobileNumber	string	`json:"mobile_number"`
	Name					string	`json:"name"`
}

// UserDetailsResponse is used to return request response in /v1/user/
type UserDetailsResponse struct {
	Error			string 		`json:"error"`
	UserData	UserData 	`json:"user_data"`
}
