package http

import (
	"errors"
	"github.com/go-park-mail-ru/2022_1_Wave/config"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/domain"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type UserHandler struct {
	userUseCase domain.UserUseCase
}

const (
	badIdErr        = "bad id"
	noSessionErr    = "no session"
	invalidUserJSON = "invalid json"
)

func NewUserHandler(e *echo.Echo, userUseCase domain.UserUseCase) {
	handler := &UserHandler{
		userUseCase: userUseCase,
	}

	g := e.Group("/users")

	// TODO: навесить мидлвары
	g.GET("/users/:id", handler.GetUser)
	g.GET("/users/self", handler.GetSelfUser)

	//g.PUT("/users/:id", handler.UpdateUser)
	g.PUT("/users/self", handler.UpdateSelfUser)
}

func (a *UserHandler) GetUser(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, getErrorUserResponse(errors.New(badIdErr)))
	}

	user, err := a.userUseCase.GetById(uint(userId))
	if err != nil {
		return c.JSON(http.StatusNotFound, getErrorUserResponse(err))
	}

	return c.JSON(http.StatusOK, getSuccessGetUserResponse(user))
}

func (a *UserHandler) GetSelfUser(c echo.Context) error {
	cookie, err := c.Cookie(config.C.SessionIDKey)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, getErrorUserResponse(errors.New(noSessionErr)))
	}

	user, err := a.userUseCase.GetBySessionId(cookie.Value)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, getErrorUserResponse(err))
	}

	return c.JSON(http.StatusOK, getSuccessGetUserResponse(user))
}

/*func (a *UserHandler) UpdateUser(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, getErrorUserResponse(errors.New(badIdErr)))
	}

	var user domain.User
	err = c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, getErrorUserResponse(errors.New(invalidUserJSON)))
	}

	err = a.userUseCase.Update(uint(userId), &user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, getErrorUserResponse(err))
	}

	return c.JSON(http.StatusOK, getSuccessUserUpdate(&user))
}*/

func (a *UserHandler) UpdateSelfUser(c echo.Context) error {
	cookie, err := c.Cookie(config.C.SessionIDKey)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, getErrorUserResponse(errors.New(noSessionErr)))
	}

	curUser, err := a.userUseCase.GetBySessionId(cookie.Value)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, getErrorUserResponse(err))
	}

	var userUpdates domain.User
	err = c.Bind(&userUpdates)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, getErrorUserResponse(errors.New(invalidUserJSON)))
	}

	err = a.userUseCase.Update(curUser.ID, &userUpdates)
	if err != nil {
		return c.JSON(http.StatusBadRequest, getErrorUserResponse(err))
	}

	return c.JSON(http.StatusOK, getSuccessUserUpdate(&userUpdates))
}
