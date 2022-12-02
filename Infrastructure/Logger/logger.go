package Logger

import (
	"GolangCodeBase/Application/Common/Interfaces"
	"GolangCodeBase/Domain/Enums"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogConfig struct {
	LogLevel string `mapstructure:"level"`
	DevMode  bool   `mapstructure:"devMode"`
	Encoder  string `mapstructure:"encoder"`
}

// Application logger
type logger struct {
	level       string
	devMode     bool
	encoding    string
	sugarLogger *zap.SugaredLogger
	logger      *zap.Logger
}

func NewLogger(cfg LogConfig) Interfaces.ILogger {
	return &logger{level: cfg.LogLevel, devMode: cfg.DevMode, encoding: cfg.Encoder}
}

// For mapping Config Logger to email_service Logger levels
var loggerLevelMap = map[string]zapcore.Level{
	"Debug":  zapcore.DebugLevel,
	"Info":   zapcore.InfoLevel,
	"Warn":   zapcore.WarnLevel,
	"Error":  zapcore.ErrorLevel,
	"DPanic": zapcore.DPanicLevel,
	"Panic":  zapcore.PanicLevel,
	"Fatal":  zapcore.FatalLevel,
}

func (l *logger) getLoggerLevel() zapcore.Level {
	level, exist := loggerLevelMap[l.level]
	if !exist {
		return zapcore.DebugLevel
	}

	return level
}

func (l *logger) setLoggerLevel(logLevel string) zapcore.Level {
	level, exist := loggerLevelMap[logLevel]
	if !exist {
		return zapcore.DebugLevel
	}
	return level
}

func (l *logger) config(logLevel zapcore.Level) {
	logWriter := zapcore.AddSync(os.Stdout)

	var encoderCfg zapcore.EncoderConfig
	if l.devMode {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderCfg = zap.NewProductionEncoderConfig()
	}

	var encoder zapcore.Encoder
	encoderCfg.NameKey = "[SERVICE]"
	encoderCfg.TimeKey = "[TIME]"
	encoderCfg.LevelKey = "[LEVEL]"
	encoderCfg.CallerKey = "[LINE]"
	encoderCfg.MessageKey = "[MESSAGE]"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderCfg.EncodeCaller = zapcore.ShortCallerEncoder
	encoderCfg.EncodeDuration = zapcore.StringDurationEncoder

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

// InitLogger Init Logger
func (l *logger) InitLogger() {
	l.config(l.getLoggerLevel())
}

func (l *logger) SetLogLevel(logLevel string) {
	logLvl := l.setLoggerLevel(logLevel)
	l.config(logLvl)
	l.logger.Info("(SET LOG LEVEL)", zap.String("LEVEL", logLvl.CapitalString()))
	l.Sync()
}

// Named add Logger microservice name
func (l *logger) Named(name string) {
	l.logger = l.logger.Named(name)
	l.sugarLogger = l.sugarLogger.Named(name)
}

// Debug uses fmt.Sprint to construct and log a message.
func (l *logger) Debug(args ...interface{}) {
	l.sugarLogger.Debug(args...)
}

// Debugf uses fmt.Sprintf to log a templated message
func (l *logger) Debugf(template string, args ...interface{}) {
	l.sugarLogger.Debugf(template, args...)
}

// Info uses fmt.Sprint to construct and log a message
func (l *logger) Info(args ...interface{}) {
	l.sugarLogger.Info(args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func (l *logger) Infof(template string, args ...interface{}) {
	l.sugarLogger.Infof(template, args...)
}

// Printf uses fmt.Sprintf to log a templated message
func (l *logger) Printf(template string, args ...interface{}) {
	l.sugarLogger.Infof(template, args...)
}

// Warn uses fmt.Sprint to construct and log a message.
func (l *logger) Warn(args ...interface{}) {
	l.sugarLogger.Warn(args...)
}

// WarnErrMsg log error message with warn level.
func (l *logger) WarnErrMsg(msg string, err error) {
	l.logger.Warn(msg, zap.String("error", err.Error()))
}

// Warnf uses fmt.Sprintf to log a templated message.
func (l *logger) Warnf(template string, args ...interface{}) {
	l.sugarLogger.Warnf(template, args...)
}

// Error uses fmt.Sprint to construct and log a message.
func (l *logger) Error(args ...interface{}) {
	l.sugarLogger.Error(args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func (l *logger) Errorf(template string, args ...interface{}) {
	l.sugarLogger.Errorf(template, args...)
}

// Err uses error to log a message.
func (l *logger) Err(msg string, err error) {
	l.logger.Error(msg, zap.Error(err))
}

// DPanic uses fmt.Sprint to construct and log a message. In development, the Logger then panics. (See DPanicLevel for details.)
func (l *logger) DPanic(args ...interface{}) {
	l.sugarLogger.DPanic(args...)
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the Logger then panics. (See DPanicLevel for details.)
func (l *logger) DPanicf(template string, args ...interface{}) {
	l.sugarLogger.DPanicf(template, args...)
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func (l *logger) Panic(args ...interface{}) {
	l.sugarLogger.Panic(args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics
func (l *logger) Panicf(template string, args ...interface{}) {
	l.sugarLogger.Panicf(template, args...)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func (l *logger) Fatal(args ...interface{}) {
	l.sugarLogger.Fatal(args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func (l *logger) Fatalf(template string, args ...interface{}) {
	l.sugarLogger.Fatalf(template, args...)
}

// Sync flushes any buffered log entries
func (l *logger) Sync() error {
	go l.logger.Sync()
	return l.sugarLogger.Sync()
}

func (l *logger) HttpMiddlewareAccessLogger(method, uri string, status int, size int64, time time.Duration) {
	l.logger.Info(
		Enums.HTTP,
		zap.String(Enums.METHOD, method),
		zap.String(Enums.URI, uri),
		zap.Int(Enums.STATUS, status),
		zap.Int64(Enums.SIZE, size),
		zap.Duration(Enums.TIME, time),
	)
}

func (l *logger) GrpcMiddlewareAccessLogger(method string, time time.Duration, metaData map[string][]string, err error) {
	l.logger.Info(
		Enums.GRPC,
		zap.String(Enums.METHOD, method),
		zap.Duration(Enums.TIME, time),
		zap.Any(Enums.METADATA, metaData),
		zap.Any(Enums.ERROR, err),
	)
}

func (l *logger) GrpcMiddlewareAccessLoggerErr(method string, time time.Duration, metaData map[string][]string, err error) {
	l.logger.Error(
		Enums.GRPC,
		zap.String(Enums.METHOD, method),
		zap.Duration(Enums.TIME, time),
		zap.Any(Enums.METADATA, metaData),
		zap.Any(Enums.ERROR, err),
	)
}

func (l *logger) GrpcClientInterceptorLogger(method string, req, reply interface{}, time time.Duration, metaData map[string][]string, err error) {
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

func (l *logger) GrpcClientInterceptorLoggerErr(method string, req, reply interface{}, time time.Duration, metaData map[string][]string, err error) {
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

func (l *logger) KafkaProcessMessage(topic string, partition int, message []byte, workerID int, offset int64, time time.Time) {
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

func (l *logger) KafkaLogCommittedMessage(topic string, partition int, offset int64) {
	l.logger.Debug(
		"(Committed Kafka message)",
		zap.String(Enums.Topic, topic),
		zap.Int(Enums.Partition, partition),
		zap.Int64(Enums.Offset, offset),
	)
}

func (l *logger) KafkaProcessMessageWithHeaders(topic string, partition int, message []byte, workerID int, offset int64, time time.Time, headers map[string]interface{}) {
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
