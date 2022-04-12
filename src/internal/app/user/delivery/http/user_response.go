package http

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	"github.com/microcosm-cc/bluemonday"
)

type UserResponse struct {
	Status string       `json:"status"`
	Error  string       `json:"error,omitempty"`
	Result *domain.User `json:"result,omitempty"`
}

type UserUpdateResponse struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
	Result string `json:"result,omitempty"`
}

type UserAvatarUploadResponse struct {
	Status string `json:"status"`
	Result string `json:"result,omitempty"`
}

const (
	StatusOK   = "OK"
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

func userSanitize(user *domain.User) {
	sanitizer := bluemonday.UGCPolicy()

	user.Username = sanitizer.Sanitize(user.Username)
	user.Email = sanitizer.Sanitize(user.Email)
	user.Avatar = sanitizer.Sanitize(user.Avatar)
}

func getSuccessGetUserResponse(user *domain.User) *UserResponse {
	userSanitize(user)

	return &UserResponse{
		Status: StatusOK,
		Result: user,
	}
}

func getSuccessUserUpdate() *UserUpdateResponse {
	return &UserUpdateResponse{
		Status: StatusOK,
		Result: "successful user update",
	}
}

func getSuccessUploadAvatar() *UserAvatarUploadResponse {
	return &UserAvatarUploadResponse{
		Status: StatusOK,
		Result: "successful avatar upload",
	}
}
