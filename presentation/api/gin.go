package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var Handler []IHandler[gin.HandlerFunc]

func init() {
	Handler = append(Handler, appHandler(handler_status))
	HttpServers["gin"] = &http.Server{
		Addr:         ":8585",
		Handler:      newGin(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

type appHandler func(ctx *gin.Context) AppError

func (fn appHandler) ServeHTTP() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := fn(c); err != (AppError{}) {
			c.JSON(503, gin.H{"status": err.Error.Error()})
			return
		}
		c.JSON(200, gin.H{"status": "OK"})
	}
}
func newGin() http.Handler {
	g := gin.Default()
	for _, h := range Handler {
		g.GET("/status", h.ServeHTTP())
	}
	return g
}

func handler_status(c *gin.Context) AppError {
	return AppError{
		Error: errors.New("gin"),
	}
}
