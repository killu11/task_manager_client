package response

type UserResponse struct {
	Token    string `json:"token,omitempty"`
	Username string `json:"username,omitempty"`
}
type RegisterResponse struct {
	UserResponse
}

type LoginResponse struct {
	UserResponse
}
