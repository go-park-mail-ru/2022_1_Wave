package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Проверить есть ли у клиента валидная сессия (токен сессии в куки).
func Session() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.JSON(http.StatusForbidden, gin.H{"error": "you are not logged in"})
	}
}
