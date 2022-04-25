package albumDeliveryHttp

import (
	"context"
	"errors"
	"fmt"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	AlbumRepo "github.com/go-park-mail-ru/2022_1_Wave/internal/app/album/repository"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/microservices/album/albumProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/microservices/common/commonProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/microservices/track/trackProto"
	"github.com/go-park-mail-ru/2022_1_Wave/pkg/webUtils"
	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
	"strconv"
)

type Handler struct {
	AlbumUseCase albumProto.AlbumUseCaseClient
	TrackUseCase trackProto.TrackUseCaseClient
}

func MakeHandler(album albumProto.AlbumUseCaseClient, track trackProto.TrackUseCaseClient) Handler {
	return Handler{
		AlbumUseCase: album,
		TrackUseCase: track,
	}

}

// GetAll godoc
// @Summary      GetAll
// @Description  getting all albums
// @Tags         album
// @Accept          application/json
// @Produce      application/json
// @Success      200  {object}  webUtils.Success
// @Failure      400  {object}  webUtils.Error  "Data is invalid"
// @Failure      405  {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/albums/ [get]
func (h Handler) GetAll(ctx echo.Context) error {
	domains, err := h.AlbumUseCase.GetAll(context.Background(), &emptypb.Empty{})
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	if domains == nil {
		domains = &albumProto.AlbumsResponse{}
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: domains})
}

// GetAllCovers godoc
// @Summary      GetAll
// @Description  getting all albums cover
// @Tags         albumCover
// @Accept          application/json
// @Produce      application/json
// @Success      200  {object}  webUtils.Success
// @Failure      400  {object}  webUtils.Error  "Data is invalid"
// @Failure      405  {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/albumCovers/ [get]
func (h Handler) GetAllCovers(ctx echo.Context) error {
	domains, err := h.AlbumUseCase.GetAllCovers(context.Background(), &emptypb.Empty{})

	fmt.Println("domains=", domains)

	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	if domains == nil {
		domains = &albumProto.AlbumsCoverResponse{}
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: domains})
}

// Create godoc
// @Summary      Create
// @Description  creating new album
// @Tags         album
// @Accept          application/json
// @Produce      application/json
// @Param        Album  body      albumProto.Album  true  "params of new album. Id will be set automatically."
// @Success      200    {object}  webUtils.Success
// @Failure      400    {object}  webUtils.Error  "Data is invalid"
// @Failure      405    {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/albums/ [post]
func (h Handler) Create(ctx echo.Context) error {
	result := albumProto.Album{}

	if err := ctx.Bind(&result); err != nil {
		return err
	}

	if err := AlbumRepo.Check(&result); err != nil {
		return err
	}

	if _, err := h.AlbumUseCase.Create(context.Background(), &result); err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	lastId, err := h.AlbumUseCase.GetLastId(context.Background(), &emptypb.Empty{})
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: constants.SuccessCreated + "(" + fmt.Sprint(lastId) + ")"})
}

// CreateCover godoc
// @Summary      Create
// @Description  creating new albumCover
// @Tags         albumCover
// @Accept       application/json
// @Produce      application/json
// @Param        Album  body      albumProto.AlbumCover  true  "params of new album cover. Id will be set automatically."
// @Success      200    {object}  webUtils.Success
// @Failure      400    {object}  webUtils.Error  "Data is invalid"
// @Failure      405    {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/albumCovers/ [post]
func (h Handler) CreateCover(ctx echo.Context) error {
	result := albumProto.AlbumCover{}

	if err := ctx.Bind(&result); err != nil {
		return err
	}

	if err := AlbumRepo.Check(&result); err != nil {
		return err
	}

	if _, err := h.AlbumUseCase.CreateCover(context.Background(), &result); err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	lastId, err := h.AlbumUseCase.GetLastCoverId(context.Background(), &emptypb.Empty{})
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
// @Description  updating album by id
// @Tags         albumCover
// @Accept          application/json
// @Produce      application/json
// @Param        Album  body      albumProto.Album  true  "id of updating album and params of it."
// @Success      200    {object}  webUtils.Success
// @Failure      400    {object}  webUtils.Error  "Data is invalid"
// @Failure      405    {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/albums/ [put]
func (h Handler) Update(ctx echo.Context) error {
	result := albumProto.Album{}

	if err := ctx.Bind(&result); err != nil {
		return err
	}

	if err := AlbumRepo.Check(&result); err != nil {
		return err
	}

	if _, err := h.AlbumUseCase.Update(context.Background(), &result); err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	id := result.Id
	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: constants.SuccessUpdated + "(" + fmt.Sprint(id) + ")"})
}

// UpdateCover godoc
// @Summary      Update
// @Description  updating album cover by id
// @Tags         albumCover
// @Accept          application/json
// @Produce      application/json
// @Param        AlbumCover  body      albumProto.AlbumCover  true  "id of updating album cover and params of it."
// @Success      200    {object}  webUtils.Success
// @Failure      400    {object}  webUtils.Error  "Data is invalid"
// @Failure      405    {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/albumCovers/ [put]
func (h Handler) UpdateCover(ctx echo.Context) error {
	result := albumProto.AlbumCover{}

	if err := ctx.Bind(&result); err != nil {
		return err
	}

	if err := AlbumRepo.Check(&result); err != nil {
		return err
	}

	if _, err := h.AlbumUseCase.UpdateCover(context.Background(), &result); err != nil {
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
// @Description  getting album by id
// @Tags         album
// @Accept          application/json
// @Produce      application/json
// @Param        id   path      integer  true  "id of album which need to be getted"
// @Success      200  {object}  albumProto.Album
// @Failure      400  {object}  webUtils.Error  "Data is invalid"
// @Failure      405  {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/albums/{id} [get]
func (h Handler) Get(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param(constants.FieldId))
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}
	if id < 0 {
		return webUtils.WriteErrorEchoServer(ctx, errors.New(constants.IndexOutOfRange), http.StatusBadRequest)
	}

	album, err := h.AlbumUseCase.GetById(context.Background(), &commonProto.IdArg{Id: int64(id)})

	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: album})
}

// GetCover godoc
// @Summary      Get
// @Description  getting album cover by id
// @Tags         albumCover
// @Accept       application/json
// @Produce      application/json
// @Param        id   path      integer  true  "id of album cover which need to be getted"
// @Success      200  {object}  albumProto.AlbumCover
// @Failure      400  {object}  webUtils.Error  "Data is invalid"
// @Failure      405  {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/albumCovers/{id} [get]
func (h Handler) GetCover(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param(constants.FieldId))
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}
	if id < 0 {
		return webUtils.WriteErrorEchoServer(ctx, errors.New(constants.IndexOutOfRange), http.StatusBadRequest)
	}

	album, err := h.AlbumUseCase.GetCoverById(context.Background(), &commonProto.IdArg{Id: int64(id)})

	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: album})
}

// Delete godoc
// @Summary      Delete
// @Description  deleting album by id
// @Tags         album
// @Accept          application/json
// @Produce      application/json
// @Param        id   path      integer  true  "id of album which need to be deleted"
// @Success      200  {object}  webUtils.Success
// @Failure      400  {object}  webUtils.Error  "Data is invalid"
// @Failure      405  {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/albums/{id} [delete]
func (h Handler) Delete(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param(constants.FieldId))
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}
	if id < 0 {
		return webUtils.WriteErrorEchoServer(ctx, errors.New(constants.IndexOutOfRange), http.StatusBadRequest)
	}

	if _, err := h.AlbumUseCase.Delete(context.Background(), &commonProto.IdArg{Id: int64(id)}); err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: constants.SuccessDeleted + "(" + fmt.Sprint(id) + ")"})
}

// DeleteCover godoc
// @Summary      Delete
// @Description  deleting album cover by id
// @Tags         albumCover
// @Accept          application/json
// @Produce      application/json
// @Param        id   path      integer  true  "id of album cover which need to be deleted"
// @Success      200  {object}  webUtils.Success
// @Failure      400  {object}  webUtils.Error  "Data is invalid"
// @Failure      405  {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/albumCovers/{id} [delete]
func (h Handler) DeleteCover(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param(constants.FieldId))
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}
	if id < 0 {
		return webUtils.WriteErrorEchoServer(ctx, errors.New(constants.IndexOutOfRange), http.StatusBadRequest)
	}

	if _, err := h.AlbumUseCase.Delete(context.Background(), &commonProto.IdArg{Id: int64(id)}); err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: constants.SuccessDeleted + "(" + fmt.Sprint(id) + ")"})
}

// GetPopular godoc
// @Summary      GetPopular
// @Description  getting top20 popular albums
// @Tags         album
// @Accept          application/json
// @Produce      application/json
// @Success      200  {object}  webUtils.Success
// @Failure      400  {object}  webUtils.Error  "Data is invalid"
// @Failure      405  {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/albums/popular [get]
func (h Handler) GetPopular(ctx echo.Context) error {
	popular, err := h.AlbumUseCase.GetPopular(context.Background(), &emptypb.Empty{})
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: popular})
}
