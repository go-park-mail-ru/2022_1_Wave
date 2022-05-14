package artistDeliveryHttp

import (
	"errors"
	"fmt"
	internal "github.com/go-park-mail-ru/2022_1_Wave/internal"
	ArtistUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/artist/useCase"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/artist/artistProto"
	Gateway "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/gateway"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/tools/utils"
	TrackUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/track/useCase"
	user_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/user"
	"github.com/go-park-mail-ru/2022_1_Wave/pkg/webUtils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Handler struct {
	UserUseCase   user_domain.UserUseCase
	ArtistUseCase ArtistUseCase.UseCase
	TrackUseCase  TrackUseCase.UseCase
}

func MakeHandler(artist ArtistUseCase.UseCase, track TrackUseCase.UseCase, user user_domain.UserUseCase) Handler {
	return Handler{
		UserUseCase:   user,
		ArtistUseCase: artist,
		TrackUseCase:  track,
	}
}

// GetAll godoc
// @Summary      GetAll
// @Description  getting all artists
// @Tags         artist
// @Accept          application/json
// @Produce      application/json
// @Success      200  {object}  webUtils.Success
// @Failure      400  {object}  webUtils.Error  "Data is invalid"
// @Failure      405  {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/artists/ [get]
func (h Handler) GetAll(ctx echo.Context) error {
	userId, err := internal.GetUserId(ctx, h.UserUseCase)
	artists, err := h.ArtistUseCase.GetAll(userId)

	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	if artists == nil {
		artists = []*artistProto.ArtistDataTransfer{}
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: utils.ArtistsToMap(artists)})
}

// Create godoc
// @Summary      Create
// @Description  creating new artist
// @Tags         artist
// @Accept          application/json
// @Produce      application/json
// @Param        Artist  body      artistProto.Artist  true  "params of new artist. Id will be set automatically."
// @Success      200     {object}  webUtils.Success
// @Failure      400     {object}  webUtils.Error  "Data is invalid"
// @Failure      405     {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/artists/ [post]
func (h Handler) Create(ctx echo.Context) error {
	result := artistProto.Artist{}

	if err := ctx.Bind(&result); err != nil {
		return err
	}

	if err := Gateway.Check(&result); err != nil {
		return err
	}

	if err := h.ArtistUseCase.Create(&result); err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	lastId, err := h.ArtistUseCase.GetLastId()
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
// @Description  updating artist by id
// @Tags         artist
// @Accept          application/json
// @Produce      application/json
// @Param        Artist  body      artistProto.Artist  true  "id of updating artist and params of it."
// @Success      200     {object}  webUtils.Success
// @Failure      400     {object}  webUtils.Error  "Data is invalid"
// @Failure      405     {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/artists/ [put]
func (h Handler) Update(ctx echo.Context) error {
	result := artistProto.Artist{}

	if err := ctx.Bind(&result); err != nil {
		return err
	}

	if err := Gateway.Check(&result); err != nil {
		return err
	}

	if err := h.ArtistUseCase.Update(&result); err != nil {
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
// @Description  getting artist by id
// @Tags         artist
// @Accept          application/json
// @Produce      application/json
// @Param        id   path      integer  true  "id of artist which need to be getted"
// @Success      200  {object}  webUtils.Success
// @Failure      400  {object}  webUtils.Error  "Data is invalid"
// @Failure      405  {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/artists/{id} [get]
func (h Handler) Get(ctx echo.Context) error {
	userId, err := internal.GetUserId(ctx, h.UserUseCase)
	id, err := internal.GetIdInt64ByFieldId(ctx)
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	if id < 0 {
		return webUtils.WriteErrorEchoServer(ctx, errors.New(internal.IndexOutOfRange), http.StatusBadRequest)
	}

	obj, err := h.ArtistUseCase.GetById(id, userId)

	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: obj})
}

// Delete godoc
// @Summary      Delete
// @Description  deleting artist by id
// @Tags         artist
// @Accept       application/json
// @Produce      application/json
// @Param        id   path      integer  true  "id of artist which need to be deleted"
// @Success      200  {object}  webUtils.Success
// @Failure      400  {object}  webUtils.Error  "Data is invalid"
// @Failure      405  {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/artists/{id} [delete]
func (h Handler) Delete(ctx echo.Context) error {
	id, err := internal.GetIdInt64ByFieldId(ctx)
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}
	if id < 0 {
		return webUtils.WriteErrorEchoServer(ctx, errors.New(internal.IndexOutOfRange), http.StatusBadRequest)
	}

	if err := h.ArtistUseCase.Delete(id); err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: internal.SuccessDeleted + "(" + fmt.Sprint(id) + ")"})
}

// GetPopular godoc
// @Summary      GetPopular
// @Description  getting top20 popular artists
// @Tags         artist
// @Accept          application/json
// @Produce      application/json
// @Success      200  {object}  webUtils.Success
// @Failure      400  {object}  webUtils.Error  "Data is invalid"
// @Failure      405  {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/artists/popular [get]
func (h Handler) GetPopular(ctx echo.Context) error {
	userId, err := internal.GetUserId(ctx, h.UserUseCase)

	popular, err := h.ArtistUseCase.GetPopular(userId)
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: utils.ArtistsToMap(popular)})
}

// GetFavorites godoc
// @Summary      GetFavorites
// @Description  getting favorites artist
// @Tags         artist
// @Accept          application/json
// @Produce      application/json
// @Success      200  {object}  webUtils.Success
// @Failure      400  {object}  webUtils.Error  "Data is invalid"
// @Failure      405  {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/artists/favorites [get]
func (h Handler) GetFavorites(ctx echo.Context) error {
	userId, err := internal.GetUserId(ctx, h.UserUseCase)
	if err != nil {
		return internal.UnauthorizedError(ctx)
	}

	favorites, err := h.ArtistUseCase.GetFavorites(userId)
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: utils.ArtistsToMap(favorites)})
}

type artistIdWrapper struct {
	ArtistId int64 `json:"artistId" example:"4"`
}

// AddToFavorites godoc
// @Summary      AddToFavorites
// @Description  add to favorite
// @Tags         artist
// @Accept          application/json
// @Produce      application/json
// @Param        artistId  body      artistIdWrapper  true  "artistId"
// @Success      200    {object}  webUtils.Success
// @Failure      400    {object}  webUtils.Error  "Data is invalid"
// @Failure      405    {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/artists/favorites [post]
func (h Handler) AddToFavorites(ctx echo.Context) error {
	userId, err := internal.GetUserId(ctx, h.UserUseCase)
	if err != nil {
		return internal.UnauthorizedError(ctx)
	}

	holder := artistIdWrapper{}

	if err := ctx.Bind(&holder); err != nil {
		return err
	}

	if err := h.ArtistUseCase.AddToFavorites(userId, holder.ArtistId); err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: internal.SuccessAddedToFavorites + "(" + fmt.Sprint(holder.ArtistId) + ")"})
}

// RemoveFromFavorites godoc
// @Summary      RemoveFromFavorites
// @Description  remove from favorites
// @Tags         artist
// @Accept          application/json
// @Produce      application/json
// @Param        id  path      integer  true  "artistId"
// @Success      200    {object}  webUtils.Success
// @Failure      400    {object}  webUtils.Error  "Data is invalid"
// @Failure      405    {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/artists/favorites/{id} [delete]
func (h Handler) RemoveFromFavorites(ctx echo.Context) error {
	userId, err := internal.GetUserId(ctx, h.UserUseCase)
	if err != nil {
		return internal.UnauthorizedError(ctx)
	}

	artistId, err := strconv.ParseInt(ctx.Param(internal.FieldId), 10, 64)

	if err := h.ArtistUseCase.RemoveFromFavorites(userId, artistId); err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: internal.SuccessRemoveFromFavorites + "(" + fmt.Sprint(artistId) + ")"})
}

// Like godoc
// @Summary      Like
// @Description  like artist by id
// @Tags         artist
// @Accept          application/json
// @Produce      application/json
// @Param        id   path      integer  true  "id of artist which need to be liked"
// @Success      200  {object}  webUtils.Success
// @Failure      400  {object}  webUtils.Error  "Data is invalid"
// @Failure      405  {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/artists/like/{id} [put]
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

	if err := h.ArtistUseCase.Like(id, userId); err != nil {
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
// @Tags         artist
// @Accept          application/json
// @Produce      application/json
// @Param        id   path      integer  true  "id of artist which need to check for like"
// @Success      200  {object}  webUtils.Success
// @Failure      400  {object}  webUtils.Error  "Data is invalid"
// @Failure      405  {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/artists/like/{id} [get]
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

	liked, err := h.ArtistUseCase.LikeCheckByUser(id, userId)
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: liked})
}
