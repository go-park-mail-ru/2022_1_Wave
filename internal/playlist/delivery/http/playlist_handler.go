package trackDeliveryHttp

import (
	"errors"
	"fmt"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	Gateway "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/gateway"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/gateway/gatewayProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/track/trackProto"
	PlaylistUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/playlist/useCase"
	"github.com/go-park-mail-ru/2022_1_Wave/pkg/webUtils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Handler struct {
	PlaylistUseCase PlaylistUseCase.PlaylistAgent
}

func MakeHandler(playlist PlaylistUseCase.PlaylistAgent, user user_do) Handler {
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

	user, err := UserUseCase.GetBySessionId(cookie.Value)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, getErrorUserResponse(err))
	}

	domains, err := h.PlaylistUseCase.GetAll()

	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	if domains == nil {
		domains = &trackProto.TracksResponse{}
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: domains})
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

	if err := h.PlaylistUseCase.Create(&result); err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	lastId, err := h.PlaylistUseCase.GetLastId()
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

	if err := h.PlaylistUseCase.Update(&result); err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	id := result.Id
	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: constants.SuccessUpdated + "(" + fmt.Sprint(id) + ")"})
}

// Get godoc
// @Summary      Get
// @Description  getting track by id
// @Tags         track
// @Accept          application/json
// @Produce      application/json
// @Param        id   path      integer  true  "id of track which need to be getted"
// @Success      200  {object}  trackProto.Track
// @Failure      400  {object}  webUtils.Error  "Data is invalid"
// @Failure      405  {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/tracks/{id} [get]
func (h Handler) Get(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param(constants.FieldId))
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}
	if id < 0 {
		return webUtils.WriteErrorEchoServer(ctx, errors.New(constants.IndexOutOfRange), http.StatusBadRequest)
	}
	track, err := h.PlaylistUseCase.GetById(&gatewayProto.IdArg{Id: int64(id)})

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
	id, err := strconv.Atoi(ctx.Param(constants.FieldId))
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}
	if id < 0 {
		return webUtils.WriteErrorEchoServer(ctx, errors.New(constants.IndexOutOfRange), http.StatusBadRequest)
	}

	if err := h.PlaylistUseCase.Delete(&gatewayProto.IdArg{Id: int64(id)}); err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: constants.SuccessDeleted + "(" + fmt.Sprint(id) + ")"})
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
	popular, err := h.PlaylistUseCase.GetPopular()
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
	id, err := strconv.Atoi(ctx.Param(constants.FieldId))
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}
	if id < 0 {
		return webUtils.WriteErrorEchoServer(ctx, errors.New(constants.IndexOutOfRange), http.StatusBadRequest)
	}

	if err := h.PlaylistUseCase.Like(&gatewayProto.IdArg{Id: int64(id)}); err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: constants.SuccessLiked + "(" + fmt.Sprint(id) + ")"})
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
	id, err := strconv.Atoi(ctx.Param(constants.FieldId))
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}
	if id < 0 {
		return webUtils.WriteErrorEchoServer(ctx, errors.New(constants.IndexOutOfRange), http.StatusBadRequest)
	}

	if err := h.PlaylistUseCase.Listen(&gatewayProto.IdArg{Id: int64(id)}); err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: constants.SuccessListened + "(" + fmt.Sprint(id) + ")"})
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
	//todo userId is not 0!!!
	userId := int64(0)
	favorites, err := h.PlaylistUseCase.GetFavorites(&gatewayProto.IdArg{Id: userId})
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
// @Param        trackId  path      int  true  "trackId"
// @Success      200    {object}  webUtils.Success
// @Failure      400    {object}  webUtils.Error  "Data is invalid"
// @Failure      405    {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/tracks/favorites/{id} [post]
func (h Handler) AddToFavorites(ctx echo.Context) error {
	trackId, err := strconv.Atoi(ctx.Param(constants.FieldId))
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}
	//todo userId is not 0!!!
	userId := int64(0)

	if _, err := h.PlaylistUseCase.AddToFavorites(&gatewayProto.UserIdTrackIdArg{
		UserId:  userId,
		TrackId: int64(trackId),
	}); err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: constants.SuccessAddedToFavorites + "(" + fmt.Sprint(trackId) + ")"})
}

// RemoveFromFavorites godoc
// @Summary      RemoveFromFavorites
// @Description  remove from favorite
// @Tags         track
// @Accept          application/json
// @Produce      application/json
// @Param        trackId  path      int  true  "trackId"
// @Success      200    {object}  webUtils.Success
// @Failure      400    {object}  webUtils.Error  "Data is invalid"
// @Failure      405    {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/tracks/favorites/{id} [delete]
func (h Handler) RemoveFromFavorites(ctx echo.Context) error {
	trackId, err := strconv.Atoi(ctx.Param(constants.FieldId))
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}
	//todo userId is not 0!!!
	userId := int64(0)

	if _, err := h.PlaylistUseCase.RemoveFromFavorites(&gatewayProto.UserIdTrackIdArg{
		UserId:  userId,
		TrackId: int64(trackId),
	}); err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: constants.SuccessRemoveFromFavorites + "(" + fmt.Sprint(trackId) + ")"})
}
