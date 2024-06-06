package logger

import (
	"go.uber.org/zap"
)

// 固定类型，速度更快,打印日志尽量用这种
func Debug(template string, fields ...zap.Field) {
	zapLog.Debug(template, fields...)
}

func Info(template string, fields ...zap.Field) {
	zapLog.Info(template, fields...)
}

func Warn(template string, fields ...zap.Field) {
	zapLog.Warn(template, fields...)
}

func Error(template string, fields ...zap.Field) {
	zapLog.Error(template, fields...)
}

func Panic(template string, fields ...zap.Field) {
	zapLog.Panic(template, fields...)
}

func Fatal(template string, fields ...zap.Field) {
	zapLog.Fatal(template, fields...)
}

// 可变参数，速度慢
func Debugf(template string, args ...interface{}) {
	zapLogSugar.Debugf(template, args...)
}

func Infof(template string, args ...interface{}) {
	zapLogSugar.Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	zapLogSugar.Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	zapLogSugar.Errorf(template, args...)
}

func Panicf(template string, args ...interface{}) {
	zapLogSugar.Panicf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	zapLogSugar.Fatalf(template, args...)
}
