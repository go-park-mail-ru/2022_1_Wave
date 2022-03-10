package api

// config
const (
	Proto             = "http://"
	Host              = "localhost"
	currentApiVersion = v1Prefix
	apiPath           = apiPrefix + currentApiVersion
)

// prefixes
const (
	apiPrefix     = "api/"
	v1Prefix      = "v1/"
	albumsPrefix  = "albums/"
	artistsPrefix = "artists/"
	songsPrefix   = "tracks/"
	usersPrefix   = "users/"
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
	CreateAlbumUrl       = "/" + apiPath + albumsPrefix
	UpdateAlbumUrl       = CreateAlbumUrl
	GetAllAlbumsUrl      = "/" + apiPath + albumsPrefix
	GetAlbumUrlWithoutId = GetAllAlbumsUrl
	GetAlbumUrl          = GetAlbumUrlWithoutId + id
	GetPopularAlbumsUrl  = GetAllAlbumsUrl + popular
	DeleteAlbumUrl       = GetAlbumUrl
)

// artists urls
const (
	CreateArtistUrl       = "/" + apiPath + artistsPrefix
	UpdateArtistUrl       = CreateArtistUrl
	GetAllArtistsUrl      = "/" + apiPath + artistsPrefix
	GetArtistUrlWithoutId = GetAllAlbumsUrl
	GetArtistUrl          = GetArtistUrlWithoutId + id
	GetPopularArtistsUrl  = GetAllArtistsUrl + popular
	DeleteArtistUrl       = GetArtistUrl
)

// tracks urls
const (
	CreateTrackUrl       = "/" + apiPath + songsPrefix
	UpdateTrackUrl       = CreateTrackUrl
	GetAllTracksUrl      = "/" + apiPath + songsPrefix
	GetTrackUrlWithoutId = GetAllTracksUrl
	GetTrackUrl          = GetTrackUrlWithoutId + id
	GetPopularTracksUrl  = GetAllTracksUrl + popular
	DeleteTrackUrl       = GetTrackUrl
)

// auth urls
const (
	LoginUrl       = "/" + apiPath + login
	LogoutUrl      = "/" + apiPath + logout
	SignUpUrl      = "/" + apiPath + signUp
	GetUserUrl     = "/" + apiPath + usersPrefix + id
	GetSelfUserUrl = "/" + apiPath + usersPrefix + self
	GetCSRFAuthUrl = "/" + apiPath + getCSRF
	GetStaticUrl   = "/" + apiPath + "/static/"
)
