package user_grpc_agent

import (
	"context"
	user_microservice_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/user"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/user/proto"
	user_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/user"
)

func getUserForRepo(userProto *proto.User) *user_microservice_domain.User {
	return &user_microservice_domain.User{
		ID:             uint(userProto.GetUserId()),
		Username:       userProto.GetUsername(),
		Email:          userProto.GetEmail(),
		Avatar:         userProto.GetAvatar(),
		Password:       userProto.GetPassword(),
		CountFollowing: int(userProto.GetCountFollowing()),
	}
}

func getProtoUser(user *user_microservice_domain.User) *proto.User {
	return &proto.User{
		UserId:         uint64(user.ID),
		Username:       user.Username,
		Email:          user.Email,
		Avatar:         user.Avatar,
		Password:       user.Password,
		CountFollowing: int64(user.CountFollowing),
	}
}

type userGRPCAgent struct {
	userClient proto.ProfileClient
}

func NewUserGRPCAgent(client proto.ProfileClient) user_domain.UserAgent {
	return &userGRPCAgent{userClient: client}
}

func (a *userGRPCAgent) Create(user *user_microservice_domain.User) error {
	userProto := getProtoUser(user)
	_, err := a.userClient.Create(context.Background(), userProto)

	return err
}

func (a *userGRPCAgent) Update(id uint, user *user_microservice_domain.User) error {
	user.ID = id
	userProto := getProtoUser(user)
	_, err := a.userClient.Update(context.Background(), userProto)

	return err
}

func (a *userGRPCAgent) Delete(id uint) error {
	_, err := a.userClient.Delete(context.Background(), &proto.UserId{UserId: uint64(id)})

	return err
}

func (a *userGRPCAgent) GetById(id uint) (*user_microservice_domain.User, error) {
	userProto, err := a.userClient.GetById(context.Background(), &proto.UserId{UserId: uint64(id)})
	if err != nil {
		return nil, err
	}

	user := getUserForRepo(userProto)
	return user, nil
}

func (a *userGRPCAgent) GetByUsername(username string) (*user_microservice_domain.User, error) {
	userProto, err := a.userClient.GetByUsername(context.Background(), &proto.Username{Username: username})
	if err != nil {
		return nil, err
	}

	user := getUserForRepo(userProto)
	return user, nil
}

func (a *userGRPCAgent) GetByEmail(email string) (*user_microservice_domain.User, error) {
	userProto, err := a.userClient.GetByEmail(context.Background(), &proto.Email{Email: email})
	if err != nil {
		return nil, err
	}

	user := getUserForRepo(userProto)
	return user, nil
}
