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

type BaseApiResponse struct {
	Error    error  `json:"error"`
	Message  string `json:"message"`
	Code     int    `json:"code"`
	httpCode int
} //@name BaseApiResponse

func NewBaseApiResponse(error error, message string, code int, httpCode int) BaseApiResponse {
	return BaseApiResponse{Error: error, Message: message, Code: code, httpCode: httpCode}
}

type AppController func(ctx *gin.Context) BaseApiResponse

func (r AppController) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		response := r(c)
		if response.Error != nil {
			c.JSON(response.httpCode, response)
			return
		}
		c.JSON(response.httpCode, response)
	}
}
