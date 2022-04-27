package UserUsecase

import (
	auth_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/auth"
	user_microservice_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/user"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/tools/utils"
	user_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/user"
)

type userUseCase struct {
	userAgent    user_domain.UserAgent
	sessionAgent auth_domain.AuthAgent
}

func NewUserUseCase(userAgent user_domain.UserAgent, sessionAgent auth_domain.AuthAgent) user_domain.UserUseCase {
	return &userUseCase{
		userAgent:    userAgent,
		sessionAgent: sessionAgent,
	}
}

func (a *userUseCase) GetById(userId uint) (*user_microservice_domain.User, error) {
	user, err := a.userAgent.GetById(userId)
	if err != nil {
		return nil, user_domain.ErrUserDoesNotExist
	}

	user.Password = ""

	return user, nil
}

func (a *userUseCase) GetByUsername(username string) (*user_microservice_domain.User, error) {
	user, err := a.userAgent.GetByUsername(username)
	if err != nil {
		return nil, user_domain.ErrUserDoesNotExist
	}

	user.Password = ""

	return user, nil
}

func (a *userUseCase) GetByEmail(email string) (*user_microservice_domain.User, error) {
	user, err := a.userAgent.GetByEmail(email)
	if err != nil {
		return nil, user_domain.ErrUserDoesNotExist
	}

	user.Password = ""

	return user, nil
}

func (a *userUseCase) GetBySessionId(sessionId string) (*user_microservice_domain.User, error) {
	session, err := a.sessionAgent.GetSession(sessionId)
	if err != nil {
		return nil, user_domain.ErrSessionDoesNotExist
	}

	user, err := a.userAgent.GetById(session.UserId)
	if err != nil {
		return nil, user_domain.ErrUserDoesNotExist
	}

	user.Password = ""

	return user, nil
}

func (a *userUseCase) DeleteById(userId uint) error {
	err := a.userAgent.Delete(userId)
	if err != nil {
		return user_domain.ErrUserDoesNotExist
	}

	return nil
}

func (a *userUseCase) DeleteByUsername(username string) error {
	user, err := a.userAgent.GetByUsername(username)
	if err != nil {
		return user_domain.ErrUserDoesNotExist
	}

	return a.DeleteById(user.ID)
}

func (a *userUseCase) DeleteByEmail(email string) error {
	user, err := a.userAgent.GetByEmail(email)
	if err != nil {
		return user_domain.ErrUserDoesNotExist
	}

	return a.DeleteById(user.ID)
}

func (a *userUseCase) DeleteBySessionId(sessionId string) error {
	session, err := a.sessionAgent.GetSession(sessionId)
	if err != nil {
		return user_domain.ErrSessionDoesNotExist
	}

	return a.DeleteById(session.UserId)
}

func (a *userUseCase) CheckUsernameAndPassword(username string, password string) bool {
	user, err := a.userAgent.GetByUsername(username)
	if err != nil {
		return false
	}

	return utils.CheckPassword(user.Password, password)
}

func (a *userUseCase) CheckEmailAndPassword(email string, password string) bool {
	user, err := a.userAgent.GetByEmail(email)
	if err != nil {
		return false
	}

	return utils.CheckPassword(user.Password, password)
}

func (a *userUseCase) Update(id uint, user *user_microservice_domain.User) error {
	curUser, err := a.userAgent.GetById(id)
	if err != nil {
		return user_domain.ErrUserDoesNotExist
	}

	_, err = a.userAgent.GetByUsername(user.Username)
	if err == nil && curUser.Username != user.Username {
		return user_domain.ErrUserAlreadyExist
	}

	_, err = a.userAgent.GetByEmail(user.Email)
	if err == nil && curUser.Email != user.Email {
		return user_domain.ErrUserAlreadyExist
	}

	if user.Password != "" {
		passwordHash, _ := utils.GetPasswordHash(user.Password)
		user.Password = string(passwordHash)
	}

	return a.userAgent.Update(id, user)
}
