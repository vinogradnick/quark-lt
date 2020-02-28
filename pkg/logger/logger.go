package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger(name string) {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   " ~/go/src/gitlab.com/quark_worker" + name + ".log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
	})
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		w,
		zap.InfoLevel,
	)
	logger := zap.New(core)
	if logger != nil {

	}
	zap.L().Info("flex")
}
