package controllers

import (
	"github.com/ehsandavari/go-logger"
	"github.com/ehsandavari/go-mediator"
	"github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type sBaseController struct {
	iLogger logger.ILogger
}

func newBaseController(logger logger.ILogger) sBaseController {
	return sBaseController{
		iLogger: logger,
	}
}

type baseController[TReq, TRes any] func(ctx *gin.Context, request TReq) (TRes, error)

func (r baseController[TReq, TRes]) Handle(iLogger logger.ILogger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request TReq
		if bindErr := ctx.ShouldBind(&request); bindErr != nil {
			iLogger.Error("Bind error :", bindErr)
			err := NewApiError(http.StatusBadRequest, "error in validate request")
			err.SetMeta(bindErr.Error())
			if validationErrors, ok := bindErr.(validator.ValidationErrors); ok {
				meta := make(map[string]string, len(validationErrors))
				for _, validationError := range validationErrors {
					meta[validationError.Field()] = validationError.Field() + " is " + validationError.Tag() + " " + validationError.Param()
				}
				err.SetMeta(meta)
			}
			ctx.JSON(http.StatusBadRequest, newBaseApiResponse[ApiError](
				false,
				err,
			))
			return
		}

		result, err := r(ctx, request)
		if err != nil {
			iError := err.(mediator.IError)
			iLogger.Error("handler error :", iError.Code(), iError.Error())
			ctx.JSON(http.StatusInternalServerError, newBaseApiResponse[ApiError](
				false,
				NewApiError(iError.Code(), i18n.MustGetMessage(iError.Error())),
			))
			return
		}

		ctx.JSON(http.StatusOK, newBaseApiResponse[TRes](
			true,
			result,
		))
	}
}

type BaseApiResponse[TD any] struct {
	IsSuccess bool `json:"isSuccess"`
	Data      TD   `json:"data"`
} //@name BaseApiResponse

func newBaseApiResponse[TD any](isSuccess bool, data TD) BaseApiResponse[TD] {
	return BaseApiResponse[TD]{
		IsSuccess: isSuccess,
		Data:      data,
	}
}

type ApiError struct {
	Code    uint   `json:"code" format:"uint32"`
	Message string `json:"message"`
	Meta    any    `json:"meta,omitempty" extensions:"x-nullable,x-omitempty"`
} //@name ApiError

func NewApiError(code uint, message string) ApiError {
	return ApiError{
		Code:    code,
		Message: message,
	}
}

func (r *ApiError) SetMeta(meta any) {
	r.Meta = meta
}
