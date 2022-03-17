package usecase

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal/domain"
)

type authUseCase struct {
	sessionRepo domain.SessionRepo
	userRepo    domain.UserRepo
}

func NewAuthUseCase(sessionRepo domain.SessionRepo, userRepo domain.UserRepo) domain.AuthUseCase {
	return &authUseCase{
		sessionRepo: sessionRepo,
		userRepo:    userRepo,
	}
}

func (a *authUseCase) Login(login string, password string) (string, error) {
	if !a.userRepo.CheckEmailAndPassword(login, password) && !a.userRepo.CheckUsernameAndPassword(login, password) {
		return "", domain.ErrInvalidLoginOrPassword
	}

	user, err := a.userRepo.SelectByUsername(login)
	if err != nil {
		user, _ = a.userRepo.SelectByEmail(login)
	}

	sessionId, err := a.sessionRepo.SetNewSession(user.ID)
	if err != nil {
		return "", domain.ErrWhileSetNewSession
	}

	return sessionId, nil
}

func (a *authUseCase) Logout(sessionId string) error {
	_ = a.sessionRepo.DeleteSession(sessionId)

	return nil
}

func (a *authUseCase) SignUp(user *domain.User) (string, error) {
	_, err := a.userRepo.SelectByEmail(user.Email)
	if err != nil {
		return "", domain.ErrUserAlreadyExist
	}

	_, err = a.userRepo.SelectByUsername(user.Username)
	if err != nil {
		return "", domain.ErrUserAlreadyExist
	}

	err = a.userRepo.Insert(user)
	if err != nil {
		return "", domain.ErrInsert
	}

	userToId, err := a.userRepo.SelectByEmail(user.Email)
	if err != nil {
		return "", domain.ErrDatabaseUnexpected
	}

	sessionId, err := a.sessionRepo.SetNewSession(userToId.ID)
	if err != nil {
		return "", domain.ErrSessionStorageUnexpected
	}

	return sessionId, nil
}
