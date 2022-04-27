package trackDeliveryHttp

import (
	"errors"
	"fmt"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	Gateway "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/gateway"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/gateway/gatewayProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/playlist/playlistProto"
	PlaylistUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/playlist/useCase"
	user_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/user"
	"github.com/go-park-mail-ru/2022_1_Wave/pkg/webUtils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Handler struct {
	PlaylistUseCase PlaylistUseCase.PlaylistAgent
	UserUseCase     user_domain.UserUseCase
}

func MakeHandler(playlist PlaylistUseCase.PlaylistAgent, user user_domain.UserUseCase) Handler {
	return Handler{
		PlaylistUseCase: playlist,
		UserUseCase:     user,
	}
}

// GetAll godoc
// @Summary      GetAll
// @Description  getting all playlists
// @Tags         playlist
// @Accept          application/json
// @Produce      application/json
// @Success      200  {object}  webUtils.Success
// @Failure      400  {object}  webUtils.Error  "Data is invalid"
// @Failure      405  {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/playlists/ [get]
func (h Handler) GetAll(ctx echo.Context) error {
	cookie, err := ctx.Cookie(constants.SessionIdKey)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, err)
	}
	user, err := h.UserUseCase.GetBySessionId(cookie.Value)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, err)
	}

	playlists, err := h.PlaylistUseCase.GetAll(&gatewayProto.IdArg{Id: int64(user.ID)})
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	if playlists == nil {
		playlists = &playlistProto.PlaylistsResponse{}
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: playlists})
}

// Create godoc
// @Summary      Create
// @Description  creating new playlist
// @Tags         playlist
// @Accept          application/json
// @Produce      application/json
// @Success      200    {object}  webUtils.Success
// @Failure      400    {object}  webUtils.Error  "Data is invalid"
// @Failure      405    {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/playlists/ [post]
func (h Handler) Create(ctx echo.Context) error {
	cookie, err := ctx.Cookie(constants.SessionIdKey)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, err)
	}
	user, err := h.UserUseCase.GetBySessionId(cookie.Value)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, err)
	}

	toCreate := playlistProto.Playlist{}

	if err := ctx.Bind(&toCreate); err != nil {
		return err
	}

	if err := Gateway.Check(&toCreate); err != nil {
		return err
	}

	if err := h.PlaylistUseCase.Create(&playlistProto.UserIdPlaylistArg{
		UserId:   int64(user.ID),
		Playlist: &toCreate,
	}); err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	lastId, err := h.PlaylistUseCase.GetLastId(&gatewayProto.IdArg{Id: int64(user.ID)})
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: constants.SuccessCreated + "(" + fmt.Sprint(lastId) + ")"})
}

// Update godoc
// @Summary      Update
// @Description  updating playlist by id
// @Tags         playlist
// @Accept          application/json
// @Produce      application/json
// @Param        playlist  body      playlistProto.Playlist  true  "id of updating playlist and params of it."
// @Success      200    {object}  webUtils.Success
// @Failure      400    {object}  webUtils.Error  "Data is invalid"
// @Failure      405    {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/playlists/ [put]
func (h Handler) Update(ctx echo.Context) error {
	cookie, err := ctx.Cookie(constants.SessionIdKey)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, err)
	}
	user, err := h.UserUseCase.GetBySessionId(cookie.Value)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, err)
	}

	toUpdate := playlistProto.Playlist{}

	if err := ctx.Bind(&toUpdate); err != nil {
		return err
	}

	if err := Gateway.Check(&toUpdate); err != nil {
		return err
	}

	if err := h.PlaylistUseCase.Update(&playlistProto.UserIdPlaylistArg{
		UserId:   int64(user.ID),
		Playlist: &toUpdate,
	}); err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	id := toUpdate.Id
	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: constants.SuccessUpdated + "(" + fmt.Sprint(id) + ")"})
}

// Get godoc
// @Summary      Get
// @Description  getting playlist by id
// @Tags         playlist
// @Accept          application/json
// @Produce      application/json
// @Param        id   path      integer  true  "id of playlist which need to be getted"
// @Success      200  {object}  webUtils.Success
// @Failure      400  {object}  webUtils.Error  "Data is invalid"
// @Failure      405  {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/playlists/{id} [get]
func (h Handler) Get(ctx echo.Context) error {
	cookie, err := ctx.Cookie(constants.SessionIdKey)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, err)
	}
	user, err := h.UserUseCase.GetBySessionId(cookie.Value)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, err)
	}

	id, err := strconv.Atoi(ctx.Param(constants.FieldId))
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}
	if id < 0 {
		return webUtils.WriteErrorEchoServer(ctx, errors.New(constants.IndexOutOfRange), http.StatusBadRequest)
	}
	playlist, err := h.PlaylistUseCase.GetById(&playlistProto.UserIdPlaylistIdArg{
		UserId:     int64(user.ID),
		PlaylistId: int64(id),
	})

	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: playlist})
}

// Delete godoc
// @Summary      Delete
// @Description  deleting playlists by id
// @Tags         playlists
// @Accept       application/json
// @Produce      application/json
// @Param        id   path      integer  true  "id of playlists which need to be deleted"
// @Success      200  {object}  webUtils.Success
// @Failure      400  {object}  webUtils.Error  "Data is invalid"
// @Failure      405  {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/playlists/{id} [delete]
func (h Handler) Delete(ctx echo.Context) error {
	cookie, err := ctx.Cookie(constants.SessionIdKey)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, err)
	}
	user, err := h.UserUseCase.GetBySessionId(cookie.Value)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, err)
	}

	id, err := strconv.Atoi(ctx.Param(constants.FieldId))
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}
	if id < 0 {
		return webUtils.WriteErrorEchoServer(ctx, errors.New(constants.IndexOutOfRange), http.StatusBadRequest)
	}

	if err := h.PlaylistUseCase.Delete(&playlistProto.UserIdPlaylistIdArg{
		UserId:     int64(user.ID),
		PlaylistId: int64(id),
	}); err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: constants.SuccessDeleted + "(" + fmt.Sprint(id) + ")"})
}
