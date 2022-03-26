package http

import "github.com/go-park-mail-ru/2022_1_Wave/internal/domain"

type UserResponse struct {
	Status string       `json:"status"`
	Error  string       `json:"error,omitempty"`
	Result *domain.User `json:"result,omitempty"`
}

const (
	statusOK   = "OK"
	statusFAIL = "FAIL"
)

const (
	successLogin  = "you are login"
	successSignUp = "you are sign up"
)

func getErrorUserResponse(err error) *UserResponse {
	return &UserResponse{
		Status: statusFAIL,
		Error:  err.Error(),
	}
}

func getSuccessGetUserResponse(user *domain.User) *UserResponse {
	return &UserResponse{
		Status: statusOK,
		Result: user,
	}
}

func getSuccessUserUpdate(user *domain.User) *UserResponse {
	return &UserResponse{
		Status: statusOK,
		Result: user,
	}
}
