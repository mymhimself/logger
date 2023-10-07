package logger

import "github.com/sirupsen/logrus"

type Logger interface {
	Info(args ...interface{})
	InfoWithFields(fs map[string]interface{}, args ...interface{})

	Infof(format string, args ...interface{})
	InfofWithFields(fs map[string]interface{}, format string, args ...interface{})

	Warn(args ...interface{})
	WarnWithFields(fs map[string]interface{}, args ...interface{})

	Warnf(format string, args ...interface{})
	WarnfWithFields(fs map[string]interface{}, format string, args ...interface{})

	Debug(args ...interface{})
	DebugWithFields(fs map[string]interface{}, args ...interface{})

	Debugf(format string, args ...interface{})
	DebugfWithFields(fs map[string]interface{}, format string, args ...interface{})

	Fatal(v ...interface{})
	FatalWithFields(fs map[string]interface{}, v ...interface{})

	Error(args ...interface{})
	ErrorWithFields(fs map[string]interface{}, args ...interface{})

	Errorf(format string, args ...interface{})
	ErrorfWithFields(fs map[string]interface{}, format string, args ...interface{})

	Clone() Logger

	ClearMDC()
	MDCPut(key string, value interface{})
	MDCRemove(key string)
}

// map[string]interface{}
var Default Logger = &myLogger{
	l: logrus.New(),
	config: config{
		level: LevelInfo,
	},
}

// ─── FUNCTIONS ──────────────────────────────────────────────────────────────────

// Info like log.Println()
func Info(args ...interface{}) {
	Default.Info(args...)
}

// Infof like log.Printf()
func Infof(format string, args ...interface{}) {
	Default.Infof(format, args...)
}

// Warn like log.Println()
func Warn(args ...interface{}) {
	Default.Warn(args...)
}

// Warnf like log.Printf()
func Warnf(format string, args ...interface{}) {
	Default.Warnf(format, args...)
}

// Fatal like log.Println()
func Debug(args ...interface{}) {
	Default.Debug(args...)
}

// Debug like log.Println()
func Debugf(format string, args ...interface{}) {
	Default.Debugf(format, args...)
}

// Debugf like log.Printf()
func Fatal(v ...interface{}) {
	Default.Fatal(v...)
}

// Error like log.Println()
func Error(args ...interface{}) {
	Default.Error(args...)
}

// Errorf like log.Printf()
func Errorf(format string, args ...interface{}) {
	Default.Errorf(format, args...)
}

func InfoWithFields(fs map[string]interface{}, args ...interface{}) {
	Default.InfoWithFields(fs, args...)
}

func InfofWithFields(fs map[string]interface{}, format string, args ...interface{}) {
	Default.InfofWithFields(fs, format, args...)
}

func WarnWithFields(fs map[string]interface{}, args ...interface{}) {
	Default.WarnWithFields(fs, args...)
}

func WarnfWithFields(fs map[string]interface{}, format string, args ...interface{}) {
	Default.WarnfWithFields(fs, format, args...)
}

func DebugWithFields(fs map[string]interface{}, args ...interface{}) {
	Default.DebugWithFields(fs, args...)
}

func DebugfWithFields(fs map[string]interface{}, format string, args ...interface{}) {
	Default.DebugfWithFields(fs, format, args...)
}

func FatalWithFields(fs map[string]interface{}, v ...interface{}) {
	Default.FatalWithFields(fs, v...)
}

func ErrorWithFields(fs map[string]interface{}, args ...interface{}) {
	Default.ErrorWithFields(fs, args...)
}

func ErrorfWithFields(fs map[string]interface{}, format string, args ...interface{}) {
	Default.ErrorfWithFields(fs, format, args...)
}
