package usecase

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal/domain"
)

type userUseCase struct {
	userRepo    domain.UserRepo
	sessionRepo domain.SessionRepo
}

func NewUserUseCase(ur domain.UserRepo, sr domain.SessionRepo) domain.UserUseCase {
	return &userUseCase{
		userRepo:    ur,
		sessionRepo: sr,
	}
}

func (a *userUseCase) GetById(userId uint) (*domain.User, error) {
	user, err := a.userRepo.SelectByID(userId)
	if err != nil {
		return nil, domain.ErrUserDoesNotExist
	}

	user.Password = ""

	return user, nil
}

func (a *userUseCase) GetByUsername(username string) (*domain.User, error) {
	user, err := a.userRepo.SelectByUsername(username)
	if err != nil {
		return nil, domain.ErrUserDoesNotExist
	}

	user.Password = ""

	return user, nil
}

func (a *userUseCase) GetByEmail(email string) (*domain.User, error) {
	user, err := a.userRepo.SelectByEmail(email)
	if err != nil {
		return nil, domain.ErrUserDoesNotExist
	}

	user.Password = ""

	return user, nil
}

func (a *userUseCase) GetBySessionId(sessionId string) (*domain.User, error) {
	session, err := a.sessionRepo.GetSession(sessionId)
	if err != nil {
		return nil, domain.ErrSessionDoesNotExist
	}

	user, err := a.userRepo.SelectByID(session.UserId)
	if err != nil {
		return nil, domain.ErrUserDoesNotExist
	}

	user.Password = ""

	return user, nil
}

func (a *userUseCase) DeleteById(userId uint) error {
	err := a.userRepo.Delete(userId)
	if err != nil {
		return domain.ErrUserDoesNotExist
	}

	return nil
}

func (a *userUseCase) DeleteByUsername(username string) error {
	user, err := a.userRepo.SelectByUsername(username)
	if err != nil {
		return domain.ErrUserDoesNotExist
	}

	return a.DeleteById(user.ID)
}

func (a *userUseCase) DeleteByEmail(email string) error {
	user, err := a.userRepo.SelectByEmail(email)
	if err != nil {
		return domain.ErrUserDoesNotExist
	}

	return a.DeleteById(user.ID)
}

func (a *userUseCase) DeleteBySessionId(sessionId string) error {
	session, err := a.sessionRepo.GetSession(sessionId)
	if err != nil {
		return domain.ErrSessionDoesNotExist
	}

	return a.DeleteById(session.UserId)
}

func (a *userUseCase) CheckUsernameAndPassword(username string, password string) bool {
	return a.userRepo.CheckUsernameAndPassword(username, password)
}

func (a *userUseCase) CheckEmailAndPassword(email string, password string) bool {
	return a.userRepo.CheckEmailAndPassword(email, password)
}
