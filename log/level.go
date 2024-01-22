package log

import "go.uber.org/zap"

// Parse the log level
func parse(level string) zap.AtomicLevel {
	// Parse
	if parsed, err := zap.ParseAtomicLevel(level); err == nil {
		return parsed
	}

	// Fallback to default in case of passing error
	return zap.NewAtomicLevelAt(zap.InfoLevel)
}
