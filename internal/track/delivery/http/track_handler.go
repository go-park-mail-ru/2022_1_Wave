package trackDeliveryHttp

import (
	"errors"
	"fmt"
	internal "github.com/go-park-mail-ru/2022_1_Wave/internal"
	Gateway "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/gateway"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/track/trackProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/structs"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/tools/utils"
	TrackUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/track/useCase"
	user_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/user"
	"github.com/go-park-mail-ru/2022_1_Wave/pkg/webUtils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	UserUseCase  user_domain.UserUseCase
	TrackUseCase TrackUseCase.TrackUseCase
}

func MakeHandler(track TrackUseCase.TrackUseCase, user user_domain.UserUseCase) Handler {
	return Handler{
		UserUseCase:  user,
		TrackUseCase: track,
	}
}

// GetAll godoc
// @Summary      GetAll
// @Description  getting all tracks
// @Tags         track
// @Accept          application/json
// @Produce      application/json
// @Success      200  {object}  webUtils.Success
// @Failure      400  {object}  webUtils.Error  "Data is invalid"
// @Failure      405  {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/tracks/ [get]
func (h Handler) GetAll(ctx echo.Context) error {
	userId, err := internal.GetUserId(ctx, h.UserUseCase)
	if err != nil {
		userId = -1
	}

	tracks, err := h.TrackUseCase.GetAll(userId)
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	if tracks == nil {
		tracks = []*trackProto.TrackDataTransfer{}
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: tracks})
}

// Create godoc
// @Summary      Create
// @Description  creating new track
// @Tags         track
// @Accept          application/json
// @Produce      application/json
// @Param        Track  body      trackProto.Track  true  "params of new track. Id will be set automatically."
// @Success      200    {object}  webUtils.Success
// @Failure      400    {object}  webUtils.Error  "Data is invalid"
// @Failure      405    {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/tracks/ [post]
func (h Handler) Create(ctx echo.Context) error {
	result := trackProto.Track{}

	if err := ctx.Bind(&result); err != nil {
		return err
	}

	if err := Gateway.Check(&result); err != nil {
		return err
	}

	if err := h.TrackUseCase.Create(&result); err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	lastId, err := h.TrackUseCase.GetLastId()
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
// @Description  updating track by id
// @Tags         track
// @Accept          application/json
// @Produce      application/json
// @Param        Track  body      trackProto.Track  true  "id of updating track and params of it."
// @Success      200    {object}  webUtils.Success
// @Failure      400    {object}  webUtils.Error  "Data is invalid"
// @Failure      405    {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/tracks/ [put]
func (h Handler) Update(ctx echo.Context) error {
	result := trackProto.Track{}

	if err := ctx.Bind(&result); err != nil {
		return err
	}

	if err := Gateway.Check(&result); err != nil {
		return err
	}

	if err := h.TrackUseCase.Update(&result); err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	id := result.Id
	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: internal.SuccessUpdated + "(" + fmt.Sprint(id) + ")"})
}

// Get godoc
// @Summary      Get
// @Description  getting track by id
// @Tags         track
// @Accept          application/json
// @Produce      application/json
// @Param        id   path      integer  true  "id of track which need to be getted"
// @Success      200  {object}  webUtils.Success
// @Failure      400  {object}  webUtils.Error  "Data is invalid"
// @Failure      405  {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/tracks/{id} [get]
func (h Handler) Get(ctx echo.Context) error {
	userId, err := internal.GetUserId(ctx, h.UserUseCase)
	if err != nil {
		userId = -1
	}

	trackId, err := internal.GetIdInt64ByFieldId(ctx)
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}
	if trackId < 0 {
		return webUtils.WriteErrorEchoServer(ctx, errors.New(internal.IndexOutOfRange), http.StatusBadRequest)
	}
	track, err := h.TrackUseCase.GetById(trackId, userId)

	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: track})
}

// Delete godoc
// @Summary      Delete
// @Description  deleting track by id
// @Tags         track
// @Accept       application/json
// @Produce      application/json
// @Param        id   path      integer  true  "id of track which need to be deleted"
// @Success      200  {object}  webUtils.Success
// @Failure      400  {object}  webUtils.Error  "Data is invalid"
// @Failure      405  {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/tracks/{id} [delete]
func (h Handler) Delete(ctx echo.Context) error {
	id, err := internal.GetIdInt64ByFieldId(ctx)
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}
	if id < 0 {
		return webUtils.WriteErrorEchoServer(ctx, errors.New(internal.IndexOutOfRange), http.StatusBadRequest)
	}

	if err := h.TrackUseCase.Delete(id); err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: internal.SuccessDeleted + "(" + fmt.Sprint(id) + ")"})
}

// GetPopular godoc
// @Summary      GetPopular
// @Description  getting top20 popular tracks
// @Tags         track
// @Accept          application/json
// @Produce      application/json
// @Success      200  {object}  webUtils.Success
// @Failure      400  {object}  webUtils.Error  "Data is invalid"
// @Failure      405  {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/tracks/popular [get]
func (h Handler) GetPopular(ctx echo.Context) error {
	userId, err := internal.GetUserId(ctx, h.UserUseCase)
	if err != nil {
		userId = -1
	}

	popular, err := h.TrackUseCase.GetPopular(userId)
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: popular})
}

// GetPopularTracks godoc
// @Summary      GetPopularTracks
// @Description  getting top20 popular tracks of this artist
// @Tags         track
// @Accept          application/json
// @Produce      application/json
// @Param        id   path      integer  true  "id of artist which need to be getted"
// @Success      200  {object}  webUtils.Success
// @Failure      400  {object}  webUtils.Error  "Data is invalid"
// @Failure      405  {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/artists/{id}/popular [get]
func (h Handler) GetPopularTracks(ctx echo.Context) error {
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

	popular, err := h.TrackUseCase.GetPopularTracksFromArtist(id, userId)
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: popular})
}

// Like godoc
// @Summary      Like
// @Description  like track by id
// @Tags         track
// @Accept          application/json
// @Produce      application/json
// @Param        id   path      integer  true  "id of track which need to be liked"
// @Success      200  {object}  webUtils.Success
// @Failure      400  {object}  webUtils.Error  "Data is invalid"
// @Failure      405  {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/tracks/like/{id} [put]
func (h Handler) Like(ctx echo.Context) error {
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

	if err := h.TrackUseCase.Like(id, userId); err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: internal.SuccessLiked + "(" + fmt.Sprint(id) + ")"})
}

// LikeCheckByUser godoc
// @Summary      LikeCheckByUser
// @Description  LikeCheckByUser
// @Tags         track
// @Accept          application/json
// @Produce      application/json
// @Param        id   path      integer  true  "id of track which need to check for like"
// @Success      200  {object}  webUtils.Success
// @Failure      400  {object}  webUtils.Error  "Data is invalid"
// @Failure      405  {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/tracks/like/{id} [get]
func (h Handler) LikeCheckByUser(ctx echo.Context) error {
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

	liked, err := h.TrackUseCase.LikeCheckByUser(id, userId)
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: liked})
}

// Listen godoc
// @Summary      Listen
// @Description  listen track by id
// @Tags         track
// @Accept          application/json
// @Produce      application/json
// @Param        id   path      integer  true  "id of track which need to be listen"
// @Success      200  {object}  webUtils.Success
// @Failure      400  {object}  webUtils.Error  "Data is invalid"
// @Failure      405  {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/tracks/listen/{id} [put]
func (h Handler) Listen(ctx echo.Context) error {
	id, err := internal.GetIdInt64ByFieldId(ctx)
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}
	if id < 0 {
		return webUtils.WriteErrorEchoServer(ctx, errors.New(internal.IndexOutOfRange), http.StatusBadRequest)
	}

	if err := h.TrackUseCase.Listen(id); err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: internal.SuccessListened + "(" + fmt.Sprint(id) + ")"})
}

// GetFavorites godoc
// @Summary      GetFavorites
// @Description  getting favorites tracks
// @Tags         track
// @Accept          application/json
// @Produce      application/json
// @Success      200  {object}  webUtils.Success
// @Failure      400  {object}  webUtils.Error  "Data is invalid"
// @Failure      405  {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/tracks/favorites [get]
func (h Handler) GetFavorites(ctx echo.Context) error {
	userId, err := internal.GetUserId(ctx, h.UserUseCase)
	if err != nil {
		return internal.UnauthorizedError(ctx)
	}
	favorites, err := h.TrackUseCase.GetFavorites(userId)
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: favorites})
}

// AddToFavorites godoc
// @Summary      AddToFavorites
// @Description  add to favorites
// @Tags         track
// @Accept          application/json
// @Produce      application/json
// @Param        trackId  body      structs.TrackIdWrapper  true  "id of track"
// @Success      200    {object}  webUtils.Success
// @Failure      400    {object}  webUtils.Error  "Data is invalid"
// @Failure      405    {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/tracks/favorites [post]
func (h Handler) AddToFavorites(ctx echo.Context) error {
	userId, err := internal.GetUserId(ctx, h.UserUseCase)
	if err != nil {
		return internal.UnauthorizedError(ctx)
	}

	holder := structs.TrackIdWrapper{}

	if err := ctx.Bind(&holder); err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	if err := h.TrackUseCase.AddToFavorites(userId, int64(holder.TrackId)); err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: internal.SuccessAddedToFavorites + "(" + fmt.Sprint(holder) + ")"})
}

// RemoveFromFavorites godoc
// @Summary      RemoveFromFavorites
// @Description  remove from favorite
// @Tags         track
// @Accept          application/json
// @Produce      application/json
// @Param        id  path      integer  true  "trackId"
// @Success      200    {object}  webUtils.Success
// @Failure      400    {object}  webUtils.Error  "Data is invalid"
// @Failure      405    {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/tracks/favorites/{id} [delete]
func (h Handler) RemoveFromFavorites(ctx echo.Context) error {
	userId, err := internal.GetUserId(ctx, h.UserUseCase)
	if err != nil {
		return internal.UnauthorizedError(ctx)
	}

	trackId, err := internal.GetIdInt64ByFieldId(ctx)
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	if err := h.TrackUseCase.RemoveFromFavorites(userId, trackId); err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: internal.SuccessRemoveFromFavorites + "(" + fmt.Sprint(trackId) + ")"})
}

// GetTracksFromPlaylist godoc
// @Summary      GetTracksFromPlaylist
// @Description  get tracks from playlist by id
// @Tags         track
// @Accept          application/json
// @Produce      application/json
// @Param        id  path      int  true  "playlistId"
// @Success      200    {object}  webUtils.Success
// @Failure      400    {object}  webUtils.Error  "Data is invalid"
// @Failure      405    {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/tracks/playlist/{id} [get]
func (h Handler) GetTracksFromPlaylist(ctx echo.Context) error {
	userId, err := internal.GetUserId(ctx, h.UserUseCase)
	if err != nil {
		userId = -1
	}

	playlistId, err := internal.GetIdInt64ByFieldId(ctx)
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}
	tracks, err := h.TrackUseCase.GetTracksFromPlaylist(playlistId, userId)
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: utils.TracksToMap(tracks)})
}
