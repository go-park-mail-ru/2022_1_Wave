package http

import (
	"errors"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type UserHandler struct {
	UserUseCase domain.UserUseCase
}

const (
	badIdErr          = "bad id"
	noSessionErr      = "no session"
	invalidUserJSON   = "invalid json"
	uploadAvatarError = "upload avatar error"

	SessionIdKey  = "session_id"
	PathToAvatars = "/users/avatars"
)

//func NewUserHandler(e *echo.Echo, userUseCase domain.UserUseCase, m *http_middleware.HttpMiddleware) {
//	handler := &UserHandler{
//		UserUseCase: userUseCase,
//	}
//
//	g := e.Group("/users")
//
//	g.GET("/users/:id", handler.GetUser)
//	g.GET("/users/self", handler.GetSelfUser, m.Auth, m.CSRF)
//
//	//g.PUT("/users/:id", handler.UpdateUser)
//	g.PUT("/users/self", handler.UpdateSelfUser, m.Auth, m.CSRF)
//	g.PUT("/users/upload_avatar", handler.UploadAvatar, m.Auth, m.CSRF)
//}

func (a *UserHandler) GetUser(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, getErrorUserResponse(errors.New(badIdErr)))
	}

	user, err := a.UserUseCase.GetById(uint(userId))
	if err != nil {
		return c.JSON(http.StatusNotFound, getErrorUserResponse(err))
	}

	return c.JSON(http.StatusOK, getSuccessGetUserResponse(user))
}

func (a *UserHandler) GetSelfUser(c echo.Context) error {
	cookie, err := c.Cookie(SessionIdKey)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, getErrorUserResponse(errors.New(noSessionErr)))
	}

	user, err := a.UserUseCase.GetBySessionId(cookie.Value)
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
	cookie, err := c.Cookie(SessionIdKey)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, getErrorUserResponse(errors.New(noSessionErr)))
	}

	curUser, err := a.UserUseCase.GetBySessionId(cookie.Value)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, getErrorUserResponse(err))
	}

	var userUpdates domain.User
	err = c.Bind(&userUpdates)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, getErrorUserResponse(errors.New(invalidUserJSON)))
	}

	err = a.UserUseCase.Update(curUser.ID, &userUpdates)
	if err != nil {
		return c.JSON(http.StatusBadRequest, getErrorUserResponse(err))
	}

	return c.JSON(http.StatusOK, getSuccessUserUpdate())
}

func (a *UserHandler) UploadAvatar(c echo.Context) error {
	file, err := c.FormFile("avatar")
	if err != nil {
		return c.JSON(http.StatusBadRequest, getErrorUserResponse(errors.New(uploadAvatarError)))
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest, getErrorUserResponse(errors.New(uploadAvatarError)))
	}
	defer src.Close()

	cookie, _ := c.Cookie(SessionIdKey)
	user, _ := a.UserUseCase.GetBySessionId(cookie.Value)

	strs := strings.Split(file.Filename, ".")

	filename := PathToAvatars + "/" + strconv.Itoa(int(user.ID)) + "." + strs[len(strs)-1]
	dst, err := os.Create(filename)
	if err != nil {
		return c.JSON(http.StatusBadRequest, getErrorUserResponse(errors.New(uploadAvatarError)))
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusBadRequest, getErrorUserResponse(errors.New(uploadAvatarError)))
	}

	user.Avatar = filename
	_ = a.UserUseCase.Update(user.ID, user)

	return c.JSON(http.StatusOK, getSuccessUploadAvatar())
}
