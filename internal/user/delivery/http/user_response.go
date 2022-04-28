package userHttp

import (
	user_microservice_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/user"
	"github.com/microcosm-cc/bluemonday"
)

type UserResponse struct {
	Status string                         `json:"status"`
	Error  string                         `json:"error,omitempty"`
	Result *user_microservice_domain.User `json:"result,omitempty"`
}

type UserResponseProto struct {
	Status string                         `json:"status"`
	Error  string                         `json:"error,omitempty"`
	Result *user_microservice_domain.User `json:"result,omitempty"`
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

func userSanitize(user *user_microservice_domain.User) {
	sanitizer := bluemonday.UGCPolicy()

	user.Username = sanitizer.Sanitize(user.Username)
	user.Email = sanitizer.Sanitize(user.Email)
	user.Avatar = sanitizer.Sanitize(user.Avatar)
}

func userSanitizeProto(user *user_microservice_domain.User) {
	sanitizer := bluemonday.UGCPolicy()

	user.Username = sanitizer.Sanitize(user.Username)
	user.Email = sanitizer.Sanitize(user.Email)
	user.Avatar = sanitizer.Sanitize(user.Avatar)
}

func getSuccessGetUserResponse(user *user_microservice_domain.User) *UserResponse {
	userSanitize(user)

	return &UserResponse{
		Status: StatusOK,
		Result: user,
	}
}

func getSuccessGetUserResponseProto(user *user_microservice_domain.User) *UserResponseProto {
	userSanitizeProto(user)

	return &UserResponseProto{
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
