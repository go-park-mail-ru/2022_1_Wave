package test

/*
import (
	"encoding/json"
	"github.com/bxcodec/faker"
	albumDeliveryHttp "github.com/go-park-mail-ru/2022_1_Wave/internal/album/delivery/http"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/album/albumProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/test/mocks"
	"github.com/go-park-mail-ru/2022_1_Wave/pkg/webUtils"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const notUser = int64(-1)

func TestGetEmptyAlbums(t *testing.T) {
	var mockAlbums []*albumProto.AlbumDataTransfer
	e := echo.New()

	req, err := http.NewRequest(echo.GET, "/albums/", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/albums/")

	mockAlbumUseCase := new(mocks.AlbumUseCase)
	mockAlbumUseCase.On("GetAll", notUser).Return(mockAlbums, nil)

	handler := albumDeliveryHttp.Handler{
		AlbumUseCase: mockAlbumUseCase,
	}

	err = handler.GetAll(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	resp := rec.Result()
	body, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)

	var result webUtils.Success
	err = json.Unmarshal(body, &result)
	require.NoError(t, err)

	resultInterface := result.Result.(map[string]interface{})
	require.Equal(t, len(resultInterface), 0)
}

func TestGetAllAlbums(t *testing.T) {
	var mockAlbums []*albumProto.AlbumDataTransfer
	assert.NoError(t, faker.FakeData(&mockAlbums))

	e := echo.New()

	req, err := http.NewRequest(echo.GET, "/albums/", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/albums/")

	mockAlbumUseCase := new(mocks.AlbumUseCase)
	mockAlbumUseCase.On("GetAll", notUser).Return(mockAlbums, nil)

	handler := albumDeliveryHttp.Handler{
		AlbumUseCase: mockAlbumUseCase,
	}

	err = handler.GetAll(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	resp := rec.Result()
	body, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)

	var result webUtils.Success
	err = json.Unmarshal(body, &result)
	require.NoError(t, err)

}*/
