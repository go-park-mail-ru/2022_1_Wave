package auth_usecase

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	auth_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/app/microservices/auth"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/microservices/auth/proto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/tools/utils"
	"time"
)

/*
type AuthorizationServer interface {
	Login(context.Context, *LoginData) (*LoginResult, error)
	Logout(context.Context, *Session) (*LogoutResult, error)
	SignUp(context.Context, *SignUpData) (*SignUpResult, error)
	GetUnauthorizedSession(context.Context, *Empty) (*GetSessionResult, error)
	IsSession(context.Context, *Session) (*Empty, error)
	IsAuthSession(context.Context, *Session) (*Empty, error)
}
*/

var SessionExpire = time.Hour * 24

type authService struct {
	authRepo auth_domain.AuthRepo
	userRepo domain.UserRepo
}

func NewAuthService(authRepo auth_domain.AuthRepo, userRepo domain.UserRepo) proto.AuthorizationServer {
	return &authService{authRepo: authRepo, userRepo: userRepo}
}

func (a *authService) Login(ctx context.Context, loginData *proto.LoginData) (*proto.LoginResult, error) {
	var user *domain.User
	user, err := a.userRepo.SelectByUsername(loginData.Login)

	if err != nil {
		user, err = a.userRepo.SelectByEmail(loginData.Login)
		if err != nil {
			return nil, errors.New(auth_domain.ErrUserDoesNotExist)
		}
	}
	if !utils.CheckPassword(user.Password, loginData.Password) {
		return nil, errors.New(auth_domain.ErrInvalidLoginOrPassword)
	}

	sessionId, err := a.authRepo.SetNewAuthorizedSession(user.ID, SessionExpire)

	if err != nil {
		return nil, errors.New(auth_domain.ErrWhileChangeSession)
	}

	var response proto.LoginResult
	response.NewSession = &proto.Session{}
	response.NewSession.SessionId = sessionId

	return &response, nil
}

func (a *authService) Logout(ctx context.Context, session *proto.Session) (*proto.Empty, error) {
	err := a.authRepo.DeleteSession(session.SessionId)

	if err != nil {
		return nil, err
	}

	return &proto.Empty{}, nil
}

func (a *authService) SignUp(ctx context.Context, data *proto.SignUpData) (*proto.SignUpResult, error) {
	_, err := a.userRepo.SelectByEmail(data.User.Email)
	if err == nil {
		return nil, errors.New(auth_domain.ErrUserAlreadyExist)
	}

	_, err = a.userRepo.SelectByUsername(data.User.Username)
	if err == nil {
		return nil, errors.New(auth_domain.ErrUserAlreadyExist)
	}

	passwordHash, _ := utils.GetPasswordHash(data.User.Password)

	var user domain.User
	user.Username = data.User.Username
	user.Email = data.User.Email

	user.Password = string(passwordHash)

	err = a.userRepo.Insert(&user)
	if err != nil {
		return nil, errors.New(auth_domain.ErrInsert)
	}

	userToId, err := a.userRepo.SelectByEmail(user.Email)
	if err != nil {
		return nil, errors.New(auth_domain.ErrDatabaseUnexpected)
	}

	sessionId, err := a.authRepo.SetNewAuthorizedSession(userToId.ID, SessionExpire)

	if err != nil {
		return nil, errors.New(auth_domain.ErrWhileChangeSession)
	}

	var response proto.SignUpResult
	response.NewSession = &proto.Session{SessionId: sessionId}

	return &response, nil
}

func (a *authService) GetUnauthorizedSession(ctx context.Context, empty *proto.Empty) (*proto.GetSessionResult, error) {
	sessionId, err := a.authRepo.SetNewUnauthorizedSession(SessionExpire)
	if err != nil {
		return nil, err
	}
	var response proto.GetSessionResult
	response.Session = &proto.Session{SessionId: sessionId}

	return &response, nil
}

func (a *authService) GetAuthorizedSession(ctx context.Context, userId *proto.UserId) (*proto.GetSessionResult, error) {
	sessionId, err := a.authRepo.SetNewAuthorizedSession(uint(userId.UserId), SessionExpire)
	if err != nil {
		return nil, err
	}
	var response proto.GetSessionResult
	response.Session = &proto.Session{SessionId: sessionId}

	return &response, nil
}

func (a *authService) IsSession(ctx context.Context, session *proto.Session) (*proto.Empty, error) {
	_, err := a.authRepo.GetSession(session.SessionId)
	return &proto.Empty{}, err
}

func (a *authService) IsAuthSession(ctx context.Context, session *proto.Session) (*proto.Empty, error) {
	sessionFromRepo, err := a.authRepo.GetSession(session.SessionId)
	if err != nil || sessionFromRepo == nil || !sessionFromRepo.IsAuthorized {
		return &proto.Empty{}, errors.New(auth_domain.ErrNotAuthSession)
	}

	return &proto.Empty{}, nil
}
