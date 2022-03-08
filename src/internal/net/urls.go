package net

// prefixes
const (
	apiPrefix     = "api/"
	v1Prefix      = "v1/"
	albumsPrefix  = "albums/"
	artistsPrefix = "artists/"
	songsPrefix   = "songs/"
	usersPrefix   = "users/"
)

// api config
const (
	currentApiVersion = v1Prefix
	apiPath           = apiPrefix + currentApiVersion
)

// destinations
const (
	login   = "login"
	logout  = "logout"
	signUp  = "signup"
	getCSRF = "get_csrf"
	self    = "self"
	popular = "popular"
	id      = "{id:[0-9]+}"
)

// albums urls
const (
	createAlbumUrl      = "/" + apiPath + albumsPrefix
	updateAlbumUrl      = createAlbumUrl
	getAllAlbumsUrl     = "/" + apiPath + albumsPrefix
	getAlbumUrl         = getAllAlbumsUrl + id
	getPopularAlbumsUrl = getAllAlbumsUrl + popular
	deleteAlbumUrl      = getAlbumUrl
)

// artists urls
const (
	createArtistUrl      = "/" + apiPath + artistsPrefix
	updateArtistUrl      = createArtistUrl
	getAllArtistsUrl     = "/" + apiPath + artistsPrefix
	getArtistUrl         = getAllArtistsUrl + id
	getPopularArtistsUrl = getAllArtistsUrl + popular
	deleteArtistUrl      = getArtistUrl
)

// songs urls
const (
	createSongUrl      = "/" + apiPath + songsPrefix
	updateSongUrl      = createSongUrl
	getAllSongsUrl     = "/" + apiPath + songsPrefix
	getSongUrl         = getAllSongsUrl + id
	getPopularSongsUrl = getAllSongsUrl + popular
	deleteSongUrl      = getSongUrl
)

// auth urls
const (
	loginUrl       = "/" + apiPath + login
	logoutUrl      = "/" + apiPath + logout
	signUpUrl      = "/" + apiPath + signUp
	getUserUrl     = "/" + apiPath + usersPrefix + id
	getSelfUserUrl = "/" + apiPath + usersPrefix + self
	getCSRFAuthUrl = "/" + apiPath + getCSRF
	getStaticUrl   = "/" + apiPath + "/static/"
)
