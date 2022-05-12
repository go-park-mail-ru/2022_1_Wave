package user_service

import (
	"context"
	"errors"
	"github.com/bxcodec/faker"
	user_microservice_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/user"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/user/mocks"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/user/proto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreate(t *testing.T) {
	var mockUser user_microservice_domain.User
	err := faker.FakeData(&mockUser)
	assert.NoError(t, err)
	mockUser.ID = 0

	mockUserRepo := new(mocks.UserRepo)
	mockUserRepo.On("Insert", &mockUser).Return(nil)

	userService := NewUserService(mockUserRepo)

	_, err = userService.Create(context.Background(), GetProtoUser(&mockUser))
	assert.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	var mockUser user_microservice_domain.User
	err := faker.FakeData(&mockUser)
	assert.NoError(t, err)

	mockUserRepo := new(mocks.UserRepo)
	mockUserRepo.On("Update", mockUser.ID, &mockUser).Return(nil)

	userService := NewUserService(mockUserRepo)

	_, err = userService.Update(context.Background(), GetProtoUser(&mockUser))
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	mockUserRepo := new(mocks.UserRepo)
	mockUserRepo.On("Delete", uint(2)).Return(nil)

	userService := NewUserService(mockUserRepo)

	_, err := userService.Delete(context.Background(), &proto.UserId{UserId: 2})
	assert.NoError(t, err)
}

func TestGetById(t *testing.T) {
	var mockUser user_microservice_domain.User
	err := faker.FakeData(&mockUser)
	assert.NoError(t, err)
	mockUser.ID = 2

	mockUserRepo := new(mocks.UserRepo)
	mockUserRepo.On("SelectByID", uint(2)).Return(&mockUser, nil)
	mockUserRepo.On("SelectByID", uint(3)).Return(nil, errors.New("error"))

	userService := NewUserService(mockUserRepo)
	userProto, err := userService.GetById(context.Background(), &proto.UserId{UserId: 2})
	assert.NoError(t, err)

	userForRepo := GetUserForRepo(userProto)
	userForRepo.ID = mockUser.ID
	assert.Equal(t, *userForRepo, mockUser)

	failRes, err := userService.GetById(context.Background(), &proto.UserId{UserId: 3})
	assert.Nil(t, failRes)
	assert.Error(t, err)
}

func TestGetByUsername(t *testing.T) {
	var mockUser user_microservice_domain.User
	err := faker.FakeData(&mockUser)
	assert.NoError(t, err)
	mockUser.ID = 2

	mockUserRepo := new(mocks.UserRepo)
	mockUserRepo.On("SelectByUsername", mockUser.Username).Return(&mockUser, nil)
	mockUserRepo.On("SelectByUsername", mockUser.Username+"a").Return(nil, errors.New("error"))

	userService := NewUserService(mockUserRepo)
	userProto, err := userService.GetByUsername(context.Background(), &proto.Username{Username: mockUser.Username})
	assert.NoError(t, err)

	userForRepo := GetUserForRepo(userProto)
	userForRepo.ID = mockUser.ID
	assert.Equal(t, *userForRepo, mockUser)

	failRes, err := userService.GetByUsername(context.Background(), &proto.Username{Username: mockUser.Username + "a"})
	assert.Nil(t, failRes)
	assert.Error(t, err)
}

func TestGetByEmail(t *testing.T) {
	var mockUser user_microservice_domain.User
	err := faker.FakeData(&mockUser)
	assert.NoError(t, err)
	mockUser.ID = 2

	mockUserRepo := new(mocks.UserRepo)
	mockUserRepo.On("SelectByEmail", mockUser.Email).Return(&mockUser, nil)
	mockUserRepo.On("SelectByEmail", mockUser.Email+"a").Return(nil, errors.New("error"))

	userService := NewUserService(mockUserRepo)
	userProto, err := userService.GetByEmail(context.Background(), &proto.Email{Email: mockUser.Email})
	assert.NoError(t, err)

	userForRepo := GetUserForRepo(userProto)
	userForRepo.ID = mockUser.ID
	assert.Equal(t, *userForRepo, mockUser)

	failRes, err := userService.GetByEmail(context.Background(), &proto.Email{Email: mockUser.Email + "a"})
	assert.Nil(t, failRes)
	assert.Error(t, err)
}
