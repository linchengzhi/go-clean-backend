package logger

import (
	"io"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Config struct {
	Level      string `yaml:"level" json:"level" default:"info"`
	Encoding   string `yaml:"encoding" json:"encoding" default:"console"`
	EncodeTime string `yaml:"encode_time" json:"encode_time" default:"2006-01-02T15:04:05.000Z0700"`
	UseStdout  bool   `yaml:"use_stdout" json:"use_stdout" default:"true"`
	Filename   string `yaml:"filename" json:"filename" default:"./logs/service.clog"`
	MaxAge     int    `yaml:"max_age" json:"max_age" default:"15"`    // days
	MaxSize    int    `yaml:"max_size" json:"max_size" default:"500"` // MB
	MaxBackups int    `yaml:"max_backups" json:"max_backups" default:"15"`
	Compress   bool   `yaml:"compress" json:"compress" default:"true"`
}

var zapLog *zap.Logger
var zapLogSugar *zap.SugaredLogger

func Init(cfg *Config) {
	logger, err := New(cfg)
	if err != nil {
		panic(err)
	}
	logger.WithOptions(zap.AddCallerSkip(1))
	zapLog = logger
	zapLogSugar = logger.Sugar()
}

// if cfg is nil, use default config
func New(cfg *Config) (*zap.Logger, error) {
	if cfg == nil {
		cfg = getDefaultCfg()
	}

	// Initialize level
	level := zap.InfoLevel
	if err := level.UnmarshalText([]byte(cfg.Level)); err != nil {
		return nil, err
	}
	autoLevel := zap.NewAtomicLevelAt(level)

	// Customize the encoder configuration
	c := zap.NewProductionConfig()

	if cfg.EncodeTime != "" {
		c.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format(cfg.EncodeTime))
		}
	}
	// Create a console encoder with the custom configuration
	consoleEncoder := zapcore.NewConsoleEncoder(c.EncoderConfig)

	// Create a zapcore.Core for each output (file and stdout)
	fileCore := zapcore.NewCore(
		consoleEncoder,
		zapcore.AddSync(newWriter(cfg)),
		autoLevel,
	)
	stdoutCore := zapcore.NewCore(
		consoleEncoder,
		zapcore.AddSync(os.Stdout),
		autoLevel,
	)
	if !cfg.UseStdout {
		stdoutCore = zapcore.NewNopCore()
	}

	// Combine the cores
	combinedCore := zapcore.NewTee(
		fileCore,
		stdoutCore,
	)

	// Create the logger with the combined core
	logger := zap.New(combinedCore, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel), zap.AddCallerSkip(0))
	return logger, nil
}

// for test, if env is "", use dev model
func InitDefault(env string) {
	var logger *zap.Logger
	switch env {
	case "dev", "":
		logger, _ = zap.NewDevelopment()
	case "example":
		logger = zap.NewExample()
	case "prod":
		logger, _ = zap.NewProduction()
	}
	zapLog = logger
	zapLogSugar = logger.Sugar()
}

func newWriter(cfg *Config) io.Writer {
	return &lumberjack.Logger{
		Filename:   cfg.Filename,
		MaxSize:    cfg.MaxSize, // megabytes
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge, //days
		Compress:   cfg.Compress,
	}
}

func getDefaultCfg() *Config {
	cfg := new(Config)
	cfg.Level = "info"
	cfg.Encoding = "console"
	cfg.EncodeTime = "2006-01-02T15:04:05.000Z0700"
	cfg.UseStdout = true
	cfg.Filename = "./logs/service.log"
	cfg.MaxAge = 15
	cfg.MaxSize = 512
	cfg.MaxBackups = 30
	cfg.Compress = true
	return cfg
}
