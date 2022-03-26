package http

type AuthResponse struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
	Result string `json:"result,omitempty"`
}

const (
	statusOK   = "OK"
	statusFAIL = "FAIL"
)

const (
	successLogin  = "you are login"
	successSignUp = "you are sign up"
)

func getErrorAuthResponse(err error) *AuthResponse {
	return &AuthResponse{
		Status: statusFAIL,
		Error:  err.Error(),
	}
}

func getSuccessLoginResponse() *AuthResponse {
	return &AuthResponse{
		Status: statusOK,
		Result: successLogin,
	}
}

func getSuccessSignUpResponse() *AuthResponse {
	return &AuthResponse{
		Status: statusOK,
		Result: successSignUp,
	}
}
