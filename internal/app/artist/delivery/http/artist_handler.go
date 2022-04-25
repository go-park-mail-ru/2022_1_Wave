package ArtistDeliveryHttp

import (
	"context"
	"errors"
	"fmt"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	AlbumRepo "github.com/go-park-mail-ru/2022_1_Wave/internal/app/album/repository"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/microservices/album/albumProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/microservices/artist/artistProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/microservices/common/commonProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/microservices/track/trackProto"
	"github.com/go-park-mail-ru/2022_1_Wave/pkg/webUtils"
	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
	"strconv"
)

type Handler struct {
	ArtistUseCase artistProto.ArtistUseCaseClient
	AlbumUseCase  albumProto.AlbumUseCaseClient
	TrackUseCase  trackProto.TrackUseCaseClient
}

func MakeHandler(artist artistProto.ArtistUseCaseClient, album albumProto.AlbumUseCaseClient, track trackProto.TrackUseCaseClient) Handler {
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
	domains, err := h.ArtistUseCase.GetAll(context.Background(), &emptypb.Empty{})

	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	if domains == nil {
		domains = &artistProto.ArtistsResponse{}
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

	if err := AlbumRepo.Check(&result); err != nil {
		return err
	}

	if _, err := h.ArtistUseCase.Create(context.Background(), &result); err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	lastId, err := h.ArtistUseCase.GetLastId(context.Background(), &emptypb.Empty{})
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

	if err := AlbumRepo.Check(&result); err != nil {
		return err
	}

	if _, err := h.ArtistUseCase.Update(context.Background(), &result); err != nil {
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
	id, err := strconv.Atoi(ctx.Param(constants.FieldId))
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}
	if id < 0 {
		return webUtils.WriteErrorEchoServer(ctx, errors.New(constants.IndexOutOfRange), http.StatusBadRequest)
	}

	obj, err := h.ArtistUseCase.GetById(context.Background(), &commonProto.IdArg{Id: int64(id)})

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

	if _, err := h.ArtistUseCase.Delete(context.Background(), &commonProto.IdArg{Id: int64(id)}); err != nil {
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
	popular, err := h.ArtistUseCase.GetPopular(context.Background(), &emptypb.Empty{})
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

	popular, err := h.TrackUseCase.GetPopularTracksFromArtist(context.Background(), &commonProto.IdArg{Id: int64(id)})
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: popular})
}
