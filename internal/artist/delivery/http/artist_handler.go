package artistDeliveryHttp

import (
	"errors"
	"fmt"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	ArtistUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/artist/useCase"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/artist/artistProto"
	Gateway "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/gateway"
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
	domains, err := h.ArtistUseCase.GetAll()

	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	if domains == nil {
		domains = []*artistProto.ArtistDataTransfer{}
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: domains})
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
			Result: constants.SuccessCreated + "(" + fmt.Sprint(lastId) + ")"})
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
			Result: constants.SuccessUpdated + "(" + fmt.Sprint(id) + ")"})
}

// Get godoc
// @Summary      Get
// @Description  getting artist by id
// @Tags         artist
// @Accept          application/json
// @Produce      application/json
// @Param        id   path      integer  true  "id of artist which need to be getted"
// @Success      200  {object}  artistProto.Artist
// @Failure      400  {object}  webUtils.Error  "Data is invalid"
// @Failure      405  {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/artists/{id} [get]
func (h Handler) Get(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param(constants.FieldId), 10, 64)
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	if id < 0 {
		return webUtils.WriteErrorEchoServer(ctx, errors.New(constants.IndexOutOfRange), http.StatusBadRequest)
	}

	obj, err := h.ArtistUseCase.GetById(id)

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
	id, err := strconv.ParseInt(ctx.Param(constants.FieldId), 10, 64)
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}
	if id < 0 {
		return webUtils.WriteErrorEchoServer(ctx, errors.New(constants.IndexOutOfRange), http.StatusBadRequest)
	}

	if err := h.ArtistUseCase.Delete(id); err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: constants.SuccessDeleted + "(" + fmt.Sprint(id) + ")"})
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
	popular, err := h.ArtistUseCase.GetPopular()
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
// @Tags         artist
// @Accept          application/json
// @Produce      application/json
// @Param        id   path      integer  true  "id of artist which need to be getted"
// @Success      200  {object}  webUtils.Success
// @Failure      400  {object}  webUtils.Error  "Data is invalid"
// @Failure      405  {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/artists/{id}/popular [get]
func (h Handler) GetPopularTracks(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param(constants.FieldId), 10, 64)
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}
	if id < 0 {
		return webUtils.WriteErrorEchoServer(ctx, errors.New(constants.IndexOutOfRange), http.StatusBadRequest)
	}

	popular, err := h.TrackUseCase.GetPopularTracksFromArtist(id)
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: popular})
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
	cookie, err := ctx.Cookie(constants.SessionIdKey)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, err)
	}

	user, err := h.UserUseCase.GetBySessionId(cookie.Value)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, err)
	}

	userId := int64(user.ID)
	favorites, err := h.ArtistUseCase.GetFavorites(userId)
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
// @Description  add to favorite
// @Tags         artist
// @Accept          application/json
// @Produce      application/json
// @Param        artistId  path      int  true  "artistId"
// @Success      200    {object}  webUtils.Success
// @Failure      400    {object}  webUtils.Error  "Data is invalid"
// @Failure      405    {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/artists/favorites/{id} [post]
func (h Handler) AddToFavorites(ctx echo.Context) error {
	cookie, err := ctx.Cookie(constants.SessionIdKey)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, err)
	}

	user, err := h.UserUseCase.GetBySessionId(cookie.Value)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, err)
	}

	userId := int64(user.ID)

	artistId, err := strconv.ParseInt(ctx.Param(constants.FieldId), 10, 64)
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	if err := h.ArtistUseCase.AddToFavorites(userId, artistId); err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: constants.SuccessAddedToFavorites + "(" + fmt.Sprint(artistId) + ")"})
}

// RemoveFromFavorites godoc
// @Summary      RemoveFromFavorites
// @Description  remove from favorites
// @Tags         artist
// @Accept          application/json
// @Produce      application/json
// @Param        artistId  path      int  true  "artistId"
// @Success      200    {object}  webUtils.Success
// @Failure      400    {object}  webUtils.Error  "Data is invalid"
// @Failure      405    {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/artists/favorites/{id} [delete]
func (h Handler) RemoveFromFavorites(ctx echo.Context) error {
	cookie, err := ctx.Cookie(constants.SessionIdKey)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, err)
	}

	user, err := h.UserUseCase.GetBySessionId(cookie.Value)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, err)
	}

	userId := int64(user.ID)

	artistId, err := strconv.ParseInt(ctx.Param(constants.FieldId), 10, 64)

	if err := h.ArtistUseCase.RemoveFromFavorites(userId, artistId); err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: constants.SuccessRemoveFromFavorites + "(" + fmt.Sprint(artistId) + ")"})
}
