package AuthUseCase

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/tools/utils"
	"time"
)

type authUseCase struct {
	sessionRepo domain.SessionRepo
	userRepo    domain.UserRepo
}

var SessionExpire = time.Hour * 24

func NewAuthUseCase(sessionRepo domain.SessionRepo, userRepo domain.UserRepo) domain.AuthUseCase {
	return &authUseCase{
		sessionRepo: sessionRepo,
		userRepo:    userRepo,
	}
}

func (a *authUseCase) Login(login string, password string, sessionId string) error {
	var user *domain.User
	user, err := a.userRepo.SelectByUsername(login)
	if err != nil {
		user, err = a.userRepo.SelectByEmail(login)
		if err != nil {
			return domain.ErrUserDoesNotExist
		}
	}
	if !utils.CheckPassword(user.Password, password) {
		return domain.ErrInvalidLoginOrPassword
	}

	err = a.sessionRepo.MakeSessionAuthorized(sessionId, user.ID)

	if err != nil {
		return domain.ErrWhileChangeSession
	}

	return nil
}

func (a *authUseCase) Logout(sessionId string) error {
	return a.sessionRepo.MakeSessionUnauthorized(sessionId)
}

func (a *authUseCase) SignUp(user *domain.User, sessionId string) error {
	_, err := a.userRepo.SelectByEmail(user.Email)
	if err == nil {
		return domain.ErrUserAlreadyExist
	}

	_, err = a.userRepo.SelectByUsername(user.Username)
	if err == nil {
		return domain.ErrUserAlreadyExist
	}

	passwordHash, _ := utils.GetPasswordHash(user.Password)

	user.Password = string(passwordHash)

	err = a.userRepo.Insert(user)
	if err != nil {
		return domain.ErrInsert
	}

	userToId, err := a.userRepo.SelectByEmail(user.Email)
	if err != nil {
		return domain.ErrDatabaseUnexpected
	}

	err = a.sessionRepo.MakeSessionAuthorized(sessionId, userToId.ID)

	if err != nil {
		return domain.ErrWhileChangeSession
	}

	return nil
}

func (a *authUseCase) GetUnauthorizedSession() (string, error) {
	return a.sessionRepo.SetNewUnauthorizedSession(SessionExpire)
}

func (a *authUseCase) IsSession(sessionId string) bool {
	_, err := a.sessionRepo.GetSession(sessionId)
	return err == nil
}

func (a *authUseCase) IsAuthSession(sessionId string) bool {
	session, err := a.sessionRepo.GetSession(sessionId)
	if err != nil || session == nil {
		return false
	}

	return session.IsAuthorized
}
