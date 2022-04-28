package user_service

import (
	"context"
	user_microservice_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/user"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/user/proto"
)

/*
type ProfileServer interface {
	Create(context.Context, *User) (*Empty, error)
	Update(context.Context, *User) (*Empty, error)
	Delete(context.Context, *UserId) (*Empty, error)
	GetById(context.Context, *UserId) (*User, error)
	GetByUsername(context.Context, *Username) (*User, error)
	GetByEmail(context.Context, *Email) (*User, error)
}
*/

type userService struct {
	userRepo user_microservice_domain.UserRepo
	proto.UnimplementedProfileServer
}

func NewUserService(userRepo user_microservice_domain.UserRepo) proto.ProfileServer {
	return &userService{userRepo: userRepo}
}

func GetUserForRepo(userProto *proto.User) *user_microservice_domain.User {
	return &user_microservice_domain.User{
		Username:       userProto.GetUsername(),
		Email:          userProto.GetEmail(),
		Avatar:         userProto.GetAvatar(),
		Password:       userProto.GetPassword(),
		CountFollowing: int(userProto.GetCountFollowing()),
	}
}

func GetProtoUser(user *user_microservice_domain.User) *proto.User {
	return &proto.User{
		UserId:         uint64(user.ID),
		Username:       user.Username,
		Email:          user.Email,
		Avatar:         user.Avatar,
		Password:       user.Password,
		CountFollowing: int64(user.CountFollowing),
	}
}

func (a *userService) Create(ctx context.Context, user *proto.User) (*proto.Empty, error) {
	userForRepo := GetUserForRepo(user)

	err := a.userRepo.Insert(userForRepo)

	return &proto.Empty{}, err
}

func (a *userService) Update(ctx context.Context, user *proto.User) (*proto.Empty, error) {
	userForRepo := GetUserForRepo(user)

	err := a.userRepo.Update(uint(user.UserId), userForRepo)

	return &proto.Empty{}, err
}

func (a *userService) Delete(ctx context.Context, userId *proto.UserId) (*proto.Empty, error) {
	err := a.userRepo.Delete(uint(userId.GetUserId()))

	return &proto.Empty{}, err
}

func (a *userService) GetById(ctx context.Context, userId *proto.UserId) (*proto.User, error) {
	user, err := a.userRepo.SelectByID(uint(userId.GetUserId()))
	if err != nil {
		return nil, err
	}

	userProto := GetProtoUser(user)

	return userProto, nil
}

func (a *userService) GetByUsername(ctx context.Context, username *proto.Username) (*proto.User, error) {
	user, err := a.userRepo.SelectByUsername(username.GetUsername())
	if err != nil {
		return nil, err
	}

	userProto := GetProtoUser(user)

	return userProto, nil
}

func (a *userService) GetByEmail(ctx context.Context, email *proto.Email) (*proto.User, error) {
	user, err := a.userRepo.SelectByEmail(email.GetEmail())
	if err != nil {
		return nil, err
	}

	userProto := GetProtoUser(user)

	return userProto, nil
}
