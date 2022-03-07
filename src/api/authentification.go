package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2022_1_Wave/config"
	"github.com/go-park-mail-ru/2022_1_Wave/db"
	"github.com/go-park-mail-ru/2022_1_Wave/forms"
	"github.com/go-park-mail-ru/2022_1_Wave/service"
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

	// зайдя сюда мы уверены, что пользователь прислал нам куку неавторизованного пользователя (навешены специальные middleware)
	var user *db.User
	if userToLogin.Username != "" {
		user, _ = db.MyUserStorage.SelectByUsername(userToLogin.Username)
	} else {
		user, _ = db.MyUserStorage.SelectByEmail(userToLogin.Email)
	}

	// и мы просто обновляем состояние текущей сессии
	session, err := service.GetSession(r)
	session.IsAuthorized = true
	session.UserId = user.ID

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
	userToLogin, err := forms.UserUnmarshal(r)
	if err != nil {
		http.Error(w, `{"error": "bad request"}`, http.StatusBadRequest)
		return
	}

	if !userToLogin.IsValid() {
		http.Error(w, `{"error": "invalid fields"}`, http.StatusBadRequest)
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
	session, err := service.GetSession(r)
	session.UserId = nowUser.ID
	session.IsAuthorized = true

	w.Write([]byte(`{"status": "you are sign up"}`))
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
	session, _ := r.Cookie(config.C.SessionIDKey)
	service.DeleteSession(session.Value)

	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
}

func GetCSRF(w http.ResponseWriter, r *http.Request) {
	if !service.IsAuthorized(r) {
		cookie, csrfToken := service.SetNewUnauthorizedSession()
		http.SetCookie(w, cookie)
		w.Header().Set("X-CSRF-TOKEN", csrfToken)
	} else {
		sessionID, _ := r.Cookie(config.C.SessionIDKey)
		csrfToken := service.SetNewCSRFToken(sessionID.Value)
		w.Header().Set("X-CSRF-TOKEN", csrfToken)
	}
}
