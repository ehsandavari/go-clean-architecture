package Interfaces

import (
	"context"
	"time"
)

//go:generate mockgen -destination=../../Mocks/MockInferastructure.go -package=Mocks GolangCodeBase/Application/Common/Interfaces IRedis,ILogger

type (
	IConfig interface {
		Get() error
	}
	IRedis interface {
		Publish(ctx context.Context, channelName string, message interface{}) error
		Subscribe(ctx context.Context, channelName string) <-chan string
	}
	ILogger interface {
		Debug(args ...interface{})
		Debugf(template string, args ...interface{})
		Info(args ...interface{})
		Infof(template string, args ...interface{})
		Warn(args ...interface{})
		Warnf(template string, args ...interface{})
		WarnErrMsg(msg string, err error)
		Error(args ...interface{})
		Errorf(template string, args ...interface{})
		Err(msg string, err error)
		DPanic(args ...interface{})
		DPanicf(template string, args ...interface{})
		Fatal(args ...interface{})
		Fatalf(template string, args ...interface{})
		Printf(template string, args ...interface{})
		Named(name string)
		HttpMiddlewareAccessLogger(method string, uri string, status int, size int64, time time.Duration)
		GrpcMiddlewareAccessLogger(method string, time time.Duration, metaData map[string][]string, err error)
		GrpcMiddlewareAccessLoggerErr(method string, time time.Duration, metaData map[string][]string, err error)
		GrpcClientInterceptorLogger(method string, req interface{}, reply interface{}, time time.Duration, metaData map[string][]string, err error)
		GrpcClientInterceptorLoggerErr(method string, req, reply interface{}, time time.Duration, metaData map[string][]string, err error)
		KafkaProcessMessage(topic string, partition int, message []byte, workerID int, offset int64, time time.Time)
		KafkaLogCommittedMessage(topic string, partition int, offset int64)
		KafkaProcessMessageWithHeaders(topic string, partition int, message []byte, workerID int, offset int64, time time.Time, headers map[string]interface{})
	}
)
