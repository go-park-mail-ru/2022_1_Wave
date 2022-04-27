package test

import (
	"errors"
	usecase2 "github.com/go-park-mail-ru/2022_1_Wave/internal/app/auth/usecase"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain/mocks"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/tools/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUseCaseLogin(t *testing.T) {
	mockSessionRepo := new(mocks.SessionRepo)
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
		mockSessionRepo.On("MakeSessionAuthorized", sessionId, mockUser.ID).Return(nil)

		usecase := usecase2.NewAuthUseCase(mockSessionRepo, mockUserRepo)
		result := usecase.Login(mockUser.Username, password, sessionId)
		assert.Nil(t, result)

		result = usecase.Login(mockUser.Email, password, sessionId)
		assert.Nil(t, result)
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

		mockSessionRepo.On("MakeSessionAuthorized", sessionId+"a", mockUser2.ID).Return(errors.New("error"))

		usecase := usecase2.NewAuthUseCase(mockSessionRepo, mockUserRepo)
		result := usecase.Login(mockUser2.Username+"a", password, sessionId)
		assert.NotNil(t, result)

		result = usecase.Login(mockUser2.Username, password+"a", sessionId)
		assert.NotNil(t, result)

		result = usecase.Login(mockUser2.Username, password, sessionId+"a")
		assert.NotNil(t, result)
	})
}

func TestUseCaseLogout(t *testing.T) {
	mockSessionRepo := new(mocks.SessionRepo)
	mockUserRepo := new(mocks.UserRepo)
	sessionId := "some-session-id"

	t.Run("success", func(t *testing.T) {
		mockSessionRepo.On("MakeSessionUnauthorized", sessionId).Return(nil)

		usecase := usecase2.NewAuthUseCase(mockSessionRepo, mockUserRepo)
		result := usecase.Logout(sessionId)
		assert.Nil(t, result)
	})
}

/*func TestSignUp(t *testing.T) {
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
*/

func TestUseCaseGetUnauthorizedSession(t *testing.T) {
	mockSessionRepo := new(mocks.SessionRepo)
	mockUserRepo := new(mocks.UserRepo)
	sessionId := "some-session-id"

	t.Run("success", func(t *testing.T) {
		mockSessionRepo.On("SetNewUnauthorizedSession", usecase2.SessionExpire).Return(sessionId, nil)

		usecase := usecase2.NewAuthUseCase(mockSessionRepo, mockUserRepo)
		sessionIdRes, result := usecase.GetUnauthorizedSession()
		assert.Nil(t, result)
		assert.Equal(t, sessionIdRes, sessionId)
	})
}

func TestUseCaseIsSession(t *testing.T) {
	mockSessionRepo := new(mocks.SessionRepo)
	mockUserRepo := new(mocks.UserRepo)
	sessionId := "some-session-id"

	t.Run("success", func(t *testing.T) {
		mockSessionRepo.On("GetSession", sessionId).Return(&domain.Session{}, nil)

		usecase := usecase2.NewAuthUseCase(mockSessionRepo, mockUserRepo)
		result := usecase.IsSession(sessionId)
		assert.True(t, result)
	})
}

func TestUseCaseIsAuthSession(t *testing.T) {
	mockSessionRepo := new(mocks.SessionRepo)
	mockUserRepo := new(mocks.UserRepo)
	sessionId := "some-session-id"

	t.Run("success", func(t *testing.T) {
		mockSessionRepo.On("GetSession", sessionId).Return(&domain.Session{IsAuthorized: true}, nil)
		mockSessionRepo.On("GetSession", sessionId+"a").Return(nil, errors.New("error"))

		usecase := usecase2.NewAuthUseCase(mockSessionRepo, mockUserRepo)
		result := usecase.IsAuthSession(sessionId)
		assert.True(t, result)

		result = usecase.IsAuthSession(sessionId + "a")
		assert.False(t, result)
	})
}
