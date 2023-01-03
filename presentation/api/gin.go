package api

import (
	"github.com/ehsandavari/go-mediator"
	"github.com/ehsandavari/golang-clean-architecture/application/handlers/order/commands/publishOrder"
	"github.com/ehsandavari/golang-clean-architecture/presentation/controller/dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var Handler []IHandler[gin.HandlerFunc]

func init() {
	Handler = append(Handler, appHandler(order))
	HttpServers["gin"] = &http.Server{
		Addr:         ":9090",
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
		g.POST("/order", h.ServeHTTP())
	}
	return g
}

func order(ctx *gin.Context) AppError {
	params := &dto.CreateOrderRequest{}
	err := ctx.ShouldBindJSON(&params)
	if err != nil {
		return AppError{
			Message: err.Error(),
			Error:   err,
		}
	}
	_, err = mediator.Send[publishOrder.SPublishOrderCommand, string](ctx, publishOrder.NewSPublishOrderCommand(params.Price, params.Title))
	if err != nil {
		return AppError{
			Message: err.Error(),
			Error:   err,
		}
	}
	return AppError{}
}
