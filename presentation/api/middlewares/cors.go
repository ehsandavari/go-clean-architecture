package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return cors.Default()
}
