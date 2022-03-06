package api

import (
	"github.com/NNKulickov/wave.music_backend/config"
	"github.com/NNKulickov/wave.music_backend/db"
	"github.com/NNKulickov/wave.music_backend/forms"
	"github.com/NNKulickov/wave.music_backend/service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/csrf"
	"net/http"
	"time"
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

// Login godoc
// @Summary      Login
// @Description  login user
// @Tags     auth
// @Accept	 application/json
// @Produce  application/json
// @Param    UserForm body forms.User true  "new jwt data"
// @Success  200 {object} forms.Result
// @Failure 460 {object} forms.Result "Data is invalid"
// @Failure 521 {object} forms.Result "Cannot create session"
// @Router   /api/login [post]
func Login(w http.ResponseWriter, r *http.Request) {
	_, err := service.GetSession(r)
	if err == nil {
		http.Error(w, `{"error": "already authorized"}`, http.StatusForbidden)
		return
	}

	userToLogin, err := forms.UserUnmarshal(r)
	if err != nil {
		http.Error(w, `{"error": "bad request"}`, http.StatusBadRequest)
		return
	}

	// проверяем логин с паролем
	var checkUser bool
	if userToLogin.Username != "" {
		checkUser = db.MyUserStorage.CheckUsernameAndPassword(userToLogin.Username, userToLogin.Password)
	} else {
		checkUser = db.MyUserStorage.CheckEmailAndPassword(userToLogin.Email, userToLogin.Password)
	}

	if !checkUser {
		http.Error(w, `{"error": "invalid login or password"}`, http.StatusBadRequest)
		return
	}

	// добавляем новую сессию
	var userCookie string
	if userToLogin.Username != "" {
		user, _ := db.MyUserStorage.SelectByUsername(userToLogin.Username)
		userCookie = service.SetNewSession(user.ID)
	} else {
		user, _ := db.MyUserStorage.SelectByEmail(userToLogin.Email)
		userCookie = service.SetNewSession(user.ID)
	}

	cookie := &http.Cookie{
		Name:    config.C.SessionIDKey,
		Value:   userCookie,
		Expires: time.Now().Add(10 * time.Hour),
	}

	http.SetCookie(w, cookie)
	w.Write([]byte(`{"status": "you are login"}`))
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
func SignUp(w http.ResponseWriter, r *http.Request) {
	_, err := service.GetSession(r)
	if err == nil {
		http.Error(w, `{"error": "already authorized"}`, http.StatusForbidden)
		return
	}

	userToLogin, err := forms.UserUnmarshal(r)
	if err != nil {
		http.Error(w, `{"error": "bad request"}`, http.StatusBadRequest)
		return
	}

	err = db.MyUserStorage.Insert(&db.User{
		Username: userToLogin.Username,
		Email:    userToLogin.Email,
		Password: userToLogin.Password,
	})

	if err != nil {
		http.Error(w, `{"error": "user already exist"}`, http.StatusBadRequest)
		return
	}

	// теперь создаем для зарегистрированного пользователя сессию
	nowUser, err := db.MyUserStorage.SelectByUsername(userToLogin.Username)
	sessionId := service.SetNewSession(nowUser.ID)

	cookie := &http.Cookie{
		Name:    config.C.SessionIDKey,
		Value:   sessionId,
		Expires: time.Now().Add(10 * time.Hour),
	}

	http.SetCookie(w, cookie)
	w.Write([]byte(`{"status": "you are sign up"}`))

	return
}

// Logout godoc
// @Summary      Logout
// @Description  sign in user
// @Tags     auth
// @Accept	 application/json
// @Produce  application/json
// @Param    UserForm body forms.User true  "new jwt data"
// @Success  200 {object} forms.Result
// @Failure 460 {object} forms.Result "Data is invalid"
// @Failure 521 {object} forms.Result "Cannot create session"
// @Router   /api/signout [post]
func Logout(w http.ResponseWriter, r *http.Request) {
	_, err := service.GetSession(r)
	if err != nil {
		http.Error(w, `{"error": "no session"}`, 401)
		return
	}

	session, _ := r.Cookie(config.C.SessionIDKey)
	service.DeleteSession(session.Value)

	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
}
