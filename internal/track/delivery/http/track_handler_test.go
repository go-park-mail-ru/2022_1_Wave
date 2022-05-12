package trackDeliveryHttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bxcodec/faker"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/track/trackProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/tools"
	mocks2 "github.com/go-park-mail-ru/2022_1_Wave/internal/track/mocks"
	"strconv"

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

func TestGetAllTracks(t *testing.T) {
	//var mockTracks = &trackProto.TracksResponse{Tracks: make([]*trackProto.Track, 10)}
	var mockTracks = make([]*trackProto.TrackDataTransfer, 10)
	for ind, _ := range mockTracks {
		mockTracks[ind] = new(trackProto.TrackDataTransfer)
	}

	err := faker.FakeData(&mockTracks)
	assert.NoError(t, err)

	e := echo.New()

	req, err := http.NewRequest(echo.GET, "/tracks/", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/tracks/")

	mockTrackUseCase := new(mocks2.UseCase)

	mockTrackUseCase.On("GetAll").Return(mockTracks, nil)

	handler := Handler{
		TrackUseCase: mockTrackUseCase,
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

	/*resultMap := result.Result.([]interface{})
	for _, resMap := range resultMap {
		cast := resMap.(map[string]interface{})
		fmt.Println(cast)
	}
	fmt.Println(resultMap)*/
	//tracks := resultMap["tracks"]

	for idx, obj := range result.Result.([]interface{}) {
		track, err := tools.CreateTrackDataTransferFromInterface(obj)
		require.NoError(t, err)
		require.Equal(t, mockTracks[idx], track)
	}
}

func TestGetTrack(t *testing.T) {
	var mockTrack trackProto.TrackDataTransfer
	err := faker.FakeData(&mockTrack)
	assert.NoError(t, err)

	e := echo.New()

	req, err := http.NewRequest(echo.GET, "/tracks/"+strconv.Itoa(int(mockTrack.Id)), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/tracks/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(int(mockTrack.Id)))

	mockTrackUseCase := new(mocks2.UseCase)

	//mockTrackUseCase.On("GetById", &gatewayProto.IdArg{Id: mockTrack.Id}).Return(&mockTrack, nil)
	mockTrackUseCase.On("GetById", mockTrack.Id).Return(&mockTrack, nil)

	handler := Handler{
		TrackUseCase: mockTrackUseCase,
	}

	err = handler.Get(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	resp := rec.Result()
	body, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)

	var result webUtils.Success
	err = json.Unmarshal(body, &result)
	require.NoError(t, err)

	album, err := tools.CreateTrackDataTransferFromInterface(result.Result)
	require.NoError(t, err)
	require.Equal(t, &mockTrack, album)
}

func TestUpdateTrack(t *testing.T) {
	var mockTrack trackProto.Track
	err := faker.FakeData(&mockTrack)
	assert.NoError(t, err)

	e := echo.New()

	dataToSend, err := json.Marshal(&mockTrack)
	req, err := http.NewRequest(echo.PUT, "/tracks/"+strconv.Itoa(int(mockTrack.Id)), bytes.NewBuffer(dataToSend))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	assert.NoError(t, err)

	c := e.NewContext(req, rec)
	c.SetPath("/tracks/")

	mockTrackUseCase := new(mocks2.UseCase)
	mockTrackUseCase.On("Update", &mockTrack).Return(nil)

	handler := Handler{
		TrackUseCase: mockTrackUseCase,
	}

	err = handler.Update(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	resp := rec.Result()
	body, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)

	var result webUtils.Success
	err = json.Unmarshal(body, &result)
	require.NoError(t, err)

	expected := webUtils.Success{
		Status: webUtils.OK,
		Result: constants.SuccessUpdated + "(" + fmt.Sprint(mockTrack.Id) + ")",
	}

	require.Equal(t, expected, result)
}

func TestCreateTrack(t *testing.T) {
	var mockTrack trackProto.Track
	err := faker.FakeData(&mockTrack)
	assert.NoError(t, err)

	e := echo.New()

	dataToSend, err := json.Marshal(&mockTrack)
	req, err := http.NewRequest(echo.POST, "/tracks/"+strconv.Itoa(int(mockTrack.Id)), bytes.NewBuffer(dataToSend))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	assert.NoError(t, err)

	c := e.NewContext(req, rec)
	c.SetPath("/tracks/")

	mockTrackUseCase := new(mocks2.UseCase)
	mockTrackUseCase.On("Create", &mockTrack).Return(nil)
	mockTrackUseCase.On("GetLastId").Return(mockTrack.Id, nil)

	handler := Handler{
		TrackUseCase: mockTrackUseCase,
	}

	err = handler.Create(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	resp := rec.Result()
	body, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)

	var result webUtils.Success
	err = json.Unmarshal(body, &result)
	require.NoError(t, err)

	expected := webUtils.Success{
		Status: webUtils.OK,
		Result: constants.SuccessCreated + "(" + fmt.Sprint(mockTrack.Id) + ")",
	}

	require.Equal(t, expected, result)
}

func TestDeleteTrack(t *testing.T) {
	var mockTrack trackProto.Track
	err := faker.FakeData(&mockTrack)
	assert.NoError(t, err)

	e := echo.New()

	req, err := http.NewRequest(echo.DELETE, "/tracks/"+strconv.Itoa(int(mockTrack.Id)), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/tracks/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(int(mockTrack.Id)))

	mockTrackUseCase := new(mocks2.UseCase)
	mockTrackUseCase.On("Delete", mockTrack.Id).Return(nil)

	handler := Handler{
		TrackUseCase: mockTrackUseCase,
	}

	err = handler.Delete(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	resp := rec.Result()
	body, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)

	var result webUtils.Success
	err = json.Unmarshal(body, &result)
	require.NoError(t, err)

	expected := webUtils.Success{
		Status: webUtils.OK,
		Result: constants.SuccessDeleted + "(" + fmt.Sprint(mockTrack.Id) + ")",
	}

	require.Equal(t, expected, result)
}
