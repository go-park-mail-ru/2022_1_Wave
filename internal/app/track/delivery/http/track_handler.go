package trackDeliveryHttp

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/structs/delivery/http"
	"github.com/labstack/echo/v4"
)

var Handler structsDeliveryHttp.Handler

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
func GetAll(ctx echo.Context) error {
	return Handler.GetAll(ctx, domain.TrackMutex)
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
func Create(ctx echo.Context) error {
	proxy, err := Handler.Create(ctx, domain.TrackMutex)
	Handler = proxy.(structsDeliveryHttp.Handler)
	return err
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
func Update(ctx echo.Context) error {
	proxy, err := Handler.Update(ctx, domain.TrackMutex)
	Handler = proxy.(structsDeliveryHttp.Handler)
	return err
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
func Get(ctx echo.Context) error {
	return Handler.Get(ctx, domain.TrackMutex)
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
func Delete(ctx echo.Context) error {
	proxy, err := Handler.Delete(ctx, domain.TrackMutex)
	Handler = proxy.(structsDeliveryHttp.Handler)
	return err
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
func GetPopular(ctx echo.Context) error {
	return Handler.GetPopular(ctx, domain.TrackMutex)
}
