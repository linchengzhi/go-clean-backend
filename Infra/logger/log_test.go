package logger

import (
	"testing"

	"go.uber.org/zap"
)

func TestInitLog(t *testing.T) {
	Init(nil)
	Debug("debug", zap.Int("player", 1)) //no print
	Info("info", zap.Int("player", 1))
	Warn("warn", zap.Int("player", 1))
	Error("cerror", zap.Int("player", 1)) //print with stack
}

func TestInitDefault(t *testing.T) {
	InitDefault("")
	Debug("test debug", zap.Int("player", 1))
	Info("test info", zap.Int("player", 1))
	Error("test cerror", zap.Int("player", 1))
}

func TestWithFixedParam(t *testing.T) {
	logger, err := New(nil)
	if err != nil {
		t.Fatal(err)
	}
	l1 := logger.Named("testLog").With(zap.String("key1", "value1")) //The module has its own fixed name or parameter
	l2 := logger.With(zap.String("key2", "value2"))
	l1.Info("test1", zap.Int("player", 1))
	l2.Info("test2", zap.Int("player", 2))

	logger.With()

}

func TestNameSpace(t *testing.T) {
	logger := zap.NewExample()
	defer logger.Sync()

	logger.Info("tracked some metrics",
		zap.Namespace("metrics"),
		zap.Int("counter", 1),
	)

	logger2 := logger.With(
		zap.Namespace("metrics"),
		zap.Int("counter", 1),
	)
	logger2.Info("tracked some metrics")
}
