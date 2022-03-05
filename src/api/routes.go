package api

import (
	"github.com/NNKulickov/wave.music_backend/middleware"
	"github.com/gin-gonic/gin"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server Petstore server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:5000
// @BasePath /api/

func DefineRoutes(router *gin.Engine) {
	// Группа неаутентифицированных запросов. Клиент может выполнять эти запросы
	// не обладая токеном аутентификации.
	//
	// В этой группе только вызов для входа и вызов для регистрации.
	unauthenticated := router.Group("/")
	{
		// Вход в систему по email и паролю.
		unauthenticated.POST("/signin", SignIn)
		// Регистрация в системе.
		unauthenticated.POST("/signup", SignUp)
		// Выйти из системы.
		unauthenticated.POST("/signout", SignOut)
	}

	// Группа аутентифицированных запросов. Клиент должен прислать валидную куки
	// аутентфикации (либо админа либо субъекта) и валидный CSRF токен в
	// заголовке X-CSRF-Token для POST запросов.
	authenticated := router.Group("/")
	authenticated.Use(middleware.Session())
	authenticated.Use(middleware.CSRF())
	// Получить CSRF токен в заголовке X-CSRF-Token.
	authenticated.GET("/csrf", Csrf)

}
