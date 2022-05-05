package internal

import (
	"errors"
	user_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/user"
	"github.com/go-park-mail-ru/2022_1_Wave/pkg/webUtils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func GetUserId(ctx echo.Context, userUseCase user_domain.UserUseCase) (int64, error) {
	cookie, err := ctx.Cookie(SessionIdKey)
	if err != nil {
		return -1, err
	}

	user, err := userUseCase.GetBySessionId(cookie.Value)
	if err != nil {
		return -1, err
	}

	userId := int64(user.ID)

	return userId, nil
}

func UnauthorizedError(ctx echo.Context) error {
	return webUtils.WriteErrorEchoServer(ctx, errors.New(Unauthorized), http.StatusUnauthorized)
}

func GetIdInt64ByFieldId(ctx echo.Context) (int64, error) {
	return strconv.ParseInt(ctx.Param(FieldId), 10, 64)
}
