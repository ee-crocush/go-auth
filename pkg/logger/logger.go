package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/grpclog"
)

//TODO Пока хз даже, нужен ли тут этот логгер или нет

// Logger предоставляет логгер для приложения
type Logger struct {
	*zap.Logger
}

// New создаёт новый экземпляр логгера
func New(devMode bool) *Logger {
	var zapConfig zap.Config

	// Выбор конфигурации в зависимости от режима
	if devMode {
		zapConfig = zap.NewDevelopmentConfig()
	} else {
		zapConfig = zap.NewProductionConfig()
		zapConfig.EncoderConfig.TimeKey = "timestamp"
		zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	}

	logger, err := zapConfig.Build()
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}

	return &Logger{logger}
}

// Sync выполняет синхронизацию буфера логов
func (l *Logger) Sync() {
	_ = l.Logger.Sync() // Игнорируем ошибку, если она возникнет
}

// ReplaceGrpcLogger устанавливает Zap в качестве логгера для gRPC
func (l *Logger) ReplaceGrpcLogger() {
	grpclog.SetLoggerV2(&zapGrpcLogger{l.Logger})
}

type zapGrpcLogger struct {
	logger *zap.Logger
}

func (l *zapGrpcLogger) Info(args ...interface{}) {
	l.logger.Sugar().Info(args...)
}

func (l *zapGrpcLogger) Infoln(args ...interface{}) {
	l.logger.Sugar().Info(args...)
}

func (l *zapGrpcLogger) Infof(format string, args ...interface{}) {
	l.logger.Sugar().Infof(format, args...)
}

func (l *zapGrpcLogger) Warning(args ...interface{}) {
	l.logger.Sugar().Warn(args...)
}

func (l *zapGrpcLogger) Warningln(args ...interface{}) {
	l.logger.Sugar().Warn(args...)
}

func (l *zapGrpcLogger) Warningf(format string, args ...interface{}) {
	l.logger.Sugar().Warnf(format, args...)
}

func (l *zapGrpcLogger) Error(args ...interface{}) {
	l.logger.Sugar().Error(args...)
}

func (l *zapGrpcLogger) Errorln(args ...interface{}) {
	l.logger.Sugar().Error(args...)
}

func (l *zapGrpcLogger) Errorf(format string, args ...interface{}) {
	l.logger.Sugar().Errorf(format, args...)
}

func (l *zapGrpcLogger) Fatal(args ...interface{}) {
	l.logger.Sugar().Fatal(args...)
}

func (l *zapGrpcLogger) Fatalln(args ...interface{}) {
	l.logger.Sugar().Fatal(args...)
}

func (l *zapGrpcLogger) Fatalf(format string, args ...interface{}) {
	l.logger.Sugar().Fatalf(format, args...)
}

func (l *zapGrpcLogger) V(level int) bool {
	return level <= int(zap.DebugLevel)
}
