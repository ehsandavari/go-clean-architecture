package controllers

import (
	"github.com/ehsandavari/golang-clean-architecture/application/common/interfaces"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	iLogger interfaces.ILogger
}

func NewController(logger interfaces.ILogger) Controller {
	return Controller{
		iLogger: logger,
	}
}

type BaseApiResponse[TD any] struct {
	err      error
	httpCode int
	Message  string `json:"message"`
	Code     int    `json:"code"`
	Data     TD     `json:"data" extensions:"x-nullable"`
} //@name BaseApiResponse

func NewBaseApiResponse[TD any](error error, message string, code int, data TD, httpCode int) BaseApiResponse[TD] {
	return BaseApiResponse[TD]{err: error, Message: message, Code: code, Data: data, httpCode: httpCode}
}

type AppController[TD any] func(ctx *gin.Context) BaseApiResponse[TD]

func (r AppController[TD]) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		response := r(c)
		if response.err != nil {
			c.JSON(response.httpCode, response)
			return
		}
		c.JSON(response.httpCode, response)
	}
}
