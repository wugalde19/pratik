package v1

// UpdatePasswordRequestParams is used to marshal request params in /v1/user/update-password
type UpdatePasswordRequestParams struct {
	NewPassword string `json:"new_password"`
	OldPassword string `json:"old_password"`
	UserId      int    `json:"user_id"`
}

// UpdatePasswordResponse is used to return request response in /v1/user/update-password
type UpdatePasswordResponse struct {
	Message string `json:"message"`
}

// UsersCountResponse is used to return request response in /v1/user/count
type UsersCountResponse struct {
	Count int    `json:"count"`
	Error string `json:"error"`
}
