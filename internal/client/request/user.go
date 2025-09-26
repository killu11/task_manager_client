package request

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type RegisterRequest struct {
	UserRequest
}

type LoginRequest struct {
	UserRequest
}

func NewRegisterRequest(username, password string) *RegisterRequest {
	return &RegisterRequest{
		UserRequest: UserRequest{
			Username: username,
			Password: password,
		},
	}
}

func NewLoginRequest(username, password string) *LoginRequest {
	return &LoginRequest{
		UserRequest: UserRequest{
			Username: username,
			Password: password,
		},
	}
}
