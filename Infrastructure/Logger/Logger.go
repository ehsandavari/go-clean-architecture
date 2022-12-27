package Logger

import (
	"GolangCodeBase/Application/Common/Interfaces"
	"GolangCodeBase/Domain/Enums"
	"GolangCodeBase/Infrastructure"
	"go.uber.org/fx"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type sLogger struct {
	level       string
	devMode     bool
	encoding    string
	sugarLogger *zap.SugaredLogger
	logger      *zap.Logger
}

func init() {
	Infrastructure.Modules = append(Infrastructure.Modules, fx.Provide(NewLogger))
}

func NewLogger(config SConfig) Interfaces.ILogger {
	logger := &sLogger{
		level:    config.LogLevel,
		devMode:  config.DevMode,
		encoding: config.Encoder,
	}
	logger.config(logger.getLoggerLevel())
	return logger
}

var loggerLevelMap = map[string]zapcore.Level{
	"Debug":  zapcore.DebugLevel,
	"Info":   zapcore.InfoLevel,
	"Warn":   zapcore.WarnLevel,
	"Error":  zapcore.ErrorLevel,
	"DPanic": zapcore.DPanicLevel,
	"Panic":  zapcore.PanicLevel,
	"Fatal":  zapcore.FatalLevel,
}

func (l *sLogger) getLoggerLevel() zapcore.Level {
	level, exist := loggerLevelMap[l.level]
	if !exist {
		return zapcore.DebugLevel
	}
	return level
}

func (l *sLogger) config(logLevel zapcore.Level) {
	logWriter := zapcore.AddSync(os.Stdout)

	var encoderCfg zapcore.EncoderConfig
	if l.devMode {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderCfg = zap.NewProductionEncoderConfig()
	}

	encoderCfg.NameKey = "[SERVICE]"
	encoderCfg.TimeKey = "[TIME]"
	encoderCfg.LevelKey = "[LEVEL]"
	encoderCfg.CallerKey = "[LINE]"
	encoderCfg.MessageKey = "[MESSAGE]"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderCfg.EncodeCaller = zapcore.ShortCallerEncoder
	encoderCfg.EncodeDuration = zapcore.StringDurationEncoder

	var encoder zapcore.Encoder
	if l.encoding == "console" {
		encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoderCfg.EncodeCaller = zapcore.FullCallerEncoder
		encoderCfg.ConsoleSeparator = " | "
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	} else {
		encoderCfg.FunctionKey = "[CALLER]"
		encoderCfg.EncodeName = zapcore.FullNameEncoder
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	}

	core := zapcore.NewCore(encoder, logWriter, zap.NewAtomicLevelAt(logLevel))
	zapLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	l.logger = zapLogger
	l.sugarLogger = zapLogger.Sugar()
}

// Named add Logger microservice name
func (l *sLogger) Named(name string) {
	l.logger = l.logger.Named(name)
	l.sugarLogger = l.sugarLogger.Named(name)
}

// Debug uses fmt.Sprint to construct and log a message.
func (l *sLogger) Debug(args ...interface{}) {
	l.sugarLogger.Debug(args...)
}

// Debugf uses fmt.Sprintf to log a templated message
func (l *sLogger) Debugf(template string, args ...interface{}) {
	l.sugarLogger.Debugf(template, args...)
}

// Info uses fmt.Sprint to construct and log a message
func (l *sLogger) Info(args ...interface{}) {
	l.sugarLogger.Info(args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func (l *sLogger) Infof(template string, args ...interface{}) {
	l.sugarLogger.Infof(template, args...)
}

// Printf uses fmt.Sprintf to log a templated message
func (l *sLogger) Printf(template string, args ...interface{}) {
	l.sugarLogger.Infof(template, args...)
}

// Warn uses fmt.Sprint to construct and log a message.
func (l *sLogger) Warn(args ...interface{}) {
	l.sugarLogger.Warn(args...)
}

// WarnErrMsg log error message with warn level.
func (l *sLogger) WarnErrMsg(msg string, err error) {
	l.logger.Warn(msg, zap.String("error", err.Error()))
}

// Warnf uses fmt.Sprintf to log a templated message.
func (l *sLogger) Warnf(template string, args ...interface{}) {
	l.sugarLogger.Warnf(template, args...)
}

// Error uses fmt.Sprint to construct and log a message.
func (l *sLogger) Error(args ...interface{}) {
	l.sugarLogger.Error(args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func (l *sLogger) Errorf(template string, args ...interface{}) {
	l.sugarLogger.Errorf(template, args...)
}

// Err uses error to log a message.
func (l *sLogger) Err(msg string, err error) {
	l.logger.Error(msg, zap.Error(err))
}

// DPanic uses fmt.Sprint to construct and log a message. In development, the Logger then panics. (See DPanicLevel for details.)
func (l *sLogger) DPanic(args ...interface{}) {
	l.sugarLogger.DPanic(args...)
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the Logger then panics. (See DPanicLevel for details.)
func (l *sLogger) DPanicf(template string, args ...interface{}) {
	l.sugarLogger.DPanicf(template, args...)
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func (l *sLogger) Panic(args ...interface{}) {
	l.sugarLogger.Panic(args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics
func (l *sLogger) Panicf(template string, args ...interface{}) {
	l.sugarLogger.Panicf(template, args...)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func (l *sLogger) Fatal(args ...interface{}) {
	l.sugarLogger.Fatal(args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func (l *sLogger) Fatalf(template string, args ...interface{}) {
	l.sugarLogger.Fatalf(template, args...)
}

func (l *sLogger) HttpMiddlewareAccessLogger(method, uri string, status int, size int64, time time.Duration) {
	l.logger.Info(
		Enums.HTTP,
		zap.String(Enums.METHOD, method),
		zap.String(Enums.URI, uri),
		zap.Int(Enums.STATUS, status),
		zap.Int64(Enums.SIZE, size),
		zap.Duration(Enums.TIME, time),
	)
}

func (l *sLogger) GrpcMiddlewareAccessLogger(method string, time time.Duration, metaData map[string][]string, err error) {
	l.logger.Info(
		Enums.GRPC,
		zap.String(Enums.METHOD, method),
		zap.Duration(Enums.TIME, time),
		zap.Any(Enums.METADATA, metaData),
		zap.Any(Enums.ERROR, err),
	)
}

func (l *sLogger) GrpcMiddlewareAccessLoggerErr(method string, time time.Duration, metaData map[string][]string, err error) {
	l.logger.Error(
		Enums.GRPC,
		zap.String(Enums.METHOD, method),
		zap.Duration(Enums.TIME, time),
		zap.Any(Enums.METADATA, metaData),
		zap.Any(Enums.ERROR, err),
	)
}

func (l *sLogger) GrpcClientInterceptorLogger(method string, req, reply interface{}, time time.Duration, metaData map[string][]string, err error) {
	l.logger.Info(
		Enums.GRPC,
		zap.String(Enums.METHOD, method),
		zap.Any(Enums.REQUEST, req),
		zap.Any(Enums.REPLY, reply),
		zap.Duration(Enums.TIME, time),
		zap.Any(Enums.METADATA, metaData),
		zap.Any(Enums.ERROR, err),
	)
}

func (l *sLogger) GrpcClientInterceptorLoggerErr(method string, req, reply interface{}, time time.Duration, metaData map[string][]string, err error) {
	l.logger.Error(
		Enums.GRPC,
		zap.String(Enums.METHOD, method),
		zap.Any(Enums.REQUEST, req),
		zap.Any(Enums.REPLY, reply),
		zap.Duration(Enums.TIME, time),
		zap.Any(Enums.METADATA, metaData),
		zap.Any(Enums.ERROR, err),
	)
}

func (l *sLogger) KafkaProcessMessage(topic string, partition int, message []byte, workerID int, offset int64, time time.Time) {
	l.logger.Debug(
		"(Processing Kafka message)",
		zap.String(Enums.Topic, topic),
		zap.Int(Enums.Partition, partition),
		zap.Int(Enums.MessageSize, len(message)),
		zap.Int(Enums.WorkerID, workerID),
		zap.Int64(Enums.Offset, offset),
		zap.Time(Enums.Time, time),
	)
}

func (l *sLogger) KafkaLogCommittedMessage(topic string, partition int, offset int64) {
	l.logger.Debug(
		"(Committed Kafka message)",
		zap.String(Enums.Topic, topic),
		zap.Int(Enums.Partition, partition),
		zap.Int64(Enums.Offset, offset),
	)
}

func (l *sLogger) KafkaProcessMessageWithHeaders(topic string, partition int, message []byte, workerID int, offset int64, time time.Time, headers map[string]interface{}) {
	l.logger.Debug(
		"(Processing Kafka message)",
		zap.String(Enums.Topic, topic),
		zap.Int(Enums.Partition, partition),
		zap.Int(Enums.MessageSize, len(message)),
		zap.Int(Enums.WorkerID, workerID),
		zap.Int64(Enums.Offset, offset),
		zap.Time(Enums.Time, time),
		zap.Any(Enums.KafkaHeaders, headers),
	)
}
