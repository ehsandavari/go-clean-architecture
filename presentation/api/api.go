package api

import "net/http"

type IHandler[TR any] interface {
	ServeHTTP() TR
}

type AppError struct {
	Error   error
	Message string
	Code    int
}

var HttpServers = map[string]*http.Server{}
