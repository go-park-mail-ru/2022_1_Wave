package api

import (
	"github.com/NNKulickov/wave.music_backend/config"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/csrf"
	"net/http"
)

// GET /api/csrf
//
// Получить CSRF токен. Токен будет находиться заголовке "X-CSRF-Token" в теле
// ответа.
//
// Заголовок X-CSRF-Token должен быть установлен в теле каждого
// аутентифицированного POST запроса. Иначе сервер выдаст статус код 403
// Forbidden.
func Csrf(c *gin.Context) {
	csrfToken := csrf.Token(c.Request)
	c.SetCookie("X-CSRF-Token", csrfToken, 5, "/", config.C.Domain, true, false)
	c.JSON(http.StatusOK, gin.H{"status": "csrf header was set"})
	return
}

// SignIn godoc
// @Summary      SignIn
// @Description  sign in user
// @Tags     auth
// @Accept	 application/json
// @Produce  application/json
// @Param    UserForm body forms.User true  "new jwt data"
// @Success  200 {object} forms.Result
// @Failure 460 {object} forms.Result "Data is invalid"
// @Failure 521 {object} forms.Result "Cannot create session"
// @Router   /api/signin [post]
func SignIn(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"status": "you are signed in"})
	return

}

// SignUp godoc
// @Summary      SignUp
// @Description  sign in user
// @Tags     auth
// @Accept	 application/json
// @Produce  application/json
// @Param    UserForm body forms.User true  "new jwt data"
// @Success  200 {object} forms.Result
// @Failure 460 {object} forms.Result "Data is invalid"
// @Failure 521 {object} forms.Result "Cannot create session"
// @Router   /api/signUp [post]
func SignUp(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"status": "you are signed up"})
	return
}

// SignOut godoc
// @Summary      SignOut
// @Description  sign in user
// @Tags     auth
// @Accept	 application/json
// @Produce  application/json
// @Param    UserForm body forms.User true  "new jwt data"
// @Success  200 {object} forms.Result
// @Failure 460 {object} forms.Result "Data is invalid"
// @Failure 521 {object} forms.Result "Cannot create session"
// @Router   /api/signout [post]
func SignOut(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"status": "you are signed out"})
	return
}
