package authHttp

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
	successLogin   = "you are login"
	successSignUp  = "you are sign up"
	successGetCSRF = "success get csrf"
	successLogout  = "you are logout"
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

func getSuccessGetCSRFResponse() *AuthResponse {
	return &AuthResponse{
		Status: statusOK,
		Result: successGetCSRF,
	}
}

func getSuccessLogoutResponse() *AuthResponse {
	return &AuthResponse{
		Status: statusOK,
		Result: successLogout,
	}
}
