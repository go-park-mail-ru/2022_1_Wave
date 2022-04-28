package auth_usecase_test

import (
	"errors"
	auth_usecase "github.com/go-park-mail-ru/2022_1_Wave/internal/auth/usecase"
	user_microservice_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/user"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/test/mocks"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/tools/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestUseCaseLogin(t *testing.T) {
	mockAuthAgent := new(mocks.AuthAgent)
	mockUserAgent := new(mocks.UserAgent)
	password := "some_password"
	ph, _ := utils.GetPasswordHash(password)
	passwordHash := string(ph)
	mockUser := &user_microservice_domain.User{
		ID:             1,
		Username:       "aboba",
		Email:          "aboba@aboba.ru",
		Avatar:         "some_url_to_avatar",
		Password:       passwordHash,
		CountFollowing: 0,
	}
	sessionId := "some-session-id"

	t.Run("success", func(t *testing.T) {
		mockUserAgent.On("GetByUsername", mockUser.Username).Return(mockUser, nil)
		mockUserAgent.On("GetByUsername", mockUser.Email).Return(nil, errors.New("error"))
		mockUserAgent.On("GetByEmail", mockUser.Email).Return(mockUser, nil)
		mockUserAgent.On("GetByEmail", mockUser.Username).Return(nil, errors.New("error"))
		mockAuthAgent.On("SetNewAuthorizedSession", mockUser.ID, auth_usecase.SessionExpires).Return(sessionId, nil)

		usecase := auth_usecase.NewAuthService(mockAuthAgent, mockUserAgent)

		sessionIdResult, err := usecase.Login(mockUser.Username, password)
		assert.Nil(t, err)
		assert.Equal(t, sessionIdResult, sessionId)

		_, err = usecase.Login(mockUser.Email, password)
		assert.Nil(t, err)
	})

	mockUser2 := &user_microservice_domain.User{
		ID:             2,
		Username:       "aboba2",
		Email:          "aboba2@aboba.ru",
		Avatar:         "some_url_to_avatar2",
		Password:       passwordHash,
		CountFollowing: 0,
	}
	t.Run("error", func(t *testing.T) {
		mockUserAgent.On("GetByUsername", mockUser2.Username+"a").Return(nil, errors.New("error"))
		mockUserAgent.On("GetByEmail", mockUser2.Username+"a").Return(nil, errors.New("error"))
		mockUserAgent.On("GetByUsername", mockUser2.Username).Return(mockUser2, nil)

		mockAuthAgent.On("SetNewAuthorizedSession", mockUser2.ID, auth_usecase.SessionExpires).Return("", errors.New("error"))

		usecase := auth_usecase.NewAuthService(mockAuthAgent, mockUserAgent)

		_, err := usecase.Login(mockUser2.Username+"a", password)
		assert.NotNil(t, err)

		_, err = usecase.Login(mockUser2.Username, password+"a")
		assert.NotNil(t, err)

		_, err = usecase.Login(mockUser2.Username, password)
		assert.NotNil(t, err)
	})
}

func TestUseCaseLogout(t *testing.T) {
	mockAuthAgent := new(mocks.AuthAgent)
	mockUserAgent := new(mocks.UserAgent)
	sessionId := "some-session-id"

	t.Run("success", func(t *testing.T) {
		mockAuthAgent.On("DeleteSession", sessionId).Return(nil)

		usecase := auth_usecase.NewAuthService(mockAuthAgent, mockUserAgent)

		err := usecase.Logout(sessionId)
		assert.Nil(t, err)
	})
	t.Run("error", func(t *testing.T) {
		mockAuthAgent.On("DeleteSession", sessionId+"a").Return(errors.New("error"))

		usecase := auth_usecase.NewAuthService(mockAuthAgent, mockUserAgent)

		err := usecase.Logout(sessionId + "a")
		assert.Error(t, err)
	})
}

func TestSignUp(t *testing.T) {
	mockAuthAgent := new(mocks.AuthAgent)
	mockUserAgent := new(mocks.UserAgent)
	password := "some_password"
	mockUser := &user_microservice_domain.User{
		ID:             1,
		Username:       "aboba",
		Email:          "aboba@aboba.ru",
		Avatar:         "some_url_to_avatar",
		Password:       password,
		CountFollowing: 0,
	}
	mockUser2 := &user_microservice_domain.User{
		ID:             2,
		Username:       "aboba2",
		Email:          "aboba2@aboba.ru",
		Avatar:         "some_url_to_avatar2",
		Password:       password,
		CountFollowing: 0,
	}
	//sessionId := "some-session-id"

	t.Run("error", func(t *testing.T) {
		mockUserAgent.On("GetByUsername", mockUser.Username).Return(nil, errors.New("error"))
		mockUserAgent.On("GetByEmail", mockUser.Email).Return(nil, nil)
		mockUserAgent.On("GetByUsername", mockUser.Username+"a").Return(nil, nil)

		mockUserAgent.On("GetByUsername", mockUser2.Username).Return(nil, errors.New("error"))
		mockUserAgent.On("GetByEmail", mockUser2.Email).Return(nil, errors.New("error"))
		mockUserAgent.On("Create", mock.Anything).Return(errors.New("error"))

		usecase := auth_usecase.NewAuthService(mockAuthAgent, mockUserAgent)
		sessionIdResult, err := usecase.SignUp(mockUser)
		assert.Error(t, err)
		assert.Equal(t, sessionIdResult, "")

		mockUser.Username += "a"
		sessionIdResult, err = usecase.SignUp(mockUser)
		assert.Error(t, err)
		assert.Equal(t, sessionIdResult, "")

		sessionIdResult, err = usecase.SignUp(mockUser2)
		assert.Error(t, err)
		assert.Equal(t, sessionIdResult, "")
	})
}

func TestUseCaseGetUnauthorizedSession(t *testing.T) {
	mockUserAgent := new(mocks.UserAgent)
	sessionId := "some-session-id"

	t.Run("success", func(t *testing.T) {
		mockAuthAgent := new(mocks.AuthAgent)
		mockAuthAgent.On("SetNewUnauthorizedSession", auth_usecase.SessionExpires).Return(sessionId, nil)

		usecase := auth_usecase.NewAuthService(mockAuthAgent, mockUserAgent)
		sessionIdRes, err := usecase.GetUnauthorizedSession()
		assert.NoError(t, err)
		assert.Equal(t, sessionIdRes, sessionId)
	})

	t.Run("error", func(t *testing.T) {
		mockAuthAgent := new(mocks.AuthAgent)
		mockAuthAgent.On("SetNewUnauthorizedSession", auth_usecase.SessionExpires).Return("", errors.New("error"))

		usecase := auth_usecase.NewAuthService(mockAuthAgent, mockUserAgent)
		sessionIdRes, err := usecase.GetUnauthorizedSession()
		assert.Error(t, err)
		assert.Equal(t, sessionIdRes, "")
	})
}

func TestUseCaseIsSession(t *testing.T) {
	mockAuthAgent := new(mocks.AuthAgent)
	mockUserAgent := new(mocks.UserAgent)
	sessionId := "some-session-id"

	t.Run("success", func(t *testing.T) {
		mockAuthAgent.On("IsSession", sessionId).Return(true)

		usecase := auth_usecase.NewAuthService(mockAuthAgent, mockUserAgent)

		res := usecase.IsSession(sessionId)
		assert.True(t, res)
	})
}

func TestUseCaseIsAuthSession(t *testing.T) {
	mockAuthAgent := new(mocks.AuthAgent)
	mockUserAgent := new(mocks.UserAgent)
	sessionId := "some-session-id"

	t.Run("success", func(t *testing.T) {
		mockAuthAgent.On("IsAuthSession", sessionId).Return(true)

		usecase := auth_usecase.NewAuthService(mockAuthAgent, mockUserAgent)

		res := usecase.IsAuthSession(sessionId)
		assert.True(t, res)
	})
}
