package albumDeliveryHttp

import (
	"errors"
	"fmt"
	internal "github.com/go-park-mail-ru/2022_1_Wave/internal"
	AlbumUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/album/useCase"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/album/albumProto"
	Gateway "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/gateway"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/tools/utils"
	user_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/user"
	"github.com/go-park-mail-ru/2022_1_Wave/pkg/webUtils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Handler struct {
	UserUseCase  user_domain.UserUseCase
	AlbumUseCase AlbumUseCase.AlbumUseCase
}

func MakeHandler(album AlbumUseCase.AlbumUseCase, user user_domain.UserUseCase) Handler {
	return Handler{
		UserUseCase:  user,
		AlbumUseCase: album,
	}
}

//func albumCoversToMap(albums []*albumProto.AlbumCoverDataTransfer) map[int64]*albumProto.AlbumCoverDataTransfer {
//	covers := map[int64]*albumProto.AlbumCoverDataTransfer{}
//	for _, obj := range albums {
//		covers[obj.] = obj
//	}
//	return covers
//}

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
	userId, err := internal.GetUserId(ctx, h.UserUseCase)
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}
	albums, err := h.AlbumUseCase.GetAll(userId)
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	if albums == nil {
		albums = []*albumProto.AlbumDataTransfer{}
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: utils.AlbumsToMap(albums)})
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
		domains = []*albumProto.AlbumCoverDataTransfer{}
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
			Result: internal.SuccessCreated + "(" + fmt.Sprint(lastId) + ")"})
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
			Result: internal.SuccessCreated + "(" + fmt.Sprint(lastId) + ")"})
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
			Result: internal.SuccessUpdated + "(" + fmt.Sprint(id) + ")"})
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
			Result: internal.SuccessUpdated + "(" + fmt.Sprint(id) + ")"})
}

// Get godoc
// @Summary      Get
// @Description  getting album by id
// @Tags         album
// @Accept          application/json
// @Produce      application/json
// @Param        id   path      integer  true  "id of album which need to be getted"
// @Success      200  {object}  webUtils.Success
// @Failure      400  {object}  webUtils.Error  "Data is invalid"
// @Failure      405  {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/albums/{id} [get]
func (h Handler) Get(ctx echo.Context) error {
	userId, err := internal.GetUserId(ctx, h.UserUseCase)
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}
	id, err := internal.GetIdInt64ByFieldId(ctx)
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}
	if id < 0 {
		return webUtils.WriteErrorEchoServer(ctx, errors.New(internal.IndexOutOfRange), http.StatusBadRequest)
	}

	album, err := h.AlbumUseCase.GetById(userId, id)

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
// @Success      200  {object}  webUtils.Success
// @Failure      400  {object}  webUtils.Error  "Data is invalid"
// @Failure      405  {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/albumCovers/{id} [get]
func (h Handler) GetCover(ctx echo.Context) error {
	id, err := internal.GetIdInt64ByFieldId(ctx)
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}
	if id < 0 {
		return webUtils.WriteErrorEchoServer(ctx, errors.New(internal.IndexOutOfRange), http.StatusBadRequest)
	}

	album, err := h.AlbumUseCase.GetCoverById(id)

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
	id, err := internal.GetIdInt64ByFieldId(ctx)
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}
	if id < 0 {
		return webUtils.WriteErrorEchoServer(ctx, errors.New(internal.IndexOutOfRange), http.StatusBadRequest)
	}

	if err := h.AlbumUseCase.Delete(id); err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: internal.SuccessDeleted + "(" + fmt.Sprint(id) + ")"})
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
	id, err := internal.GetIdInt64ByFieldId(ctx)
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}
	if id < 0 {
		return webUtils.WriteErrorEchoServer(ctx, errors.New(internal.IndexOutOfRange), http.StatusBadRequest)
	}

	if err := h.AlbumUseCase.DeleteCover(id); err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: internal.SuccessDeleted + "(" + fmt.Sprint(id) + ")"})
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
	userId, err := internal.GetUserId(ctx, h.UserUseCase)
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}
	popular, err := h.AlbumUseCase.GetPopular(userId)
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: utils.AlbumsToMap(popular)})
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
	userId, err := internal.GetUserId(ctx, h.UserUseCase)
	if err != nil {
		return internal.UnauthorizedError(ctx)
	}
	favorites, err := h.AlbumUseCase.GetFavorites(userId)
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: utils.AlbumsToMap(favorites)})
}

type albumIdWrapper struct {
	AlbumId int64 `json:"albumId" example:"4"`
}

// AddToFavorites godoc
// @Summary      AddToFavorites
// @Description  add to favorite
// @Tags         album
// @Accept          application/json
// @Produce      application/json
// @Param        albumId  body      albumIdWrapper  true  "albumId"
// @Success      200    {object}  webUtils.Success
// @Failure      400    {object}  webUtils.Error  "Data is invalid"
// @Failure      405    {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/albums/favorites [post]
func (h Handler) AddToFavorites(ctx echo.Context) error {
	userId, err := internal.GetUserId(ctx, h.UserUseCase)
	if err != nil {
		return internal.UnauthorizedError(ctx)
	}

	holder := albumIdWrapper{}

	if err := ctx.Bind(&holder); err != nil {
		return err
	}

	if err := h.AlbumUseCase.AddToFavorites(userId, holder.AlbumId); err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: internal.SuccessAddedToFavorites + "(" + fmt.Sprint(holder.AlbumId) + ")"})
}

// RemoveFromFavorites godoc
// @Summary      RemoveFromFavorites
// @Description  remove from favorites
// @Tags         album
// @Accept          application/json
// @Produce      application/json
// @Param        id  path      integer  true  "albumId"
// @Success      200    {object}  webUtils.Success
// @Failure      400    {object}  webUtils.Error  "Data is invalid"
// @Failure      405    {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/albums/favorites/{id} [delete]
func (h Handler) RemoveFromFavorites(ctx echo.Context) error {
	userId, err := internal.GetUserId(ctx, h.UserUseCase)
	if err != nil {
		return internal.UnauthorizedError(ctx)
	}

	albumId, err := strconv.ParseInt(ctx.Param(internal.FieldId), 10, 64)
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	if err := h.AlbumUseCase.RemoveFromFavorites(userId, albumId); err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: internal.SuccessRemoveFromFavorites + "(" + fmt.Sprint(albumId) + ")"})
}

// Like godoc
// @Summary      Like
// @Description  like album by id
// @Tags         album
// @Accept          application/json
// @Produce      application/json
// @Param        id   path      integer  true  "id of album which need to be liked"
// @Success      200  {object}  webUtils.Success
// @Failure      400  {object}  webUtils.Error  "Data is invalid"
// @Failure      405  {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/albums/like/{id} [put]
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

	if err := h.AlbumUseCase.Like(id, userId); err != nil {
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
// @Tags         album
// @Accept          application/json
// @Produce      application/json
// @Param        id   path      integer  true  "id of album which need to check for like"
// @Success      200  {object}  webUtils.Success
// @Failure      400  {object}  webUtils.Error  "Data is invalid"
// @Failure      405  {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/albums/like/{id} [get]
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

	liked, err := h.AlbumUseCase.LikeCheckByUser(id, userId)

	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: liked})
}
