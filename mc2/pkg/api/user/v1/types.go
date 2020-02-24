package v1

// UsersCountResponse is used to return request response in /v1/user/count
type UsersCountResponse struct {
	Count int    `json:"count"`
	Error string `json:"error"`
}
