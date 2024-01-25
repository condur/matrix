package log

import (
	"fmt"

	"go.uber.org/zap"
)

// Define and initialize the logger
var logger, _ = defaultConfig().Build()

// Config
// sets the application name ('app') value and the minimum enabled logging level
//
// Default level value: "info".
// Accepted values: "panic", "fatal", "error", "warn", "warning", "info", "debug".
// In case of invalid input value, the log level is reset to default value.
func Config(app, level string) {
	// Gen the logger default config
	config := defaultConfig()

	// Set log level
	config.Level = parse(level)

	// Reinitialize the logger
	logger, _ = config.Build(zap.Fields(zap.String("app", app)))
}

// Debug logs a message at DebugLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Debug(msg string, fields ...Field) {
	logger.Debug(msg, convert(fields)...)
}

// Info logs a message at InfoLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Info(msg string, fields ...Field) {
	logger.Info(msg, convert(fields)...)
}

// Warn logs a message at WarnLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Warn(msg string, fields ...Field) {
	logger.Warn(msg, convert(fields)...)
}

// Error logs a message at ErrorLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Error(msg string, fields ...Field) {
	logger.Error(msg, convert(fields)...)
}

// Panic logs a message at PanicLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then panics, even if logging at PanicLevel is disabled.
func Panic(msg string, fields ...Field) {
	logger.Panic(msg, convert(fields)...)
}

// Fatal logs a message at FatalLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then calls os.Exit(1), even if logging at FatalLevel is
// disabled.
func Fatal(msg string, fields ...Field) {
	logger.Fatal(msg, convert(fields)...)
}

// Debugf uses fmt.Sprintf to log a templated message.
func Debugf(template string, args ...any) {
	logger.Debug(fmt.Sprintf(template, args...))
}

// Infof uses fmt.Sprintf to log a templated message.
func Infof(template string, args ...any) {
	logger.Info(fmt.Sprintf(template, args...))
}

// Warnf uses fmt.Sprintf to log a templated message.
func Warnf(template string, args ...any) {
	logger.Warn(fmt.Sprintf(template, args...))
}

// Errorf uses fmt.Sprintf to log a templated message.
func Errorf(template string, args ...any) {
	logger.Error(fmt.Sprintf(template, args...))
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func Panicf(template string, args ...any) {
	logger.Panic(fmt.Sprintf(template, args...))
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func Fatalf(template string, args ...any) {
	logger.Fatal(fmt.Sprintf(template, args...))
}
