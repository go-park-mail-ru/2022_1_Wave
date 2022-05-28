package auth_usecase_test

/*
import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/domain/mocks"
	auth_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/auth"
	auth_mocks "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/auth/mocks"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/auth/proto"
	auth_usecase "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/auth/usecase"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/tools/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUseCaseLogin(t *testing.T) {
	mockSessionRepo := new(auth_mocks.AuthRepo)
	mockUserRepo := new(mocks.UserRepo)
	password := "some_password"
	ph, _ := utils.GetPasswordHash(password)
	passwordHash := string(ph)
	mockUser := &domain.User{
		ID:             1,
		Username:       "aboba",
		Email:          "aboba@aboba.ru",
		Avatar:         "some_url_to_avatar",
		Password:       passwordHash,
		CountFollowing: 0,
	}
	sessionId := "some-session-id"

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("SelectByUsername", mockUser.Username).Return(mockUser, nil)
		mockUserRepo.On("SelectByUsername", mockUser.Email).Return(nil, errors.New("error"))
		mockUserRepo.On("SelectByEmail", mockUser.Email).Return(mockUser, nil)
		mockUserRepo.On("SelectByEmail", mockUser.Username).Return(nil, errors.New("error"))
		mockSessionRepo.On("SetNewAuthorizedSession", mockUser.ID, auth_usecase.SessionExpire).Return(sessionId, nil)

		usecase := auth_usecase.NewAuthService(mockSessionRepo, mockUserRepo)

		var loginData proto.LoginData
		loginData.Login = mockUser.Username
		loginData.Password = password
		sessionResult, err := usecase.Login(context.Background(), &loginData)
		assert.Nil(t, err)
		assert.Equal(t, sessionResult.NewSession.SessionId, sessionId)

		loginData.Login = mockUser.Email
		_, err = usecase.Login(context.Background(), &loginData)
		assert.Nil(t, err)
	})

	mockUser2 := &domain.User{
		ID:             2,
		Username:       "aboba2",
		Email:          "aboba2@aboba.ru",
		Avatar:         "some_url_to_avatar2",
		Password:       passwordHash,
		CountFollowing: 0,
	}
	t.Run("error", func(t *testing.T) {
		mockUserRepo.On("SelectByUsername", mockUser2.Username+"a").Return(nil, errors.New("error"))
		mockUserRepo.On("SelectByEmail", mockUser2.Username+"a").Return(nil, errors.New("error"))
		mockUserRepo.On("SelectByUsername", mockUser2.Username).Return(mockUser2, nil)

		mockSessionRepo.On("SetNewAuthorizedSession", mockUser2.ID, auth_usecase.SessionExpire).Return("", errors.New("error"))

		usecase := auth_usecase.NewAuthService(mockSessionRepo, mockUserRepo)

		var loginData proto.LoginData
		loginData.Login = mockUser2.Username + "a"
		loginData.Password = password

		_, err := usecase.Login(context.Background(), &loginData)
		assert.NotNil(t, err)

		loginData.Login = mockUser2.Username
		loginData.Password = password + "a"

		_, err = usecase.Login(context.Background(), &loginData)
		assert.NotNil(t, err)

		loginData.Login = mockUser2.Username
		loginData.Password = password
		_, err = usecase.Login(context.Background(), &loginData)
		assert.NotNil(t, err)
	})
}

func TestUseCaseLogout(t *testing.T) {
	mockSessionRepo := new(auth_mocks.AuthRepo)
	mockUserRepo := new(mocks.UserRepo)
	sessionId := "some-session-id"

	t.Run("success", func(t *testing.T) {
		mockSessionRepo.On("DeleteSession", sessionId).Return(nil)

		usecase := auth_usecase.NewAuthService(mockSessionRepo, mockUserRepo)

		var session proto.Session
		session.SessionId = sessionId
		_, err := usecase.Logout(context.Background(), &session)
		assert.Nil(t, err)
	})
}

func TestSignUp(t *testing.T) {
	mockSessionRepo := new(mocks.SessionRepo)
	mockUserRepo := new(mocks.UserRepo)
	password := "some_password"
	mockUser := &domain.User{
		ID:             1,
		Username:       "aboba",
		Email:          "aboba@aboba.ru",
		Avatar:         "some_url_to_avatar",
		Password:       password,
		CountFollowing: 0,
	}
	sessionId := "some-session-id"

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("SelectByUsername", mockUser.Username).Return(nil, errors.New("error"))
		mockUserRepo.On("SelectByEmail", mockUser.Email).Return(nil, errors.New("error"))

		mockSessionRepo.On("MakeSessionAuthorized", sessionId, mockUser.ID).Return(nil)

		usecase:= NewAuthUseCase(mockSessionRepo, mockUserRepo)
		result := usecase.Login(mockUser.Username, password, sessionId)
		assert.Nil(t, result)

		result = usecase.Login(mockUser.Email, password, sessionId)
		assert.Nil(t, result)
	})
}


func TestUseCaseGetUnauthorizedSession(t *testing.T) {
	mockSessionRepo := new(auth_mocks.AuthRepo)
	mockUserRepo := new(mocks.UserRepo)
	sessionId := "some-session-id"

	t.Run("success", func(t *testing.T) {
		mockSessionRepo.On("SetNewUnauthorizedSession", auth_usecase.SessionExpire).Return(sessionId, nil)

		usecase := auth_usecase.NewAuthService(mockSessionRepo, mockUserRepo)
		sessionIdRes, result := usecase.GetUnauthorizedSession(context.Background(), &proto.Empty{})
		assert.Nil(t, result)
		assert.Equal(t, sessionIdRes.Session.SessionId, sessionId)
	})
}

func TestUseCaseGetAuthorizedSession(t *testing.T) {
	mockSessionRepo := new(auth_mocks.AuthRepo)
	mockUserRepo := new(mocks.UserRepo)
	sessionId := "some-session-id"

	t.Run("success", func(t *testing.T) {
		mockSessionRepo.On("SetNewAuthorizedSession", uint(1), auth_usecase.SessionExpire).Return(sessionId, nil)

		usecase := auth_usecase.NewAuthService(mockSessionRepo, mockUserRepo)
		sessionIdRes, result := usecase.GetAuthorizedSession(context.Background(), &proto.UserId{UserId: 1})
		assert.Nil(t, result)
		assert.Equal(t, sessionIdRes.Session.SessionId, sessionId)
	})
}

func TestUseCaseIsSession(t *testing.T) {
	mockSessionRepo := new(auth_mocks.AuthRepo)
	mockUserRepo := new(mocks.UserRepo)
	sessionId := "some-session-id"

	t.Run("success", func(t *testing.T) {
		mockSessionRepo.On("GetSession", sessionId).Return(&auth_domain.Session{}, nil)

		usecase := auth_usecase.NewAuthService(mockSessionRepo, mockUserRepo)

		var session proto.Session
		session.SessionId = sessionId

		_, err := usecase.IsSession(context.Background(), &session)
		assert.Nil(t, err)
	})
}

func TestUseCaseIsAuthSession(t *testing.T) {
	mockSessionRepo := new(auth_mocks.AuthRepo)
	mockUserRepo := new(mocks.UserRepo)
	sessionId := "some-session-id"

	t.Run("success", func(t *testing.T) {
		mockSessionRepo.On("GetSession", sessionId).Return(&auth_domain.Session{IsAuthorized: true}, nil)
		mockSessionRepo.On("GetSession", sessionId+"a").Return(nil, errors.New("error"))

		usecase := auth_usecase.NewAuthService(mockSessionRepo, mockUserRepo)

		var session proto.Session
		session.SessionId = sessionId

		_, err := usecase.IsAuthSession(context.Background(), &session)
		assert.Nil(t, err)

		session.SessionId += "a"
		_, err = usecase.IsAuthSession(context.Background(), &session)
		assert.NotNil(t, err)
	})
}
*/
