package gatewayDeliveryHttp

import (
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	AlbumUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/album/useCase"
	ArtistUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/artist/useCase"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/album/albumProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/artist/artistProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/track/trackProto"
	TrackUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/track/useCase"
	user_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/user"
	"github.com/go-park-mail-ru/2022_1_Wave/pkg/webUtils"
	"github.com/labstack/echo/v4"
	"net/http"
	"sync"
)

type Handler struct {
	ArtistUseCase ArtistUseCase.UseCase
	AlbumUseCase  AlbumUseCase.UseCase
	TrackUseCase  TrackUseCase.UseCase
	UserUseCase   user_domain.UserUseCase
}

func MakeHandler(album AlbumUseCase.UseCase, artist ArtistUseCase.UseCase, track TrackUseCase.UseCase, user user_domain.UserUseCase) Handler {
	return Handler{
		ArtistUseCase: artist,
		AlbumUseCase:  album,
		TrackUseCase:  track,
		UserUseCase:   user,
	}
}

type SearchResult struct {
	Albums  []*albumProto.AlbumDataTransfer   `json:"MatchedAlbums"`
	Artists []*artistProto.ArtistDataTransfer `json:"MatchedArtists"`
	Tracks  []*trackProto.TrackDataTransfer   `json:"MatchedTracks"`
}

// Search godoc
// @Summary      Search
// @Description  general search
// @Tags         gateway
// @Accept          application/json
// @Produce      application/json
// @Param        toFind   path      string  true  "string for search"
// @Success      200  {object}  webUtils.Success
// @Failure      400  {object}  webUtils.Error  "Data is invalid"
// @Failure      405  {object}  webUtils.Error  "Method is not allowed"
// @Router       /api/v1/search/{toFind} [get]
func (h Handler) Search(ctx echo.Context) error {
	userId, _ := constants.GetUserId(ctx, h.UserUseCase)

	searchString := ctx.Param(constants.FieldToFind)

	albumsChan := make(chan []*albumProto.AlbumDataTransfer, 1)
	artistsChan := make(chan []*artistProto.ArtistDataTransfer, 1)
	tracksChan := make(chan []*trackProto.TrackDataTransfer, 1)
	albumsErrorChan := make(chan error, 1)
	artistsErrorChan := make(chan error, 1)
	tracksErrorChan := make(chan error, 1)

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func(albumsChan chan []*albumProto.AlbumDataTransfer, errorChan chan error, wg *sync.WaitGroup) {
		defer wg.Done()
		albums, err := h.AlbumUseCase.SearchByTitle(userId, searchString)
		if err != nil {
			errorChan <- err
		} else {
			albumsChan <- albums
		}
		close(albumsChan)
		close(errorChan)
	}(albumsChan, albumsErrorChan, wg)

	wg.Add(1)
	go func(artistsChan chan []*artistProto.ArtistDataTransfer, errorChan chan error, wg *sync.WaitGroup) {
		wg.Done()
		artists, err := h.ArtistUseCase.SearchByName(userId, searchString)
		if err != nil {
			errorChan <- err
		} else {
			artistsChan <- artists
		}
		close(artistsChan)
		close(errorChan)
	}(artistsChan, artistsErrorChan, wg)

	wg.Add(1)
	go func(tracksChan chan []*trackProto.TrackDataTransfer, errorChan chan error, wg *sync.WaitGroup) {
		wg.Done()
		tracks, err := h.TrackUseCase.SearchByTitle(searchString, userId)
		if err != nil {
			errorChan <- err
		} else {
			tracksChan <- tracks
		}
		close(tracksChan)
		close(errorChan)
	}(tracksChan, tracksErrorChan, wg)

	wg.Wait()

	for err := range albumsErrorChan {
		if err != nil {
			return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
		}
	}

	for err := range artistsErrorChan {
		if err != nil {
			return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
		}
	}

	for err := range tracksErrorChan {
		if err != nil {
			return webUtils.WriteErrorEchoServer(ctx, err, http.StatusBadRequest)
		}
	}

	albums := <-albumsChan
	artists := <-artistsChan
	tracks := <-tracksChan

	result := SearchResult{
		Albums:  albums,
		Artists: artists,
		Tracks:  tracks,
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: result})
}
