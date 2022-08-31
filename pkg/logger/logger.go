package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

func getLogWriter(filename string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    5,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   true,
	}
	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func New(enable bool, filename, level string) *zap.SugaredLogger {
	if !enable {
		return zap.NewNop().Sugar()
	}
	writeSyncer := getLogWriter(filename)
	encoder := getEncoder()

	l, err := zapcore.ParseLevel(level)
	if err != nil {
		l = zapcore.InfoLevel
	}
	core := zapcore.NewCore(encoder, writeSyncer, l)
	return zap.New(core, zap.AddCaller()).Sugar()
}
