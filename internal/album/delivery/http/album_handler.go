package albumDeliveryHttp

import (
	"errors"
	"fmt"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	AlbumUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/album/useCase"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/album/albumProto"
	Gateway "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/gateway"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/gateway/gatewayProto"
	"github.com/go-park-mail-ru/2022_1_Wave/pkg/webUtils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Handler struct {
	AlbumUseCase AlbumUseCase.AlbumAgent
}

func MakeHandler(album AlbumUseCase.AlbumAgent) Handler {
	return Handler{
		AlbumUseCase: album,
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
	domains, err := h.AlbumUseCase.GetAll()
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
	domains, err := h.AlbumUseCase.GetAllCovers()

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

	if err := Gateway.Check(&result); err != nil {
		return err
	}

	if err := h.AlbumUseCase.Create(&result); err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	lastId, err := h.AlbumUseCase.GetLastId()
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

	if err := Gateway.Check(&result); err != nil {
		return err
	}

	if err := h.AlbumUseCase.CreateCover(&result); err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	lastId, err := h.AlbumUseCase.GetLastCoverId()
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

	if err := Gateway.Check(&result); err != nil {
		return err
	}

	if err := h.AlbumUseCase.Update(&result); err != nil {
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

	if err := Gateway.Check(&result); err != nil {
		return err
	}

	if err := h.AlbumUseCase.UpdateCover(&result); err != nil {
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

	album, err := h.AlbumUseCase.GetById(&gatewayProto.IdArg{Id: int64(id)})

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

	album, err := h.AlbumUseCase.GetCoverById(&gatewayProto.IdArg{Id: int64(id)})

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

	if err := h.AlbumUseCase.Delete(&gatewayProto.IdArg{Id: int64(id)}); err != nil {
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

	if err := h.AlbumUseCase.Delete(&gatewayProto.IdArg{Id: int64(id)}); err != nil {
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
	popular, err := h.AlbumUseCase.GetPopular()
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
// @Description  getting favorites albums
// @Tags         album
// @Accept          application/json
// @Produce      application/json
// @Success      200  {object}  webUtils.Success
// @Failure      400  {object}  webUtils.Error  "Data is invalid"
// @Failure      405  {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/albums/favorites [get]
func (h Handler) GetFavorites(ctx echo.Context) error {
	//todo userId is not 0!!!
	userId := int64(0)
	favorites, err := h.AlbumUseCase.GetFavorites(&gatewayProto.IdArg{Id: userId})
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
// @Tags         album
// @Accept          application/json
// @Produce      application/json
// @Param        albumId  path      int  true  "albumId"
// @Success      200    {object}  webUtils.Success
// @Failure      400    {object}  webUtils.Error  "Data is invalid"
// @Failure      405    {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/albums/favorites/{id} [post]
func (h Handler) AddToFavorites(ctx echo.Context) error {
	trackId, err := strconv.Atoi(ctx.Param(constants.FieldId))
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}
	//todo userId is not 0!!!
	userId := int64(0)

	if _, err := h.AlbumUseCase.AddToFavorites(&gatewayProto.UserIdAlbumIdArg{
		UserId:  userId,
		AlbumId: int64(trackId),
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
// @Description  remove from favorites
// @Tags         album
// @Accept          application/json
// @Produce      application/json
// @Param        albumId  path      int  true  "albumId"
// @Success      200    {object}  webUtils.Success
// @Failure      400    {object}  webUtils.Error  "Data is invalid"
// @Failure      405    {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/albums/favorites/{id} [delete]
func (h Handler) RemoveFromFavorites(ctx echo.Context) error {
	albumId, err := strconv.Atoi(ctx.Param(constants.FieldId))
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}
	//todo userId is not 0!!!
	userId := int64(0)

	if _, err := h.AlbumUseCase.RemoveFromFavorites(&gatewayProto.UserIdAlbumIdArg{
		UserId:  userId,
		AlbumId: int64(albumId),
	}); err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: constants.SuccessRemoveFromFavorites + "(" + fmt.Sprint(albumId) + ")"})
}
