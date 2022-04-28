package user_grpc_agent_test

import (
	"errors"
	"github.com/bxcodec/faker"
	user_microservice_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/user"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/user/mocks"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/user/proto"
	user_service "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/user/service"
	user_grpc_agent "github.com/go-park-mail-ru/2022_1_Wave/internal/user/client/grpc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreate(t *testing.T) {
	var mockUser user_microservice_domain.User
	err := faker.FakeData(&mockUser)
	assert.NoError(t, err)
	protoUser := user_service.GetProtoUser(&mockUser)

	mockProfileClient := mocks.ProfileClient{}
	mockProfileClient.On("Create", mock.Anything, protoUser).Return(nil, nil)
	userAgent := user_grpc_agent.NewUserGRPCAgent(&mockProfileClient)

	err = userAgent.Create(&mockUser)
	assert.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	var mockUser user_microservice_domain.User
	err := faker.FakeData(&mockUser)
	assert.NoError(t, err)
	protoUser := user_service.GetProtoUser(&mockUser)

	mockProfileClient := mocks.ProfileClient{}
	mockProfileClient.On("Update", mock.Anything, protoUser).Return(nil, nil)
	userAgent := user_grpc_agent.NewUserGRPCAgent(&mockProfileClient)

	err = userAgent.Update(mockUser.ID, &mockUser)
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	var mockUser user_microservice_domain.User
	err := faker.FakeData(&mockUser)
	assert.NoError(t, err)

	mockProfileClient := mocks.ProfileClient{}
	mockProfileClient.On("Delete", mock.Anything, &proto.UserId{UserId: uint64(mockUser.ID)}).Return(nil, nil)
	userAgent := user_grpc_agent.NewUserGRPCAgent(&mockProfileClient)

	err = userAgent.Delete(mockUser.ID)
	assert.NoError(t, err)
}

func TestGetById(t *testing.T) {
	var mockUser user_microservice_domain.User
	err := faker.FakeData(&mockUser)
	assert.NoError(t, err)
	protoUser := user_service.GetProtoUser(&mockUser)

	mockProfileClient := mocks.ProfileClient{}
	mockProfileClient.On("GetById", mock.Anything, &proto.UserId{UserId: uint64(mockUser.ID)}).Return(protoUser, nil)
	mockProfileClient.On("GetById", mock.Anything, &proto.UserId{UserId: uint64(mockUser.ID + 1)}).Return(nil, errors.New("error"))
	userAgent := user_grpc_agent.NewUserGRPCAgent(&mockProfileClient)

	userResult, err := userAgent.GetById(mockUser.ID)
	assert.NoError(t, err)
	assert.Equal(t, *userResult, mockUser)

	userResult, err = userAgent.GetById(mockUser.ID + 1)
	assert.Error(t, err)
	assert.Nil(t, userResult)
}

func TestGetByUsername(t *testing.T) {
	var mockUser user_microservice_domain.User
	err := faker.FakeData(&mockUser)
	assert.NoError(t, err)
	protoUser := user_service.GetProtoUser(&mockUser)

	mockProfileClient := mocks.ProfileClient{}
	mockProfileClient.On("GetByUsername", mock.Anything, &proto.Username{Username: mockUser.Username}).Return(protoUser, nil)
	mockProfileClient.On("GetByUsername", mock.Anything, &proto.Username{Username: mockUser.Username + "a"}).Return(nil, errors.New("error"))
	userAgent := user_grpc_agent.NewUserGRPCAgent(&mockProfileClient)

	userResult, err := userAgent.GetByUsername(mockUser.Username)
	assert.NoError(t, err)
	assert.Equal(t, *userResult, mockUser)

	userResult, err = userAgent.GetByUsername(mockUser.Username + "a")
	assert.Error(t, err)
	assert.Nil(t, userResult)
}

func TestGetByEmail(t *testing.T) {
	var mockUser user_microservice_domain.User
	err := faker.FakeData(&mockUser)
	assert.NoError(t, err)
	protoUser := user_service.GetProtoUser(&mockUser)

	mockProfileClient := mocks.ProfileClient{}
	mockProfileClient.On("GetByEmail", mock.Anything, &proto.Email{Email: mockUser.Email}).Return(protoUser, nil)
	mockProfileClient.On("GetByEmail", mock.Anything, &proto.Email{Email: mockUser.Email + "a"}).Return(nil, errors.New("error"))
	userAgent := user_grpc_agent.NewUserGRPCAgent(&mockProfileClient)

	userResult, err := userAgent.GetByEmail(mockUser.Email)
	assert.NoError(t, err)
	assert.Equal(t, *userResult, mockUser)

	userResult, err = userAgent.GetByEmail(mockUser.Email + "a")
	assert.Error(t, err)
	assert.Nil(t, userResult)
}
