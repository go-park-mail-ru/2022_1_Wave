package test

import (
	"encoding/json"
	"errors"
	"github.com/bxcodec/faker"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain/mocks"
	userHttp "github.com/go-park-mail-ru/2022_1_Wave/internal/app/user/delivery/userHttp"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"image"
	"image/color"
	"image/png"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestGetUser(t *testing.T) {
	var mockUser domain.User
	err := faker.FakeData(&mockUser)
	assert.NoError(t, err)
	mockUseCase := new(mocks.UserUseCase)
	mockUseCase.On("GetById", mockUser.ID).Return(&mockUser, nil)
	mockUseCase.On("GetById", mockUser.ID+1).Return(nil, errors.New("error"))

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/users/"+strconv.Itoa(int(mockUser.ID)), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(int(mockUser.ID)))
	handler := userHttp.UserHandler{
		UserUseCase: mockUseCase,
	}

	err = handler.GetUser(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var userResponse userHttp.UserResponse
	rightUserResponse := &userHttp.UserResponse{
		Status: userHttp.StatusOK,
		Result: &mockUser,
	}
	body := rec.Body.String()
	json.Unmarshal([]byte(body), &userResponse)

	assert.Equal(t, userResponse, *rightUserResponse)

	req, err = http.NewRequest(echo.GET, "/users/"+strconv.Itoa(int(mockUser.ID+1)), strings.NewReader(""))
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(int(mockUser.ID + 1)))

	err = handler.GetUser(c)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestGetSelfUser(t *testing.T) {
	var mockUser domain.User
	err := faker.FakeData(&mockUser)
	assert.NoError(t, err)

	mockUseCase := new(mocks.UserUseCase)
	sessionId := "some_session_id"
	mockUseCase.On("GetBySessionId", sessionId).Return(&mockUser, nil)
	mockUseCase.On("GetBySessionId", sessionId+"a").Return(nil, errors.New("error"))

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/users/self", strings.NewReader(""))
	assert.NoError(t, err)

	cookie := &http.Cookie{
		Name:     userHttp.SessionIdKey,
		Value:    sessionId,
		HttpOnly: true,
	}
	req.AddCookie(cookie)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/self")
	handler := userHttp.UserHandler{
		UserUseCase: mockUseCase,
	}

	err = handler.GetSelfUser(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	req, err = http.NewRequest(echo.GET, "/users/self", strings.NewReader(""))
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetPath("/users/self")

	err = handler.GetSelfUser(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, rec.Code)

	req, err = http.NewRequest(echo.GET, "/users/self", strings.NewReader(""))
	cookie.Value += "a"
	req.AddCookie(cookie)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetPath("/users/self")

	err = handler.GetSelfUser(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestUpdateSelfUser(t *testing.T) {
	var mockUser domain.User
	var changeUser domain.User
	err := faker.FakeData(&mockUser)
	assert.NoError(t, err)
	err = faker.FakeData(&changeUser)
	assert.NoError(t, err)

	mockUser2 := mockUser
	mockUser2.ID += 1

	mockUseCase := new(mocks.UserUseCase)
	sessionId := "some_session_id"
	sessionId2 := "some_session_id2"
	mockUseCase.On("GetBySessionId", sessionId).Return(&mockUser, nil)
	mockUseCase.On("GetBySessionId", sessionId+"a").Return(nil, errors.New("error"))
	mockUseCase.On("GetBySessionId", sessionId2).Return(&mockUser2, nil)
	mockUseCase.On("Update", mockUser.ID, &changeUser).Return(nil)
	mockUseCase.On("Update", mockUser.ID+1, &changeUser).Return(errors.New("error"))

	e := echo.New()
	jsonUser, _ := json.Marshal(&changeUser)
	req, err := http.NewRequest(echo.PUT, "/users/self", strings.NewReader(string(jsonUser)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	cookie := &http.Cookie{
		Name:     userHttp.SessionIdKey,
		Value:    sessionId,
		HttpOnly: true,
	}
	req.AddCookie(cookie)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/self")
	handler := userHttp.UserHandler{
		UserUseCase: mockUseCase,
	}

	err = handler.UpdateSelfUser(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	req, err = http.NewRequest(echo.PUT, "/users/self", strings.NewReader(string(jsonUser)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	cookie.Value += "a"
	req.AddCookie(cookie)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetPath("/users/self")

	err = handler.UpdateSelfUser(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, rec.Code)

	req, err = http.NewRequest(echo.PUT, "/users/self", strings.NewReader(string(jsonUser)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	cookie.Name += "a"
	req.AddCookie(cookie)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetPath("/users/self")

	err = handler.UpdateSelfUser(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, rec.Code)

	cookie = &http.Cookie{
		Name:     userHttp.SessionIdKey,
		Value:    sessionId,
		HttpOnly: true,
	}
	req, err = http.NewRequest(echo.PUT, "/users/self", strings.NewReader("aboba"))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	req.AddCookie(cookie)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetPath("/users/self")

	err = handler.UpdateSelfUser(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)

	cookie = &http.Cookie{
		Name:     userHttp.SessionIdKey,
		Value:    sessionId2,
		HttpOnly: true,
	}
	req, err = http.NewRequest(echo.PUT, "/users/self", strings.NewReader(string(jsonUser)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	req.AddCookie(cookie)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetPath("/users/self")

	err = handler.UpdateSelfUser(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func createImage() *image.RGBA {
	width := 100
	height := 100

	upLeft := image.Point{}
	lowRight := image.Point{X: width, Y: height}

	img := image.NewRGBA(image.Rectangle{Min: upLeft, Max: lowRight})

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	cyan := color.RGBA{R: 100, G: 200, B: 200, A: 0xff}

	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			switch {
			case x < width/2 && y < height/2: // upper left quadrant
				img.Set(x, y, cyan)
			case x >= width/2 && y >= height/2: // lower right quadrant
				img.Set(x, y, color.White)
			default:
				// Use zero value.
			}
		}
	}

	// Encode as PNG.
	return img
}

func TestUploadAvatar(t *testing.T) {
	var mockUser domain.User
	err := faker.FakeData(&mockUser)
	assert.NoError(t, err)

	mockUseCase := new(mocks.UserUseCase)
	sessionId := "some_session_id"
	mockUseCase.On("GetBySessionId", sessionId).Return(&mockUser, nil)
	mockUseCase.On("Update", mockUser.ID, mock.Anything).Return(nil)

	pr, pw := io.Pipe()

	multipartWriter := multipart.NewWriter(pw)
	go func() {
		defer multipartWriter.Close()
		img := createImage()
		part, _ := multipartWriter.CreateFormFile("avatar", "someimg.png")

		err = png.Encode(part, img)
	}()

	e := echo.New()

	req, err := http.NewRequest(echo.PUT, "/users/self/upload_avatar", pr)
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, multipartWriter.FormDataContentType())

	cookie := &http.Cookie{
		Name:     userHttp.SessionIdKey,
		Value:    sessionId,
		HttpOnly: true,
	}
	req.AddCookie(cookie)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/self/upload_avatar")
	handler := userHttp.UserHandler{
		UserUseCase: mockUseCase,
	}

	split := strings.Split(userHttp.PathToAvatars, "/")
	now_dir := ""
	for _, s := range split {
		now_dir += s
		err = os.Mkdir(now_dir, 0777)
		now_dir += "/"
	}
	err = handler.UploadAvatar(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
