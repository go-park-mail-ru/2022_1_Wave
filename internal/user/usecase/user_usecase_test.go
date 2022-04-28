package user_usecase_test

import (
	"errors"
	auth_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/auth"
	user_microservice_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/user"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/test/mocks"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/tools/utils"
	user_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/user"
	user_usecase "github.com/go-park-mail-ru/2022_1_Wave/internal/user/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestGetById(t *testing.T) {
	mockUserAgent := new(mocks.UserAgent)
	mockAuthAgent := new(mocks.AuthAgent)
	mockUser := &user_microservice_domain.User{
		ID:             1,
		Username:       "aboba",
		Email:          "aboba@aboba.ru",
		Avatar:         "some_url_to_avatar",
		Password:       "some_password",
		CountFollowing: 0,
	}
	useCaseResultUser := &user_microservice_domain.User{
		ID:             1,
		Username:       "aboba",
		Email:          "aboba@aboba.ru",
		Avatar:         "some_url_to_avatar",
		Password:       "",
		CountFollowing: 0,
	}
	t.Run("success", func(t *testing.T) {
		mockUserAgent.On("GetById", uint(1)).Return(mockUser, nil)
		usecase := user_usecase.NewUserUseCase(mockUserAgent, mockAuthAgent)
		result, err := usecase.GetById(1)

		assert.NoError(t, err)
		assert.Equal(t, result, useCaseResultUser)
	})
	t.Run("error", func(t *testing.T) {
		mockUserAgent.On("GetById", uint(2)).Return(nil, errors.New("error select user"))
		usecase := user_usecase.NewUserUseCase(mockUserAgent, mockAuthAgent)
		result, err := usecase.GetById(2)

		assert.ErrorIs(t, err, user_domain.ErrUserDoesNotExist)
		assert.Nil(t, result)
	})
}

func TestGetByUsername(t *testing.T) {
	mockUserAgent := new(mocks.UserAgent)
	mockAuthAgent := new(mocks.AuthAgent)
	mockUser := &user_microservice_domain.User{
		ID:             1,
		Username:       "aboba",
		Email:          "aboba@aboba.ru",
		Avatar:         "some_url_to_avatar",
		Password:       "some_password",
		CountFollowing: 0,
	}
	useCaseResultUser := &user_microservice_domain.User{
		ID:             1,
		Username:       "aboba",
		Email:          "aboba@aboba.ru",
		Avatar:         "some_url_to_avatar",
		Password:       "",
		CountFollowing: 0,
	}
	t.Run("success", func(t *testing.T) {
		mockUserAgent.On("GetByUsername", mockUser.Username).Return(mockUser, nil)
		usecase := user_usecase.NewUserUseCase(mockUserAgent, mockAuthAgent)
		result, err := usecase.GetByUsername(mockUser.Username)

		assert.NoError(t, err)
		assert.Equal(t, result, useCaseResultUser)
	})
	t.Run("error", func(t *testing.T) {
		mockUserAgent.On("GetByUsername", "doesnt_exist").Return(nil, errors.New("error select user"))
		usecase := user_usecase.NewUserUseCase(mockUserAgent, mockAuthAgent)
		result, err := usecase.GetByUsername("doesnt_exist")

		assert.ErrorIs(t, err, user_domain.ErrUserDoesNotExist)
		assert.Nil(t, result)
	})
}

func TestGetByEmail(t *testing.T) {
	mockUserAgent := new(mocks.UserAgent)
	mockAuthAgent := new(mocks.AuthAgent)
	mockUser := &user_microservice_domain.User{
		ID:             1,
		Username:       "aboba",
		Email:          "aboba@aboba.ru",
		Avatar:         "some_url_to_avatar",
		Password:       "some_password",
		CountFollowing: 0,
	}
	useCaseResultUser := &user_microservice_domain.User{
		ID:             1,
		Username:       "aboba",
		Email:          "aboba@aboba.ru",
		Avatar:         "some_url_to_avatar",
		Password:       "",
		CountFollowing: 0,
	}
	t.Run("success", func(t *testing.T) {
		mockUserAgent.On("GetByEmail", mockUser.Email).Return(mockUser, nil)
		usecase := user_usecase.NewUserUseCase(mockUserAgent, mockAuthAgent)
		result, err := usecase.GetByEmail(mockUser.Email)

		assert.NoError(t, err)
		assert.Equal(t, result, useCaseResultUser)
	})
	t.Run("error", func(t *testing.T) {
		mockUserAgent.On("GetByEmail", "doesnt_exist").Return(nil, errors.New("error select user"))
		usecase := user_usecase.NewUserUseCase(mockUserAgent, mockAuthAgent)
		result, err := usecase.GetByEmail("doesnt_exist")

		assert.ErrorIs(t, err, user_domain.ErrUserDoesNotExist)
		assert.Nil(t, result)
	})
}

func TestGetBySessionId(t *testing.T) {
	mockUserAgent := new(mocks.UserAgent)
	mockAuthAgent := new(mocks.AuthAgent)
	mockUser := &user_microservice_domain.User{
		ID:             1,
		Username:       "aboba",
		Email:          "aboba@aboba.ru",
		Avatar:         "some_url_to_avatar",
		Password:       "some_password",
		CountFollowing: 0,
	}
	useCaseResultUser := &user_microservice_domain.User{
		ID:             1,
		Username:       "aboba",
		Email:          "aboba@aboba.ru",
		Avatar:         "some_url_to_avatar",
		Password:       "",
		CountFollowing: 0,
	}
	sessionResult := &auth_domain.Session{
		UserId:       mockUser.ID,
		IsAuthorized: true,
	}

	t.Run("success", func(t *testing.T) {
		sessionId := "some_session_id"
		mockUserAgent.On("GetById", mockUser.ID).Return(mockUser, nil)
		mockAuthAgent.On("GetSession", sessionId).Return(sessionResult, nil)

		usecase := user_usecase.NewUserUseCase(mockUserAgent, mockAuthAgent)
		result, err := usecase.GetBySessionId(sessionId)

		assert.NoError(t, err)
		assert.Equal(t, result, useCaseResultUser)
	})

	t.Run("error-get-session", func(t *testing.T) {
		sessionId := "some_session_id2"
		mockUserAgent.On("GetById", mockUser.ID).Return(mockUser, nil)
		mockAuthAgent.On("GetSession", sessionId).Return(nil, errors.New("no such session"))

		usecase := user_usecase.NewUserUseCase(mockUserAgent, mockAuthAgent)
		result, err := usecase.GetBySessionId(sessionId)

		assert.ErrorIs(t, err, user_domain.ErrSessionDoesNotExist)
		assert.Nil(t, result)
	})

	t.Run("error-get-user", func(t *testing.T) {
		sessionId := "some_session_id3"
		sessionResult.UserId = 3
		mockUserAgent.On("GetById", sessionResult.UserId).Return(nil, errors.New("no such user"))
		mockAuthAgent.On("GetSession", sessionId).Return(sessionResult, nil)

		usecase := user_usecase.NewUserUseCase(mockUserAgent, mockAuthAgent)
		result, err := usecase.GetBySessionId(sessionId)

		assert.ErrorIs(t, err, user_domain.ErrUserDoesNotExist)
		assert.Nil(t, result)
	})
}

func TestDelete(t *testing.T) {
	mockUserAgent := new(mocks.UserAgent)
	mockAuthAgent := new(mocks.AuthAgent)

	mockUser := &user_microservice_domain.User{
		ID:             1,
		Username:       "aboba",
		Email:          "aboba@aboba.ru",
		Avatar:         "some_url_to_avatar",
		Password:       "some_password",
		CountFollowing: 0,
	}
	mockUser2 := &user_microservice_domain.User{
		ID:             2,
		Username:       "aboba2",
		Email:          "aboba2@aboba.ru",
		Avatar:         "some_url_to_avatar2",
		Password:       "some_password2",
		CountFollowing: 1,
	}

	sessionResult := &auth_domain.Session{
		UserId:       mockUser.ID,
		IsAuthorized: true,
	}

	t.Run("success", func(t *testing.T) {
		mockUserAgent.On("Delete", mockUser.ID).Return(nil)
		mockUserAgent.On("GetByUsername", mockUser.Username).Return(mockUser, nil)
		mockUserAgent.On("GetByEmail", mockUser.Email).Return(mockUser, nil)
		sessionId := "some session id"
		mockAuthAgent.On("GetSession", sessionId).Return(sessionResult, nil)

		usecase := user_usecase.NewUserUseCase(mockUserAgent, mockAuthAgent)

		err := usecase.DeleteById(mockUser.ID)
		assert.NoError(t, err)
		err = usecase.DeleteByUsername(mockUser.Username)
		assert.NoError(t, err)
		err = usecase.DeleteByEmail(mockUser.Email)
		assert.NoError(t, err)
		err = usecase.DeleteBySessionId(sessionId)
		assert.NoError(t, err)
	})
	t.Run("error", func(t *testing.T) {
		mockUserAgent.On("Delete", mockUser2.ID).Return(errors.New("error delete"))
		mockUserAgent.On("GetByUsername", mockUser2.Username).Return(nil, errors.New("error select"))
		mockUserAgent.On("GetByEmail", mockUser2.Email).Return(nil, errors.New("error select"))
		sessionId := "some session id 2"
		mockAuthAgent.On("GetSession", sessionId).Return(nil, errors.New("error get session"))

		usecase := user_usecase.NewUserUseCase(mockUserAgent, mockAuthAgent)

		err := usecase.DeleteById(mockUser2.ID)
		assert.ErrorIs(t, err, user_domain.ErrUserDoesNotExist)
		err = usecase.DeleteByUsername(mockUser2.Username)
		assert.ErrorIs(t, err, user_domain.ErrUserDoesNotExist)
		err = usecase.DeleteByEmail(mockUser2.Email)
		assert.ErrorIs(t, err, user_domain.ErrUserDoesNotExist)
		err = usecase.DeleteBySessionId(sessionId)
		assert.ErrorIs(t, err, user_domain.ErrSessionDoesNotExist)
	})
}

func TestCheckUsernameAndPassword(t *testing.T) {
	mockUserAgent := new(mocks.UserAgent)
	mockAuthAgent := new(mocks.AuthAgent)
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

	t.Run("success", func(t *testing.T) {
		mockUserAgent.On("GetByUsername", mockUser.Username).Return(mockUser, nil)

		usecase := user_usecase.NewUserUseCase(mockUserAgent, mockAuthAgent)
		result := usecase.CheckUsernameAndPassword(mockUser.Username, password)

		assert.True(t, result)
	})
	t.Run("error", func(t *testing.T) {
		mockUserAgent.On("GetByUsername", mockUser.Username).Return(mockUser, nil)
		mockUserAgent.On("GetByUsername", "not such user").Return(nil, errors.New("no such user"))

		usecase := user_usecase.NewUserUseCase(mockUserAgent, mockAuthAgent)
		result1 := usecase.CheckUsernameAndPassword(mockUser.Username, password+"ab")
		result2 := usecase.CheckUsernameAndPassword("not such user", password)

		assert.False(t, result1)
		assert.False(t, result2)
	})
}

func TestCheckEmailAndPassword(t *testing.T) {
	mockUserAgent := new(mocks.UserAgent)
	mockAuthAgent := new(mocks.AuthAgent)
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

	t.Run("success", func(t *testing.T) {
		mockUserAgent.On("GetByEmail", mockUser.Email).Return(mockUser, nil)

		usecase := user_usecase.NewUserUseCase(mockUserAgent, mockAuthAgent)
		result := usecase.CheckEmailAndPassword(mockUser.Email, password)

		assert.True(t, result)
	})
	t.Run("error", func(t *testing.T) {
		mockUserAgent.On("GetByEmail", mockUser.Email).Return(mockUser, nil)
		mockUserAgent.On("GetByEmail", "not such user").Return(nil, errors.New("no such user"))

		usecase := user_usecase.NewUserUseCase(mockUserAgent, mockAuthAgent)
		result1 := usecase.CheckEmailAndPassword(mockUser.Email, password+"ab")
		result2 := usecase.CheckEmailAndPassword("not such user", password)

		assert.False(t, result1)
		assert.False(t, result2)
	})
}

func TestUpdate(t *testing.T) {
	mockUserAgent := new(mocks.UserAgent)
	mockAuthAgent := new(mocks.AuthAgent)

	mockUser := &user_microservice_domain.User{
		ID:             1,
		Username:       "aboba",
		Email:          "aboba@aboba.ru",
		Avatar:         "some_url_to_avatar",
		Password:       "some_password",
		CountFollowing: 0,
	}
	changesToUser := &user_microservice_domain.User{
		ID:             1,
		Username:       "change_aboba",
		Email:          "change_aboba@aboba.ru",
		Avatar:         "change_some_url_to_avatar",
		Password:       "change_some_password",
		CountFollowing: 1,
	}

	t.Run("success", func(t *testing.T) {
		mockUserAgent.On("GetById", mockUser.ID).Return(mockUser, nil)
		mockUserAgent.On("GetByUsername", changesToUser.Username).Return(nil, errors.New("error"))
		mockUserAgent.On("GetByEmail", changesToUser.Email).Return(nil, errors.New("error"))
		mockUserAgent.On("Update", mockUser.ID, mock.Anything).Return(nil)

		usecase := user_usecase.NewUserUseCase(mockUserAgent, mockAuthAgent)
		err := usecase.Update(mockUser.ID, changesToUser)

		assert.Nil(t, err)
	})

	mockUser2 := &user_microservice_domain.User{
		ID:             2,
		Username:       "aboba2",
		Email:          "aboba2@aboba.ru",
		Avatar:         "some_url_to_avatar2",
		Password:       "some_password2",
		CountFollowing: 0,
	}
	changesToUser2 := &user_microservice_domain.User{
		ID:             2,
		Username:       "change_aboba2",
		Email:          "change_aboba2@aboba.ru",
		Avatar:         "change_some_url_to_avatar2",
		Password:       "change_some_password2",
		CountFollowing: 1,
	}

	t.Run("error", func(t *testing.T) {
		mockUserAgent.On("GetById", uint(3)).Return(nil, errors.New("error"))
		mockUserAgent.On("GetById", mockUser2.ID).Return(mockUser2, nil)
		mockUserAgent.On("GetByUsername", changesToUser2.Username).Return(changesToUser2, nil)
		mockUserAgent.On("GetByUsername", "username1").Return(nil, errors.New("error"))
		mockUserAgent.On("GetByEmail", changesToUser2.Email).Return(changesToUser2, nil)

		usecase := user_usecase.NewUserUseCase(mockUserAgent, mockAuthAgent)

		err := usecase.Update(3, changesToUser2)
		assert.ErrorIs(t, err, user_domain.ErrUserDoesNotExist)
		err = usecase.Update(mockUser2.ID, changesToUser2)
		assert.ErrorIs(t, err, user_domain.ErrUserAlreadyExist)

		changesToUser2.Username = "username1"
		err = usecase.Update(mockUser2.ID, changesToUser2)
		assert.ErrorIs(t, err, user_domain.ErrUserAlreadyExist)
	})
}
