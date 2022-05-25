package playlistDeliveryHttp

import (
	"errors"
	"fmt"
	internal "github.com/go-park-mail-ru/2022_1_Wave/internal"
	Gateway "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/gateway"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/playlist/playlistProto"
	PlaylistUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/playlist/useCase"
	user_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/user"
	"github.com/go-park-mail-ru/2022_1_Wave/pkg/webUtils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Handler struct {
	PlaylistUseCase PlaylistUseCase.PlaylistUseCase
	UserUseCase     user_domain.UserUseCase
}

func MakeHandler(playlist PlaylistUseCase.PlaylistUseCase, user user_domain.UserUseCase) Handler {
	return Handler{
		PlaylistUseCase: playlist,
		UserUseCase:     user,
	}
}

func toMap(playlists []*playlistProto.PlaylistDataTransfer) map[int64]*playlistProto.PlaylistDataTransfer {
	playlistMap := map[int64]*playlistProto.PlaylistDataTransfer{}
	for _, obj := range playlists {
		playlistMap[obj.Id] = obj
	}
	return playlistMap
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
	userId, err := internal.GetUserId(ctx, h.UserUseCase)
	if err != nil {
		userId = -1
	}
	playlists, err := h.PlaylistUseCase.GetAll(userId)
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	if playlists == nil {
		playlists = []*playlistProto.PlaylistDataTransfer{}
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: toMap(playlists)})
}

// GetAllOfCurrentUser godoc
// @Summary      GetAllOfCurrentUser
// @Description  getting all playlists of user
// @Tags         playlist
// @Accept          application/json
// @Produce      application/json
// @Success      200  {object}  webUtils.Success
// @Failure      400  {object}  webUtils.Error  "Data is invalid"
// @Failure      405  {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/playlists/ofUser [get]
func (h Handler) GetAllOfCurrentUser(ctx echo.Context) error {
	userId, err := internal.GetUserId(ctx, h.UserUseCase)
	if err != nil {
		return internal.UnauthorizedError(ctx)
	}

	playlists, err := h.PlaylistUseCase.GetAllOfCurrentUser(userId)
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	if playlists == nil {
		playlists = []*playlistProto.PlaylistDataTransfer{}
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: toMap(playlists)})
}

// Create godoc
// @Summary      Create
// @Description  creating new playlist
// @Tags         playlist
// @Accept          application/json
// @Produce      application/json
// @Param        Playlist  body      playlistProto.Playlist  true  "params of new playlist. Id will be set automatically."
// @Success      200    {object}  webUtils.Success
// @Failure      400    {object}  webUtils.Error  "Data is invalid"
// @Failure      405    {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/playlists/ [post]
func (h Handler) Create(ctx echo.Context) error {
	userId, err := internal.GetUserId(ctx, h.UserUseCase)
	if err != nil {
		return internal.UnauthorizedError(ctx)
	}

	toCreate := playlistProto.Playlist{}

	if err := ctx.Bind(&toCreate); err != nil {
		return err
	}

	if err := Gateway.Check(&toCreate); err != nil {
		return err
	}

	if err := h.PlaylistUseCase.Create(userId, &toCreate); err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	lastId, err := h.PlaylistUseCase.GetLastIdOfCurrentUser(userId)
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: internal.SuccessCreated + "(" + fmt.Sprint(lastId) + ")"})
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
	userId, err := internal.GetUserId(ctx, h.UserUseCase)
	if err != nil {
		return internal.UnauthorizedError(ctx)
	}

	toUpdate := playlistProto.Playlist{}

	if err := ctx.Bind(&toUpdate); err != nil {
		return err
	}

	if err := Gateway.Check(&toUpdate); err != nil {
		return err
	}

	if err := h.PlaylistUseCase.Update(userId, &toUpdate); err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	id := toUpdate.Id
	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: internal.SuccessUpdated + "(" + fmt.Sprint(id) + ")"})
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
	userId, err := internal.GetUserId(ctx, h.UserUseCase)
	if err != nil {
		userId = -1
	}
	id, err := internal.GetIdInt64ByFieldId(ctx)
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	if id < 0 {
		return webUtils.WriteErrorEchoServer(ctx, errors.New(internal.IndexOutOfRange), http.StatusBadRequest)
	}
	playlist, err := h.PlaylistUseCase.GetById(id, userId)

	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: playlist})
}

// GetOfCurrentUser godoc
// @Summary      GetOfCurrentUser
// @Description  getting playlist by id of current user
// @Tags         playlist
// @Accept          application/json
// @Produce      application/json
// @Param        id   path      integer  true  "id of playlist which need to be getted"
// @Success      200  {object}  webUtils.Success
// @Failure      400  {object}  webUtils.Error  "Data is invalid"
// @Failure      405  {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/playlists/ofUser/{id} [get]
func (h Handler) GetOfCurrentUser(ctx echo.Context) error {
	userId, err := internal.GetUserId(ctx, h.UserUseCase)
	if err != nil {
		return internal.UnauthorizedError(ctx)
	}

	id, err := internal.GetIdInt64ByFieldId(ctx)
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	if id < 0 {
		return webUtils.WriteErrorEchoServer(ctx, errors.New(internal.IndexOutOfRange), http.StatusBadRequest)
	}
	playlist, err := h.PlaylistUseCase.GetByIdOfCurrentUser(userId, id)

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
// @Tags         playlist
// @Accept       application/json
// @Produce      application/json
// @Param        id   path      integer  true  "id of playlists which need to be deleted"
// @Success      200  {object}  webUtils.Success
// @Failure      400  {object}  webUtils.Error  "Data is invalid"
// @Failure      405  {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/playlists/{id} [delete]
func (h Handler) Delete(ctx echo.Context) error {
	userId, err := internal.GetUserId(ctx, h.UserUseCase)
	if err != nil {
		return internal.UnauthorizedError(ctx)
	}

	id, err := internal.GetIdInt64ByFieldId(ctx)
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	if id < 0 {
		return webUtils.WriteErrorEchoServer(ctx, errors.New(internal.IndexOutOfRange), http.StatusBadRequest)
	}

	if err := h.PlaylistUseCase.Delete(userId, id); err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: internal.SuccessDeleted + "(" + fmt.Sprint(id) + ")"})
}

// AddToPlaylist godoc
// @Summary      AddToPlaylist
// @Description  adding track_id to playlist
// @Tags         playlist
// @Accept       application/json
// @Produce      application/json
// @Param        playlistIdTrackId  body      playlistProto.PlaylistIdTrackId  true  "ids of playlist and track"
// @Success      200  {object}  webUtils.Success
// @Failure      400  {object}  webUtils.Error  "Data is invalid"
// @Failure      405  {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/playlists/ofUser [post]
func (h Handler) AddToPlaylist(ctx echo.Context) error {
	userId, err := internal.GetUserId(ctx, h.UserUseCase)
	if err != nil {
		return internal.UnauthorizedError(ctx)
	}

	idHolder := playlistProto.PlaylistIdTrackId{}

	if err := ctx.Bind(&idHolder); err != nil {
		return err
	}

	playlistId := idHolder.PlaylistId
	trackId := idHolder.TrackId

	if err := h.PlaylistUseCase.AddToPlaylist(userId, playlistId, trackId); err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: internal.SuccessAdded + "(" + fmt.Sprint(trackId) + ")" + internal.ToPlaylist + "(" + fmt.Sprint(playlistId) + ")"})
}

// RemoveFromPlaylist godoc
// @Summary      RemoveFromPlaylist
// @Description  remove track_id from playlist
// @Tags         playlist
// @Accept       application/json
// @Produce      application/json
// @Param        playlistId   query      integer  true  "id of playlists"
// @Param        trackId   query      integer  true  "id of track"
// @Success      200  {object}  webUtils.Success
// @Failure      400  {object}  webUtils.Error  "Data is invalid"
// @Failure      405  {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/playlists/ofUser [delete]
func (h Handler) RemoveFromPlaylist(ctx echo.Context) error {
	userId, err := internal.GetUserId(ctx, h.UserUseCase)
	if err != nil {
		return internal.UnauthorizedError(ctx)
	}

	playlistId, err := strconv.ParseInt(ctx.QueryParam(internal.PlaylistId), 10, 64)
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	trackId, err := strconv.ParseInt(ctx.QueryParam(internal.TrackId), 10, 64)
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	if err := h.PlaylistUseCase.RemoveFromPlaylist(userId, playlistId, trackId); err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: internal.SuccessRemoved + "(" + fmt.Sprint(trackId) + ")" + internal.FromPlaylist + "(" + fmt.Sprint(playlistId) + ")"})
}
