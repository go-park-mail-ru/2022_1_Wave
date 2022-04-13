package trackDeliveryHttp

import (
	"errors"
	"fmt"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	ArtistUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/artist/usecase"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	TrackUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/track/usecase"
	"github.com/go-park-mail-ru/2022_1_Wave/pkg/webUtils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Handler struct {
	artistUseCase ArtistUseCase.ArtistUseCase
	trackUseCase  TrackUseCase.TrackUseCase
}

func MakeHandler(artist ArtistUseCase.ArtistUseCase, track TrackUseCase.TrackUseCase) Handler {
	return Handler{
		artistUseCase: artist,
		trackUseCase:  track,
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
	domains, err := h.trackUseCase.GetAll()

	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	if domains == nil {
		domains = []domain.TrackDataTransfer{}
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
// @Param        Track  body      domain.Track  true  "params of new track. Id will be set automatically."
// @Success      200    {object}  webUtils.Success
// @Failure      400    {object}  webUtils.Error  "Data is invalid"
// @Failure      405    {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/tracks/ [post]
func (h Handler) Create(ctx echo.Context) error {
	result := domain.Track{}

	if err := ctx.Bind(&result); err != nil {
		return err
	}

	if err := result.Check(); err != nil {
		return err
	}

	if err := h.trackUseCase.Create(result); err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	lastId, err := h.trackUseCase.GetLastId()
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
// @Param        Track  body      domain.Track  true  "id of updating track and params of it."
// @Success      200    {object}  webUtils.Success
// @Failure      400    {object}  webUtils.Error  "Data is invalid"
// @Failure      405    {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/tracks/ [put]
func (h Handler) Update(ctx echo.Context) error {
	result := domain.Track{}

	if err := ctx.Bind(&result); err != nil {
		return err
	}

	if err := result.Check(); err != nil {
		return err
	}

	if err := h.trackUseCase.Update(result); err != nil {
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
// @Success      200  {object}  domain.Track
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

	track, err := h.trackUseCase.GetById(id)

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

	if err := h.trackUseCase.Delete(id); err != nil {
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
	popular, err := h.trackUseCase.GetPopular()
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: popular})
}
