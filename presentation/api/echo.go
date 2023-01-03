package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"time"
)

var HandlerEcho []IHandler[echo.HandlerFunc]

func init() {
	HandlerEcho = append(HandlerEcho, appHandlerEcho(handler_status_echo))
	HttpServers["echo"] = &http.Server{
		Addr:         ":8181",
		Handler:      newEcho(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

type appHandlerEcho func(ctx echo.Context) AppError

func (fn appHandlerEcho) ServeHTTP() echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := fn(c); err != (AppError{}) {
			return c.JSON(503, gin.H{"status": err.Error.Error()})
		}
		return c.JSON(200, gin.H{"status": "OK"})
	}
}
func newEcho() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	for _, h := range HandlerEcho {
		e.GET("/status", h.ServeHTTP())
	}
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		code := http.StatusInternalServerError
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
		}
		c.Logger().Error(err, code)
	}

	return e
}

func handler_status_echo(c echo.Context) AppError {
	return AppError{
		Error: errors.New("echo"),
	}
}
