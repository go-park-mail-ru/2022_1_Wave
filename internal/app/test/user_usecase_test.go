package test

import (
	"errors"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain/mocks"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/tools/utils"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/user/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestGetById(t *testing.T) {
	mockUserRepo := new(mocks.UserRepo)
	mockSessionRepo := new(mocks.SessionRepo)
	mockUser := &domain.User{
		ID:             1,
		Username:       "aboba",
		Email:          "aboba@aboba.ru",
		Avatar:         "some_url_to_avatar",
		Password:       "some_password",
		CountFollowing: 0,
	}
	useCaseResultUser := &domain.User{
		ID:             1,
		Username:       "aboba",
		Email:          "aboba@aboba.ru",
		Avatar:         "some_url_to_avatar",
		Password:       "",
		CountFollowing: 0,
	}
	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("SelectByID", uint(1)).Return(mockUser, nil)
		usecase := usecase.NewUserUseCase(mockUserRepo, mockSessionRepo)
		result, err := usecase.GetById(1)

		assert.NoError(t, err)
		assert.Equal(t, result, useCaseResultUser)
	})
	t.Run("error", func(t *testing.T) {
		mockUserRepo.On("SelectByID", uint(2)).Return(nil, errors.New("error select user"))
		usecase := usecase.NewUserUseCase(mockUserRepo, mockSessionRepo)
		result, err := usecase.GetById(2)

		assert.ErrorIs(t, err, domain.ErrUserDoesNotExist)
		assert.Nil(t, result)
	})
}

func TestGetByUsername(t *testing.T) {
	mockUserRepo := new(mocks.UserRepo)
	mockSessionRepo := new(mocks.SessionRepo)
	mockUser := &domain.User{
		ID:             1,
		Username:       "aboba",
		Email:          "aboba@aboba.ru",
		Avatar:         "some_url_to_avatar",
		Password:       "some_password",
		CountFollowing: 0,
	}
	useCaseResultUser := &domain.User{
		ID:             1,
		Username:       "aboba",
		Email:          "aboba@aboba.ru",
		Avatar:         "some_url_to_avatar",
		Password:       "",
		CountFollowing: 0,
	}
	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("SelectByUsername", mockUser.Username).Return(mockUser, nil)
		usecase := usecase.NewUserUseCase(mockUserRepo, mockSessionRepo)
		result, err := usecase.GetByUsername(mockUser.Username)

		assert.NoError(t, err)
		assert.Equal(t, result, useCaseResultUser)
	})
	t.Run("error", func(t *testing.T) {
		mockUserRepo.On("SelectByUsername", "doesnt_exist").Return(nil, errors.New("error select user"))
		usecase := usecase.NewUserUseCase(mockUserRepo, mockSessionRepo)
		result, err := usecase.GetByUsername("doesnt_exist")

		assert.ErrorIs(t, err, domain.ErrUserDoesNotExist)
		assert.Nil(t, result)
	})
}

func TestGetByEmail(t *testing.T) {
	mockUserRepo := new(mocks.UserRepo)
	mockSessionRepo := new(mocks.SessionRepo)
	mockUser := &domain.User{
		ID:             1,
		Username:       "aboba",
		Email:          "aboba@aboba.ru",
		Avatar:         "some_url_to_avatar",
		Password:       "some_password",
		CountFollowing: 0,
	}
	useCaseResultUser := &domain.User{
		ID:             1,
		Username:       "aboba",
		Email:          "aboba@aboba.ru",
		Avatar:         "some_url_to_avatar",
		Password:       "",
		CountFollowing: 0,
	}
	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("SelectByEmail", mockUser.Email).Return(mockUser, nil)
		usecase := usecase.NewUserUseCase(mockUserRepo, mockSessionRepo)
		result, err := usecase.GetByEmail(mockUser.Email)

		assert.NoError(t, err)
		assert.Equal(t, result, useCaseResultUser)
	})
	t.Run("error", func(t *testing.T) {
		mockUserRepo.On("SelectByEmail", "doesnt_exist").Return(nil, errors.New("error select user"))
		usecase := usecase.NewUserUseCase(mockUserRepo, mockSessionRepo)
		result, err := usecase.GetByEmail("doesnt_exist")

		assert.ErrorIs(t, err, domain.ErrUserDoesNotExist)
		assert.Nil(t, result)
	})
}

func TestGetBySessionId(t *testing.T) {
	mockUserRepo := new(mocks.UserRepo)
	mockSessionRepo := new(mocks.SessionRepo)
	mockUser := &domain.User{
		ID:             1,
		Username:       "aboba",
		Email:          "aboba@aboba.ru",
		Avatar:         "some_url_to_avatar",
		Password:       "some_password",
		CountFollowing: 0,
	}
	useCaseResultUser := &domain.User{
		ID:             1,
		Username:       "aboba",
		Email:          "aboba@aboba.ru",
		Avatar:         "some_url_to_avatar",
		Password:       "",
		CountFollowing: 0,
	}
	sessionResult := &domain.Session{
		UserId:       mockUser.ID,
		IsAuthorized: true,
	}

	t.Run("success", func(t *testing.T) {
		sessionId := "some_session_id"
		mockUserRepo.On("SelectByID", mockUser.ID).Return(mockUser, nil)
		mockSessionRepo.On("GetSession", sessionId).Return(sessionResult, nil)

		usecase := usecase.NewUserUseCase(mockUserRepo, mockSessionRepo)
		result, err := usecase.GetBySessionId(sessionId)

		assert.NoError(t, err)
		assert.Equal(t, result, useCaseResultUser)
	})

	t.Run("error-get-session", func(t *testing.T) {
		sessionId := "some_session_id2"
		mockUserRepo.On("SelectByID", mockUser.ID).Return(mockUser, nil)
		mockSessionRepo.On("GetSession", sessionId).Return(nil, errors.New("no such session"))

		usecase := usecase.NewUserUseCase(mockUserRepo, mockSessionRepo)
		result, err := usecase.GetBySessionId(sessionId)

		assert.ErrorIs(t, err, domain.ErrSessionDoesNotExist)
		assert.Nil(t, result)
	})

	t.Run("error-get-user", func(t *testing.T) {
		sessionId := "some_session_id3"
		sessionResult.UserId = 3
		mockUserRepo.On("SelectByID", sessionResult.UserId).Return(nil, errors.New("no such user"))
		mockSessionRepo.On("GetSession", sessionId).Return(sessionResult, nil)

		usecase := usecase.NewUserUseCase(mockUserRepo, mockSessionRepo)
		result, err := usecase.GetBySessionId(sessionId)

		assert.ErrorIs(t, err, domain.ErrUserDoesNotExist)
		assert.Nil(t, result)
	})
}

func TestDelete(t *testing.T) {
	mockUserRepo := new(mocks.UserRepo)
	mockSessionRepo := new(mocks.SessionRepo)

	mockUser := &domain.User{
		ID:             1,
		Username:       "aboba",
		Email:          "aboba@aboba.ru",
		Avatar:         "some_url_to_avatar",
		Password:       "some_password",
		CountFollowing: 0,
	}
	mockUser2 := &domain.User{
		ID:             2,
		Username:       "aboba2",
		Email:          "aboba2@aboba.ru",
		Avatar:         "some_url_to_avatar2",
		Password:       "some_password2",
		CountFollowing: 1,
	}

	sessionResult := &domain.Session{
		UserId:       mockUser.ID,
		IsAuthorized: true,
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("Delete", mockUser.ID).Return(nil)
		mockUserRepo.On("SelectByUsername", mockUser.Username).Return(mockUser, nil)
		mockUserRepo.On("SelectByEmail", mockUser.Email).Return(mockUser, nil)
		sessionId := "some session id"
		mockSessionRepo.On("GetSession", sessionId).Return(sessionResult, nil)

		usecase := usecase.NewUserUseCase(mockUserRepo, mockSessionRepo)

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
		mockUserRepo.On("Delete", mockUser2.ID).Return(errors.New("error delete"))
		mockUserRepo.On("SelectByUsername", mockUser2.Username).Return(nil, errors.New("error select"))
		mockUserRepo.On("SelectByEmail", mockUser2.Email).Return(nil, errors.New("error select"))
		sessionId := "some session id 2"
		mockSessionRepo.On("GetSession", sessionId).Return(nil, errors.New("error get session"))

		usecase := usecase.NewUserUseCase(mockUserRepo, mockSessionRepo)

		err := usecase.DeleteById(mockUser2.ID)
		assert.ErrorIs(t, err, domain.ErrUserDoesNotExist)
		err = usecase.DeleteByUsername(mockUser2.Username)
		assert.ErrorIs(t, err, domain.ErrUserDoesNotExist)
		err = usecase.DeleteByEmail(mockUser2.Email)
		assert.ErrorIs(t, err, domain.ErrUserDoesNotExist)
		err = usecase.DeleteBySessionId(sessionId)
		assert.ErrorIs(t, err, domain.ErrSessionDoesNotExist)
	})
}

func TestCheckUsernameAndPassword(t *testing.T) {
	mockUserRepo := new(mocks.UserRepo)
	mockSessionRepo := new(mocks.SessionRepo)
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

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("SelectByUsername", mockUser.Username).Return(mockUser, nil)

		usecase := usecase.NewUserUseCase(mockUserRepo, mockSessionRepo)
		result := usecase.CheckUsernameAndPassword(mockUser.Username, password)

		assert.True(t, result)
	})
	t.Run("error", func(t *testing.T) {
		mockUserRepo.On("SelectByUsername", mockUser.Username).Return(mockUser, nil)
		mockUserRepo.On("SelectByUsername", "not such user").Return(nil, errors.New("no such user"))

		usecase := usecase.NewUserUseCase(mockUserRepo, mockSessionRepo)
		result1 := usecase.CheckUsernameAndPassword(mockUser.Username, password+"ab")
		result2 := usecase.CheckUsernameAndPassword("not such user", password)

		assert.False(t, result1)
		assert.False(t, result2)
	})
}

func TestCheckEmailAndPassword(t *testing.T) {
	mockUserRepo := new(mocks.UserRepo)
	mockSessionRepo := new(mocks.SessionRepo)
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

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("SelectByEmail", mockUser.Email).Return(mockUser, nil)

		usecase := usecase.NewUserUseCase(mockUserRepo, mockSessionRepo)
		result := usecase.CheckEmailAndPassword(mockUser.Email, password)

		assert.True(t, result)
	})
	t.Run("error", func(t *testing.T) {
		mockUserRepo.On("SelectByEmail", mockUser.Email).Return(mockUser, nil)
		mockUserRepo.On("SelectByEmail", "not such user").Return(nil, errors.New("no such user"))

		usecase := usecase.NewUserUseCase(mockUserRepo, mockSessionRepo)
		result1 := usecase.CheckEmailAndPassword(mockUser.Email, password+"ab")
		result2 := usecase.CheckEmailAndPassword("not such user", password)

		assert.False(t, result1)
		assert.False(t, result2)
	})
}

func TestUpdate(t *testing.T) {
	mockUserRepo := new(mocks.UserRepo)
	mockSessionRepo := new(mocks.SessionRepo)

	mockUser := &domain.User{
		ID:             1,
		Username:       "aboba",
		Email:          "aboba@aboba.ru",
		Avatar:         "some_url_to_avatar",
		Password:       "some_password",
		CountFollowing: 0,
	}
	changesToUser := &domain.User{
		ID:             1,
		Username:       "change_aboba",
		Email:          "change_aboba@aboba.ru",
		Avatar:         "change_some_url_to_avatar",
		Password:       "change_some_password",
		CountFollowing: 1,
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("SelectByID", mockUser.ID).Return(mockUser, nil)
		mockUserRepo.On("SelectByUsername", changesToUser.Username).Return(nil, errors.New("error"))
		mockUserRepo.On("SelectByEmail", changesToUser.Email).Return(nil, errors.New("error"))
		mockUserRepo.On("Update", mockUser.ID, mock.Anything).Return(nil)

		usecase := usecase.NewUserUseCase(mockUserRepo, mockSessionRepo)
		err := usecase.Update(mockUser.ID, changesToUser)

		assert.Nil(t, err)
	})

	mockUser2 := &domain.User{
		ID:             2,
		Username:       "aboba2",
		Email:          "aboba2@aboba.ru",
		Avatar:         "some_url_to_avatar2",
		Password:       "some_password2",
		CountFollowing: 0,
	}
	changesToUser2 := &domain.User{
		ID:             2,
		Username:       "change_aboba2",
		Email:          "change_aboba2@aboba.ru",
		Avatar:         "change_some_url_to_avatar2",
		Password:       "change_some_password2",
		CountFollowing: 1,
	}

	t.Run("error", func(t *testing.T) {
		mockUserRepo.On("SelectByID", uint(3)).Return(nil, errors.New("error"))
		mockUserRepo.On("SelectByID", mockUser2.ID).Return(mockUser2, nil)
		mockUserRepo.On("SelectByUsername", changesToUser2.Username).Return(changesToUser2, nil)
		mockUserRepo.On("SelectByUsername", "username1").Return(nil, errors.New("error"))
		mockUserRepo.On("SelectByEmail", changesToUser2.Email).Return(changesToUser2, nil)

		usecase := usecase.NewUserUseCase(mockUserRepo, mockSessionRepo)

		err := usecase.Update(3, changesToUser2)
		assert.ErrorIs(t, err, domain.ErrUserDoesNotExist)
		err = usecase.Update(mockUser2.ID, changesToUser2)
		assert.ErrorIs(t, err, domain.ErrUserAlreadyExist)

		changesToUser2.Username = "username1"
		err = usecase.Update(mockUser2.ID, changesToUser2)
		assert.ErrorIs(t, err, domain.ErrUserAlreadyExist)
	})
}
