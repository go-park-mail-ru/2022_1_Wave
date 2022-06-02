package router

import (
	_ "github.com/go-park-mail-ru/2022_1_Wave/docs"
	"github.com/go-park-mail-ru/2022_1_Wave/init/logger"
	albumDeliveryHttp "github.com/go-park-mail-ru/2022_1_Wave/internal/album/delivery/http"
	AlbumUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/album/useCase"
	artistDeliveryHttp "github.com/go-park-mail-ru/2022_1_Wave/internal/artist/delivery/http"
	ArtistUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/artist/useCase"
	auth_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/auth"
	authHttp "github.com/go-park-mail-ru/2022_1_Wave/internal/auth/delivery/http"
	auth_middleware "github.com/go-park-mail-ru/2022_1_Wave/internal/auth/delivery/http/middleware"
	gatewayDeliveryHttp "github.com/go-park-mail-ru/2022_1_Wave/internal/gateway/delivery/http"
	linkerDeliveryHttp "github.com/go-park-mail-ru/2022_1_Wave/internal/linker/delivery/http"
	LinkerUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/linker/useCase"
	playlistDeliveryHttp "github.com/go-park-mail-ru/2022_1_Wave/internal/playlist/delivery/http"
	PlaylistUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/playlist/useCase"
	trackDeliveryHttp "github.com/go-park-mail-ru/2022_1_Wave/internal/track/delivery/http"
	TrackUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/track/useCase"
	user_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/user"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/user/client/s3"
	userHttp "github.com/go-park-mail-ru/2022_1_Wave/internal/user/delivery/http"
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
)

func Router(e *echo.Echo,
	auth auth_domain.AuthUseCase,
	album AlbumUseCase.AlbumUseCase,
	artist ArtistUseCase.ArtistUseCase,
	track TrackUseCase.TrackUseCase,
	playlist PlaylistUseCase.PlaylistUseCase,
	user user_domain.UserUseCase,
	linker LinkerUseCase.LinkerUseCase,
	s3Handler *s3.Handler) error {

	api := e.Group(apiPrefix)
	v1 := api.Group(v1Prefix)
	albumHandler := albumDeliveryHttp.MakeHandler(album, user)
	artistHandler := artistDeliveryHttp.MakeHandler(artist, track, user)
	trackHandler := trackDeliveryHttp.MakeHandler(track, user)
	playlistHandler := playlistDeliveryHttp.MakeHandler(playlist, user)
	authHandler := authHttp.MakeHandler(auth)
	linkerHandler := linkerDeliveryHttp.MakeHandler(linker)
	userHandler := userHttp.MakeHandler(user, s3Handler)
	gatewayHandler := gatewayDeliveryHttp.MakeHandler(album, artist, track, user)

	m := auth_middleware.InitMiddleware(auth)

	logger.GlobalLogger.Logrus.Warnln("api version:", v1Prefix)

	SetLinkerRoutes(e, linkerHandler)
	logger.GlobalLogger.Logrus.Warnln("setting linker routes")

	SetAlbumsRoutes(v1, albumHandler)
	logger.GlobalLogger.Logrus.Warnln("setting albums routes")

	SetArtistsRoutes(v1, artistHandler, trackHandler)
	logger.GlobalLogger.Logrus.Warnln("setting artists routes")

	SetTracksRoutes(v1, trackHandler)
	logger.GlobalLogger.Logrus.Warnln("setting tracks routes")

	SetPlaylistsRoutes(v1, playlistHandler)
	logger.GlobalLogger.Logrus.Warnln("setting playlists routes")

	SetGatewayRoutes(v1, gatewayHandler)
	logger.GlobalLogger.Logrus.Warnln("setting gateway routes")

	SetDocsPath(v1)
	logger.GlobalLogger.Logrus.Warnln("setting docs routes")

	SetAuthRoutes(v1, authHandler, m)
	logger.GlobalLogger.Logrus.Warnln("setting auth routes")

	SetUserRoutes(v1, userHandler, m)
	logger.GlobalLogger.Logrus.Warnln("setting user routes")

	return nil
}

// SetAlbumsRoutes albums
func SetAlbumsRoutes(apiVersion *echo.Group, handler albumDeliveryHttp.Handler) {
	albumRoutes := apiVersion.Group(albumsPrefix)
	albumRoutes.GET(idEchoPattern, handler.Get)
	albumRoutes.GET(locate, handler.GetAll)
	albumRoutes.POST(locate, handler.Create)
	albumRoutes.PUT(locate, handler.Update)
	albumRoutes.GET(popularPrefix, handler.GetPopular)
	albumRoutes.GET(popularOfWeekPrefix, handler.GetPopularOfWeek)
	albumRoutes.GET(favoritesPrefix, handler.GetFavorites)
	albumRoutes.POST(favoritesPrefix, handler.AddToFavorites)
	albumRoutes.DELETE(favoritesPrefix+idEchoPattern, handler.RemoveFromFavorites)
	albumRoutes.DELETE(idEchoPattern, handler.Delete)
	albumRoutes.PUT(likePrefix+idEchoPattern, handler.Like)
	albumRoutes.GET(likePrefix+idEchoPattern, handler.LikeCheckByUser)

	coverRoutes := apiVersion.Group(albumCoversPrefix)
	coverRoutes.GET(idEchoPattern, handler.GetCover)
	coverRoutes.GET(locate, handler.GetAllCovers)
	coverRoutes.POST(locate, handler.CreateCover)
	coverRoutes.PUT(locate, handler.UpdateCover)
	coverRoutes.DELETE(idEchoPattern, handler.DeleteCover)

}

// SetArtistsRoutes artists
func SetArtistsRoutes(apiVersion *echo.Group, handler artistDeliveryHttp.Handler, trackHandler trackDeliveryHttp.Handler) {
	artistRoutes := apiVersion.Group(artistsPrefix)

	artistRoutes.GET(idEchoPattern, handler.Get)
	artistRoutes.GET(locate, handler.GetAll)
	artistRoutes.POST(locate, handler.Create)
	artistRoutes.PUT(locate, handler.Update)
	artistRoutes.GET(popularPrefix, handler.GetPopular)
	artistRoutes.GET(idEchoPattern+popularPrefix, trackHandler.GetPopularTracks)
	artistRoutes.GET(favoritesPrefix, handler.GetFavorites)
	artistRoutes.POST(favoritesPrefix, handler.AddToFavorites)
	artistRoutes.DELETE(favoritesPrefix+idEchoPattern, handler.RemoveFromFavorites)
	artistRoutes.DELETE(idEchoPattern, handler.Delete)
	artistRoutes.PUT(likePrefix+idEchoPattern, handler.Like)
	artistRoutes.GET(likePrefix+idEchoPattern, handler.LikeCheckByUser)
}

// SetTracksRoutes tracks
func SetTracksRoutes(apiVersion *echo.Group, handler trackDeliveryHttp.Handler) {
	trackRoutes := apiVersion.Group(tracksPrefix)

	trackRoutes.GET(idEchoPattern, handler.Get)
	trackRoutes.GET(locate, handler.GetAll)
	trackRoutes.POST(locate, handler.Create)
	trackRoutes.PUT(locate, handler.Update)
	trackRoutes.GET(popularPrefix, handler.GetPopular)
	trackRoutes.GET(popularOfWeekPrefix, handler.GetPopularOfWeek)
	trackRoutes.GET(favoritesPrefix, handler.GetFavorites)
	trackRoutes.POST(favoritesPrefix, handler.AddToFavorites)
	trackRoutes.DELETE(favoritesPrefix+idEchoPattern, handler.RemoveFromFavorites)
	trackRoutes.DELETE(idEchoPattern, handler.Delete)
	trackRoutes.PUT(likePrefix+idEchoPattern, handler.Like)
	trackRoutes.GET(likePrefix+idEchoPattern, handler.LikeCheckByUser)
	trackRoutes.PUT(listenPrefix+idEchoPattern, handler.Listen)
	trackRoutes.GET(playlistPrefix+idEchoPattern, handler.GetTracksFromPlaylist)
}

// SetPlaylistsRoutes tracks
func SetPlaylistsRoutes(apiVersion *echo.Group, handler playlistDeliveryHttp.Handler) {
	playlistRoutes := apiVersion.Group(playlistPrefix)

	playlistRoutes.GET(locate, handler.GetAll)
	playlistRoutes.GET(ofUserPrefix, handler.GetAllOfCurrentUser)
	playlistRoutes.POST(locate, handler.Create)
	playlistRoutes.PUT(locate, handler.Update)
	playlistRoutes.GET(idEchoPattern, handler.Get)
	playlistRoutes.GET(ofUserPrefix+idEchoPattern, handler.GetOfCurrentUser)
	playlistRoutes.DELETE(idEchoPattern, handler.Delete)
	playlistRoutes.POST(ofUserPrefix, handler.AddToPlaylist)
	playlistRoutes.DELETE(ofUserPrefix, handler.RemoveFromPlaylist)
}

// SetGatewayRoutes songs
func SetGatewayRoutes(apiVersion *echo.Group, handler gatewayDeliveryHttp.Handler) {
	searchRoutes := apiVersion.Group(searchPrefix)
	searchRoutes.GET(strEchoToFindPattern, handler.Search)
}

// SetLinkerRoutes songs
func SetLinkerRoutes(e *echo.Echo, handler linkerDeliveryHttp.Handler) {
	e.GET(strEchoHashPattern, handler.Get)
	e.POST(locate, handler.Create)
	e.GET(strCountPattern+strEchoHashPattern, handler.Count)
}

func SetUserRoutes(apiVersion *echo.Group, handler userHttp.UserHandler, m *auth_middleware.HttpMiddleware) {
	userRoutes := apiVersion.Group(usersPrefix)

	userRoutes.GET("/:id", handler.GetUser)
	userRoutes.GET("/self", handler.GetSelfUser, m.Auth, m.CSRF)

	userRoutes.PATCH("/self", handler.UpdateSelfUser, m.Auth, m.CSRF)
	userRoutes.PATCH("/upload_avatar", handler.UploadAvatar, m.Auth, m.CSRF)
}

// InitAuthModule auth
func SetAuthRoutes(apiVersion *echo.Group, handler authHttp.AuthHandler, m *auth_middleware.HttpMiddleware) {
	apiVersion.POST(loginPrefix, handler.Login, m.CSRF)
	apiVersion.POST(logoutPrefix, handler.Logout, m.CSRF)
	apiVersion.POST(signUpPrefix, handler.SignUp, m.CSRF)
	apiVersion.GET(getCSRFPrefix, handler.GetCSRF)
}

// SetDocsPath docs
func SetDocsPath(apiVersion *echo.Group) {
	docRoutes := apiVersion.Group(docsPrefix)
	docRoutes.GET(locate+"*", echoSwagger.WrapHandler)
}

// prefixes
const (
	apiPrefix           = "/api"
	v1Prefix            = "/v1"
	albumsPrefix        = "/albums"
	albumCoversPrefix   = "/albumCovers"
	artistsPrefix       = "/artists"
	tracksPrefix        = "/tracks"
	usersPrefix         = "/users"
	searchPrefix        = "/search"
	docsPrefix          = "/docs"
	popularPrefix       = "/popular"
	popularOfWeekPrefix = "/popular/week"
	playlistPrefix      = "/playlists"
	favoritesPrefix     = "/favorites"
	likePrefix          = "/like"
	listenPrefix        = "/listen"
	loginPrefix         = "/login"
	logoutPrefix        = "/logout"
	signUpPrefix        = "/signup"
	getCSRFPrefix       = "/get_csrf"
	ofUserPrefix        = "/ofUser"
)

const (
	locate = "/"
)

// destinations
const (
	idEchoPattern        = "/:id"
	strEchoToFindPattern = "/:toFind"
	strEchoHashPattern   = "/:hash"
	strCountPattern      = "/count"
)
