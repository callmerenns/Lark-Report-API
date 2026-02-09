package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Loggers struct {
	App      *zap.Logger
	Access   *zap.Logger
	Error    *zap.Logger
	Security *zap.Logger
}

var Log Loggers

func newCore(filename string, level zapcore.Level) zapcore.Core {
	writer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filename,
		MaxSize:    20, // MB
		MaxBackups: 10,
		MaxAge:     14, // days
		Compress:   true,
	})

	encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:      "timestamp",
		LevelKey:     "level",
		MessageKey:   "message",
		CallerKey:    "caller",
		EncodeTime:   zapcore.ISO8601TimeEncoder,
		EncodeLevel:  zapcore.LowercaseLevelEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	})

	return zapcore.NewCore(encoder, writer, level)
}

func InitLogger() {
	Log = Loggers{
		App: zap.New(
			newCore("logs/app.log", zap.InfoLevel),
			zap.AddCaller(),
		),
		Access: zap.New(
			newCore("logs/access.log", zap.InfoLevel),
		),
		Error: zap.New(
			newCore("logs/error.log", zap.ErrorLevel),
			zap.AddCaller(),
			zap.AddStacktrace(zap.ErrorLevel),
		),
		Security: zap.New(
			newCore("logs/security.log", zap.WarnLevel),
			zap.AddCaller(),
		),
	}
}
