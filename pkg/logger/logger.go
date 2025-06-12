package logger

import (
	"os"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger    *zap.Logger
	startTime time.Time
	initOnce  sync.Once
)

func ZapLoggerInit() {
	initOnce.Do(func() {
		startTime = time.Now()
		cfg := zap.NewProductionEncoderConfig()
		cfg.EncodeTime = zapcore.ISO8601TimeEncoder

		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(cfg),
			zapcore.AddSync(os.Stderr),
			zapcore.DebugLevel,
		)

		logger = zap.New(core, zap.AddCaller())
	})
}

func Info(msg string, fields ...zap.Field)  { logger.Info(msg, fields...) }
func Error(msg string, fields ...zap.Field) { logger.Error(msg, fields...) }
func Fatal(msg string, fields ...zap.Field) { logger.Fatal(msg, fields...) }
