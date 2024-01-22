package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func defaultConfig() zap.Config {
	// Generate default config
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	config.DisableCaller = true
	config.DisableStacktrace = true
	config.EncoderConfig.TimeKey = "time"
	config.EncoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder

	return config
}
