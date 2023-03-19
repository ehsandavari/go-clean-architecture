package controllers

import (
	"github.com/ehsandavari/golang-clean-architecture/application/common/interfaces"
	"github.com/gin-gonic/gin"
)

type sBaseController struct {
	iLogger interfaces.ILogger
}

func newBaseController(logger interfaces.ILogger) sBaseController {
	return sBaseController{
		iLogger: logger,
	}
}

type baseController[TD any] func(ctx *gin.Context) BaseApiResponse[TD]

func (r baseController[TD]) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		response := r(c)
		if response.err != nil {
			c.JSON(response.httpCode, response)
			return
		}
		c.JSON(response.httpCode, response)
	}
}

type BaseApiResponse[TD any] struct {
	err      error
	httpCode int
	Message  string `json:"message"`
	Code     int64  `json:"code" format:"int64"`
	Data     TD     `json:"data"`
} //@name BaseApiResponse

func newBaseApiResponse[TD any](error error, message string, code int64, data TD, httpCode int) BaseApiResponse[TD] {
	return BaseApiResponse[TD]{
		err:      error,
		Message:  message,
		Code:     code,
		Data:     data,
		httpCode: httpCode,
	}
}
