package middlewares

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func RequestId() gin.HandlerFunc {
	return requestid.New()
}
