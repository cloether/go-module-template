// Package applog defines depth logging.
package applog

import "os"

var (
	// Logger is the logger used for the non-depth log functions.
	Logger LoggerI

	// DepthLogger is the logger used for the depth log functions.
	DepthLogger DepthLoggerI
)

// InfoDepth logs to the INFO log at the specified depth.
func InfoDepth(depth int, args ...interface{}) {
	if DepthLogger != nil {
		DepthLogger.InfoDepth(depth, args...)
	} else {
		Logger.Infoln(args...)
	}
}

// WarningDepth logs to the WARNING log at the specified depth.
func WarningDepth(depth int, args ...interface{}) {
	if DepthLogger != nil {
		DepthLogger.WarningDepth(depth, args...)
	} else {
		Logger.Warningln(args...)
	}
}

// ErrorDepth logs to the ERROR log at the specified depth.
func ErrorDepth(depth int, args ...interface{}) {
	if DepthLogger != nil {
		DepthLogger.ErrorDepth(depth, args...)
	} else {
		Logger.Errorln(args...)
	}
}

// FatalDepth logs to the FATAL log at the specified depth.
func FatalDepth(depth int, args ...interface{}) {
	if DepthLogger != nil {
		DepthLogger.FatalDepth(depth, args...)
	} else {
		Logger.Fatalln(args...)
	}
	os.Exit(1)
}

// LoggerI does underlying logging work.
// This is a copy of the Logger defined in the external applog package. It
// is defined here to avoid a circular dependency.
type LoggerI interface {

	// Info logs to INFO log.
	// Arguments are handled in the manner of fmt.Print.
	Info(args ...interface{})

	// Infoln logs to INFO log.
	// Arguments are handled in the manner of fmt.Println.
	Infoln(args ...interface{})

	// Infof logs to INFO log.
	// Arguments are handled in the manner of fmt.Printf.
	Infof(format string, args ...interface{})

	// Warning logs to WARNING log.
	// Arguments are handled in the manner of fmt.Print.
	Warning(args ...interface{})

	// Warningln logs to WARNING log.
	// Arguments are handled in the manner of fmt.Println.
	Warningln(args ...interface{})

	// Warningf logs to WARNING log.
	// Arguments are handled in the manner of fmt.Printf.
	Warningf(format string, args ...interface{})

	// Error logs to ERROR log.
	// Arguments are handled in the manner of fmt.Print.
	Error(args ...interface{})

	// Errorln logs to ERROR log.
	// Arguments are handled in the manner of fmt.Println.
	Errorln(args ...interface{})

	// Errorf logs to ERROR log.
	// Arguments are handled in the manner of fmt.Printf.
	Errorf(format string, args ...interface{})

	// Fatal logs to ERROR log.
	// Arguments are handled in the manner of fmt.Print.
	// gRPC ensures that all Fatal logs will exit with os.Exit(1).
	// Implementations may also call os.Exit() with a non-zero exit code.
	Fatal(args ...interface{})

	// Fatalln logs to ERROR log.
	// Arguments are handled in the manner of fmt.Println.
	// gRPC ensures that all Fatal logs will exit with os.Exit(1).
	// Implementations may also call os.Exit() with a non-zero exit code.
	Fatalln(args ...interface{})

	// Fatalf logs to ERROR log.
	// Arguments are handled in the manner of fmt.Printf.
	// gRPC ensures that all Fatal logs will exit with os.Exit(1).
	// Implementations may also call os.Exit() with a non-zero exit code.
	Fatalf(format string, args ...interface{})

	// V reports whether verbosity level l is at least the requested verbose level.
	V(l int) bool
}

// DepthLoggerI logs at a specified call frame. If a LoggerI also implements
// DepthLoggerI, the below functions will be called with the appropriate stack
// depth set for trivial functions the logger may ignore.
// This is a copy of the DepthLogger defined in the external applog package.
// It is defined here to avoid a circular dependency.
//
// This API is EXPERIMENTAL.
type DepthLoggerI interface {
	// InfoDepth logs to INFO log at the specified depth.
	// Arguments are handled in the manner of fmt.Print.
	InfoDepth(depth int, args ...interface{})

	// WarningDepth logs to WARNING log at the specified depth.
	// Arguments are handled in the manner of fmt.Print.
	WarningDepth(depth int, args ...interface{})

	// ErrorDepth logs to ERROR log at the specified depth.
	// Arguments are handled in the manner of fmt.Print.
	ErrorDepth(depth int, args ...interface{})

	// FatalDepth logs to FATAL log at the specified depth.
	// Arguments are handled in the manner of fmt.Print.
	FatalDepth(depth int, args ...interface{})
}
