package v1

// LoginData is used to decode login request data
type LoginData struct {
	Password     string `json:"password"`
	MobileNumber string `json:"mobile_number"`
}

// LoginResponse is used to return request response
type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}
