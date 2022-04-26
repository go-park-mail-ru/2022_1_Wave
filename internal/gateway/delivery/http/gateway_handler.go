package gatewayDeliveryHttp

import (
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	AlbumUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/album/useCase"
	ArtistUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/artist/useCase"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/album/albumProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/artist/artistProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/gateway/gatewayProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/track/trackProto"
	TrackUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/track/useCase"
	"github.com/go-park-mail-ru/2022_1_Wave/pkg/webUtils"
	"github.com/labstack/echo/v4"
	"net/http"
	"sync"
)

type Handler struct {
	ArtistUseCase ArtistUseCase.ArtistAgent
	AlbumUseCase  AlbumUseCase.AlbumAgent
	TrackUseCase  TrackUseCase.TrackAgent
}

func MakeHandler(album AlbumUseCase.AlbumAgent, artist ArtistUseCase.ArtistAgent, track TrackUseCase.TrackAgent) Handler {
	return Handler{
		ArtistUseCase: artist,
		AlbumUseCase:  album,
		TrackUseCase:  track,
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
	searchString := ctx.Param(constants.FieldToFind)

	albumsChan := make(chan *albumProto.AlbumsResponse, 1)
	artistsChan := make(chan *artistProto.ArtistsResponse, 1)
	tracksChan := make(chan *trackProto.TracksResponse, 1)
	albumsErrorChan := make(chan error, 1)
	artistsErrorChan := make(chan error, 1)
	tracksErrorChan := make(chan error, 1)

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func(albumsChan chan *albumProto.AlbumsResponse, errorChan chan error, wg *sync.WaitGroup) {
		defer wg.Done()
		albums, err := h.AlbumUseCase.SearchByTitle(&gatewayProto.StringArg{Str: searchString})
		if err != nil {
			errorChan <- err
		} else {
			albumsChan <- albums
		}
		close(albumsChan)
		close(errorChan)
	}(albumsChan, albumsErrorChan, wg)

	wg.Add(1)
	go func(artistsChan chan *artistProto.ArtistsResponse, errorChan chan error, wg *sync.WaitGroup) {
		wg.Done()
		artists, err := h.ArtistUseCase.SearchByName(&gatewayProto.StringArg{Str: searchString})
		if err != nil {
			errorChan <- err
		} else {
			artistsChan <- artists
		}
		close(artistsChan)
		close(errorChan)
	}(artistsChan, artistsErrorChan, wg)

	wg.Add(1)
	go func(tracksChan chan *trackProto.TracksResponse, errorChan chan error, wg *sync.WaitGroup) {
		wg.Done()
		tracks, err := h.TrackUseCase.SearchByTitle(&gatewayProto.StringArg{Str: searchString})
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
		Albums:  albums.Albums,
		Artists: artists.Artists,
		Tracks:  tracks.Tracks,
	}

	return ctx.JSON(http.StatusOK,
		webUtils.Success{
			Status: webUtils.OK,
			Result: result})
}
