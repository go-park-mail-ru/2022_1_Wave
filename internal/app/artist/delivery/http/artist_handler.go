package ArtistDeliveryHttp

import (
	"errors"
	"fmt"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	AlbumUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/album/usecase"
	ArtistUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/artist/usecase"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	TrackUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/track/usecase"
	"github.com/go-park-mail-ru/2022_1_Wave/pkg/webUtils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Handler struct {
	ArtistUseCase ArtistUseCase.ArtistUseCaseInterface
	AlbumUseCase  AlbumUseCase.AlbumUseCaseInterface
	TrackUseCase  TrackUseCase.TrackUseCaseInterface
}

func MakeHandler(artist ArtistUseCase.ArtistUseCase, album AlbumUseCase.AlbumUseCase, track TrackUseCase.TrackUseCase) Handler {
	return Handler{
		ArtistUseCase: artist,
		AlbumUseCase:  album,
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
	domains, err := h.ArtistUseCase.GetAll(h.AlbumUseCase, h.TrackUseCase)

	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	if domains == nil {
		domains = []domain.ArtistDataTransfer{}
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
// @Param        Artist  body      domain.Artist  true  "params of new artist. Id will be set automatically."
// @Success      200     {object}  webUtils.Success
// @Failure      400     {object}  webUtils.Error  "Data is invalid"
// @Failure      405     {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/artists/ [post]
func (h Handler) Create(ctx echo.Context) error {
	result := domain.Artist{}

	if err := ctx.Bind(&result); err != nil {
		return err
	}

	if err := result.Check(); err != nil {
		return err
	}

	if err := h.ArtistUseCase.Create(result); err != nil {
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
// @Param        Artist  body      domain.Artist  true  "id of updating artist and params of it."
// @Success      200     {object}  webUtils.Success
// @Failure      400     {object}  webUtils.Error  "Data is invalid"
// @Failure      405     {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/artists/ [put]
func (h Handler) Update(ctx echo.Context) error {
	result := domain.Artist{}

	if err := ctx.Bind(&result); err != nil {
		return err
	}

	if err := result.Check(); err != nil {
		return err
	}

	if err := h.ArtistUseCase.Update(result); err != nil {
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
// @Success      200  {object}  domain.Artist
// @Failure      400  {object}  webUtils.Error  "Data is invalid"
// @Failure      405  {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/artists/{id} [get]
func (h Handler) Get(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param(constants.FieldId))
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}
	if id < 0 {
		return webUtils.WriteErrorEchoServer(ctx, errors.New(constants.IndexOutOfRange), http.StatusBadRequest)
	}

	obj, err := h.ArtistUseCase.GetById(h.TrackUseCase, h.AlbumUseCase, id)

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
	id, err := strconv.Atoi(ctx.Param(constants.FieldId))
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
	popular, err := h.ArtistUseCase.GetPopular(h.TrackUseCase, h.AlbumUseCase)
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
	id, err := strconv.Atoi(ctx.Param(constants.FieldId))
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
