package artistDeliveryHttp

import (
	"fmt"
	artistUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/artist/usecase"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	utilsInterfaces "github.com/go-park-mail-ru/2022_1_Wave/internal/app/interfaces"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/structs/delivery/http"
	dataTransferCreator "github.com/go-park-mail-ru/2022_1_Wave/internal/app/tools/dataTransfer"
	"github.com/go-park-mail-ru/2022_1_Wave/pkg/webUtils"
	"github.com/labstack/echo/v4"
	"net/http"
	"sync"
)

var Handler structsDeliveryHttp.Handler

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
func GetAll(ctx echo.Context) error {
	return Handler.GetAll(ctx, domain.ArtistMutex)
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
func Create(ctx echo.Context) error {
	proxy, err := Handler.Create(ctx, domain.ArtistMutex)
	Handler = proxy.(structsDeliveryHttp.Handler)
	return err
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
func Update(ctx echo.Context) error {
	proxy, err := Handler.Update(ctx, domain.ArtistMutex)
	Handler = proxy.(structsDeliveryHttp.Handler)
	return err
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
func Get(ctx echo.Context) error {
	return Handler.Get(ctx, domain.ArtistMutex)
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
func Delete(ctx echo.Context) error {
	proxy, err := Handler.Delete(ctx, domain.ArtistMutex)
	Handler = proxy.(structsDeliveryHttp.Handler)
	return err
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
func GetPopular(ctx echo.Context) error {
	return Handler.GetPopular(ctx, domain.ArtistMutex)
}

func GetPopularTracksHandler(ctx echo.Context, mutex *sync.RWMutex, useCase utilsInterfaces.UseCaseInterface) error {
	id, err := structsDeliveryHttp.ReadGetDeleteRequest(ctx)

	fmt.Println(id)

	popular, err := useCase.GetPopularTracksFromArtist(uint64(id), mutex)
	if err != nil {
		return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
	}

	popularData := popular.([]domain.Track)
	dataTransfers := make([]domain.TrackDataTransfer, len(popularData))

	for i, pop := range popularData {
		dataTransfer, err := dataTransferCreator.CreateDataTransfer(pop)
		if err != nil {
			return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
		}

		dataTransfers[i] = dataTransfer.(domain.TrackDataTransfer)

	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: dataTransfers})
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
func GetPopularTracks(ctx echo.Context) error {
	return GetPopularTracksHandler(ctx, domain.ArtistMutex, artistUseCase.UseCase)
}
