package albumDeliveryHttp

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/structs/delivery/http"
	"github.com/labstack/echo/v4"
)

var Handler structsDeliveryHttp.Handler

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
func GetAll(ctx echo.Context) error {
	return Handler.GetAll(ctx, domain.AlbumMutex)
}

// Create godoc
// @Summary      Create
// @Description  creating new album
// @Tags         album
// @Accept          application/json
// @Produce      application/json
// @Param        Album  body      domain.Album  true  "params of new album. Id will be set automatically."
// @Success      200    {object}  webUtils.Success
// @Failure      400    {object}  webUtils.Error  "Data is invalid"
// @Failure      405    {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/albums/ [post]
func Create(ctx echo.Context) error {
	proxy, err := Handler.Create(ctx, domain.AlbumMutex)
	Handler = proxy.(structsDeliveryHttp.Handler)
	return err
}

// Update godoc
// @Summary      Update
// @Description  updating album by id
// @Tags         album
// @Accept          application/json
// @Produce      application/json
// @Param        Album  body      domain.Album  true  "id of updating album and params of it."
// @Success      200    {object}  webUtils.Success
// @Failure      400    {object}  webUtils.Error  "Data is invalid"
// @Failure      405    {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/albums/ [put]
func Update(ctx echo.Context) error {
	proxy, err := Handler.Update(ctx, domain.AlbumMutex)
	Handler = proxy.(structsDeliveryHttp.Handler)
	return err
}

// Get godoc
// @Summary      Get
// @Description  getting album by id
// @Tags         album
// @Accept          application/json
// @Produce      application/json
// @Param        id   path      integer  true  "id of album which need to be getted"
// @Success      200  {object}  domain.Album
// @Failure      400  {object}  webUtils.Error  "Data is invalid"
// @Failure      405  {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/albums/{id} [get]
func Get(ctx echo.Context) error {
	return Handler.Get(ctx, domain.AlbumMutex)
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
func Delete(ctx echo.Context) error {
	proxy, err := Handler.Delete(ctx, domain.AlbumMutex)
	Handler = proxy.(structsDeliveryHttp.Handler)
	return err
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
func GetPopular(ctx echo.Context) error {
	return Handler.GetPopular(ctx, domain.AlbumMutex)
}
