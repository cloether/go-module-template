package applog

import (
	"github.com/cloether/go-module-template/internal/applog"
	"os"
)

func init() {
	SetLogger(newLogger())
}

// V reports whether verbosity level l is at least the requested verbose level.
func V(l int) bool {
	return applog.Logger.V(l)
}

// Info logs to the INFO log.
func Info(args ...interface{}) {
	applog.Logger.Info(args...)
}

// Infof logs to the INFO log.
// Arguments are handled in the manner of fmt.Printf.
func Infof(format string, args ...interface{}) {
	applog.Logger.Infof(format, args...)
}

// Infoln logs to the INFO log.
// Arguments are handled in the manner of fmt.Println.
func Infoln(args ...interface{}) {
	applog.Logger.Infoln(args...)
}

// Warning logs to the WARNING log.
func Warning(args ...interface{}) {
	applog.Logger.Warning(args...)
}

// Warningf logs to the WARNING log.
// Arguments are handled in the manner of fmt.Printf.
func Warningf(format string, args ...interface{}) {
	applog.Logger.Warningf(format, args...)
}

// Warningln logs to the WARNING log.
// Arguments are handled in the manner of fmt.Println.
func Warningln(args ...interface{}) {
	applog.Logger.Warningln(args...)
}

// Error logs to the ERROR log.
func Error(args ...interface{}) {
	applog.Logger.Error(args...)
}

// Errorf logs to the ERROR log.
// Arguments are handled in the manner of fmt.Printf.
func Errorf(format string, args ...interface{}) {
	applog.Logger.Errorf(format, args...)
}

// Errorln logs to the ERROR log.
// Arguments are handled in the manner of fmt.Println.
func Errorln(args ...interface{}) {
	applog.Logger.Errorln(args...)
}

// Fatal logs to the FATAL log.
// Arguments are handled in the manner of fmt.Print.
// It calls os.Exit() with exit code 1.
func Fatal(args ...interface{}) {
	applog.Logger.Fatal(args...)
	// Make sure fatal logs will exit.
	os.Exit(1)
}

// Fatalf logs to the FATAL log.
// Arguments are handled in the manner of fmt.Printf.
// It calls os.Exit() with exit code 1.
func Fatalf(format string, args ...interface{}) {
	applog.Logger.Fatalf(format, args...)
	// Make sure fatal logs will exit.
	os.Exit(1)
}

// Fatalln logs to the FATAL log.
// Arguments are handled in the manner of fmt.Println.
// It calls os.Exit()) with exit code 1.
func Fatalln(args ...interface{}) {
	applog.Logger.Fatalln(args...)
	// Make sure fatal logs will exit.
	os.Exit(1)
}
