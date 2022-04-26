package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bxcodec/faker"
	constants "github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/gateway/gatewayProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/track/trackProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/test/mocks"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/tools"
	trackDeliveryHttp "github.com/go-park-mail-ru/2022_1_Wave/internal/track/delivery/http"
	"github.com/go-park-mail-ru/2022_1_Wave/pkg/webUtils"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestGetAllTracks(t *testing.T) {
	var mockTracks = &trackProto.TracksResponse{Tracks: make([]*trackProto.TrackDataTransfer, 10)}

	err := faker.FakeData(&mockTracks)
	assert.NoError(t, err)

	e := echo.New()

	req, err := http.NewRequest(echo.GET, "/tracks/", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/tracks/")

	mockTrackUseCase := new(mocks.TrackAgent)

	mockTrackUseCase.On("GetAll").Return(mockTracks, nil)

	handler := trackDeliveryHttp.Handler{
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

	resultMap := result.Result.(map[string]interface{})
	tracks := resultMap["tracks"]

	for idx, obj := range tracks.([]interface{}) {
		track, err := tools.CreateTrackDataTransferFromInterface(obj)
		require.NoError(t, err)
		require.Equal(t, mockTracks.Tracks[idx], track)
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

	mockTrackUseCase := new(mocks.TrackAgent)

	mockTrackUseCase.On("GetById", &gatewayProto.IdArg{Id: mockTrack.Id}).Return(&mockTrack, nil)

	handler := trackDeliveryHttp.Handler{
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

	mockTrackUseCase := new(mocks.TrackAgent)
	mockTrackUseCase.On("Update", &mockTrack).Return(nil)

	handler := trackDeliveryHttp.Handler{
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

	mockTrackUseCase := new(mocks.TrackAgent)
	mockTrackUseCase.On("Create", &mockTrack).Return(nil)
	mockTrackUseCase.On("GetLastId").Return(&gatewayProto.IntResponse{Data: mockTrack.Id}, nil)

	handler := trackDeliveryHttp.Handler{
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
		Result: constants.SuccessCreated + "(Data:" + fmt.Sprint(mockTrack.Id) + ")",
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

	mockTrackUseCase := new(mocks.TrackAgent)
	mockTrackUseCase.On("Delete", &gatewayProto.IdArg{Id: mockTrack.Id}).Return(nil)

	handler := trackDeliveryHttp.Handler{
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

//func TestGetSelfUser(t *testing.T) {
//	var mockUser domain.User
//	err := faker.FakeData(&mockUser)
//	assert.NoError(t, err)
//
//	mockUseCase := new(mocks.UserUseCase)
//	sessionId := "some_session_id"
//	mockUseCase.On("GetBySessionId", sessionId).Return(&mockUser, nil)
//	mockUseCase.On("GetBySessionId", sessionId+"a").Return(nil, errors.New("error"))
//
//	e := echo.New()
//	req, err := http.NewRequest(echo.GET, "/users/self", strings.NewReader(""))
//	assert.NoError(t, err)
//
//	cookie := &http.Cookie{
//		Name:     userHttp.SessionIdKey,
//		Value:    sessionId,
//		HttpOnly: true,
//	}
//	req.AddCookie(cookie)
//	rec := httptest.NewRecorder()
//	c := e.NewContext(req, rec)
//	c.SetPath("/users/self")
//	handler := userHttp.UserHandler{
//		UserUseCase: mockUseCase,
//	}
//
//	err = handler.GetSelfUser(c)
//	assert.NoError(t, err)
//	assert.Equal(t, http.StatusOK, rec.Code)
//
//	req, err = http.NewRequest(echo.GET, "/users/self", strings.NewReader(""))
//	rec = httptest.NewRecorder()
//	c = e.NewContext(req, rec)
//	c.SetPath("/users/self")
//
//	err = handler.GetSelfUser(c)
//	assert.NoError(t, err)
//	assert.Equal(t, http.StatusUnauthorized, rec.Code)
//
//	req, err = http.NewRequest(echo.GET, "/users/self", strings.NewReader(""))
//	cookie.Value += "a"
//	req.AddCookie(cookie)
//	rec = httptest.NewRecorder()
//	c = e.NewContext(req, rec)
//	c.SetPath("/users/self")
//
//	err = handler.GetSelfUser(c)
//	assert.NoError(t, err)
//	assert.Equal(t, http.StatusUnauthorized, rec.Code)
//}
//
//func TestUpdateSelfUser(t *testing.T) {
//	var mockUser domain.User
//	var changeUser domain.User
//	err := faker.FakeData(&mockUser)
//	assert.NoError(t, err)
//	err = faker.FakeData(&changeUser)
//	assert.NoError(t, err)
//
//	mockUser2 := mockUser
//	mockUser2.ID += 1
//
//	mockUseCase := new(mocks.UserUseCase)
//	sessionId := "some_session_id"
//	sessionId2 := "some_session_id2"
//	mockUseCase.On("GetBySessionId", sessionId).Return(&mockUser, nil)
//	mockUseCase.On("GetBySessionId", sessionId+"a").Return(nil, errors.New("error"))
//	mockUseCase.On("GetBySessionId", sessionId2).Return(&mockUser2, nil)
//	mockUseCase.On("Update", mockUser.ID, &changeUser).Return(nil)
//	mockUseCase.On("Update", mockUser.ID+1, &changeUser).Return(errors.New("error"))
//
//	e := echo.New()
//	jsonUser, _ := json.Marshal(&changeUser)
//	req, err := http.NewRequest(echo.PUT, "/users/self", strings.NewReader(string(jsonUser)))
//	assert.NoError(t, err)
//	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
//
//	cookie := &http.Cookie{
//		Name:     userHttp.SessionIdKey,
//		Value:    sessionId,
//		HttpOnly: true,
//	}
//	req.AddCookie(cookie)
//	rec := httptest.NewRecorder()
//	c := e.NewContext(req, rec)
//	c.SetPath("/users/self")
//	handler := userHttp.UserHandler{
//		UserUseCase: mockUseCase,
//	}
//
//	err = handler.UpdateSelfUser(c)
//	assert.NoError(t, err)
//	assert.Equal(t, http.StatusOK, rec.Code)
//
//	req, err = http.NewRequest(echo.PUT, "/users/self", strings.NewReader(string(jsonUser)))
//	assert.NoError(t, err)
//	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
//
//	cookie.Value += "a"
//	req.AddCookie(cookie)
//	rec = httptest.NewRecorder()
//	c = e.NewContext(req, rec)
//	c.SetPath("/users/self")
//
//	err = handler.UpdateSelfUser(c)
//	assert.NoError(t, err)
//	assert.Equal(t, http.StatusUnauthorized, rec.Code)
//
//	req, err = http.NewRequest(echo.PUT, "/users/self", strings.NewReader(string(jsonUser)))
//	assert.NoError(t, err)
//	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
//
//	cookie.Name += "a"
//	req.AddCookie(cookie)
//	rec = httptest.NewRecorder()
//	c = e.NewContext(req, rec)
//	c.SetPath("/users/self")
//
//	err = handler.UpdateSelfUser(c)
//	assert.NoError(t, err)
//	assert.Equal(t, http.StatusUnauthorized, rec.Code)
//
//	cookie = &http.Cookie{
//		Name:     userHttp.SessionIdKey,
//		Value:    sessionId,
//		HttpOnly: true,
//	}
//	req, err = http.NewRequest(echo.PUT, "/users/self", strings.NewReader("aboba"))
//	assert.NoError(t, err)
//	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
//
//	req.AddCookie(cookie)
//	rec = httptest.NewRecorder()
//	c = e.NewContext(req, rec)
//	c.SetPath("/users/self")
//
//	err = handler.UpdateSelfUser(c)
//	assert.NoError(t, err)
//	assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
//
//	cookie = &http.Cookie{
//		Name:     userHttp.SessionIdKey,
//		Value:    sessionId2,
//		HttpOnly: true,
//	}
//	req, err = http.NewRequest(echo.PUT, "/users/self", strings.NewReader(string(jsonUser)))
//	assert.NoError(t, err)
//	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
//
//	req.AddCookie(cookie)
//	rec = httptest.NewRecorder()
//	c = e.NewContext(req, rec)
//	c.SetPath("/users/self")
//
//	err = handler.UpdateSelfUser(c)
//	assert.NoError(t, err)
//	assert.Equal(t, http.StatusBadRequest, rec.Code)
//}
//
//func createImage() *image.RGBA {
//	width := 100
//	height := 100
//
//	upLeft := image.Point{}
//	lowRight := image.Point{X: width, Y: height}
//
//	img := image.NewRGBA(image.Rectangle{Min: upLeft, Max: lowRight})
//
//	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
//	cyan := color.RGBA{R: 100, G: 200, B: 200, A: 0xff}
//
//	// Set color for each pixel.
//	for x := 0; x < width; x++ {
//		for y := 0; y < height; y++ {
//			switch {
//			case x < width/2 && y < height/2: // upper left quadrant
//				img.Set(x, y, cyan)
//			case x >= width/2 && y >= height/2: // lower right quadrant
//				img.Set(x, y, color.White)
//			default:
//				// Use zero value.
//			}
//		}
//	}
//
//	// Encode as PNG.
//	return img
//}
//
//func TestUploadAvatar(t *testing.T) {
//	var mockUser domain.User
//	err := faker.FakeData(&mockUser)
//	assert.NoError(t, err)
//
//	mockUseCase := new(mocks.UserUseCase)
//	sessionId := "some_session_id"
//	mockUseCase.On("GetBySessionId", sessionId).Return(&mockUser, nil)
//	mockUseCase.On("Update", mockUser.ID, mock.Anything).Return(nil)
//
//	pr, pw := io.Pipe()
//
//	multipartWriter := multipart.NewWriter(pw)
//	go func() {
//		defer multipartWriter.Close()
//		img := createImage()
//		part, _ := multipartWriter.CreateFormFile("avatar", "someimg.png")
//
//		err = png.Encode(part, img)
//	}()
//
//	e := echo.New()
//
//	req, err := http.NewRequest(echo.PUT, "/users/self/upload_avatar", pr)
//	assert.NoError(t, err)
//	req.Header.Set(echo.HeaderContentType, multipartWriter.FormDataContentType())
//
//	cookie := &http.Cookie{
//		Name:     userHttp.SessionIdKey,
//		Value:    sessionId,
//		HttpOnly: true,
//	}
//	req.AddCookie(cookie)
//	rec := httptest.NewRecorder()
//	c := e.NewContext(req, rec)
//	c.SetPath("/users/self/upload_avatar")
//	handler := userHttp.UserHandler{
//		UserUseCase: mockUseCase,
//	}
//
//	split := strings.Split(userHttp.PathToAvatars, "/")
//	now_dir := ""
//	for _, s := range split {
//		now_dir += s
//		err = os.Mkdir(now_dir, 0777)
//		now_dir += "/"
//	}
//	err = handler.UploadAvatar(c)
//	assert.NoError(t, err)
//	assert.Equal(t, http.StatusBadRequest, rec.Code)
//}
