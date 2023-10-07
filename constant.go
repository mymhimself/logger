package logger

import "github.com/sirupsen/logrus"

// Level the level of logging
type Level = logrus.Level

const (

	// LevelPanic level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	LevelPanic = logrus.PanicLevel

	// LevelFatal level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	LevelFatal = logrus.FatalLevel

	// LevelError level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	LevelError = logrus.ErrorLevel

	// LevelWarn level. Non-critical entries that deserve eyes.
	LevelWarn = logrus.WarnLevel

	// LevelInfo level. General operational entries about what's going on inside the
	// application.
	LevelInfo = logrus.InfoLevel

	// LevelDebug level. Usually only enabled when debugging. Very verbose logging.
	LevelDebug = logrus.DebugLevel

	// LevelTrace level. Designates finer-grained informational events than the Debug.
	LevelTrace = logrus.TraceLevel
)

// Formatter the type of each log record
type Formatter = logrus.Formatter

var (
	//FormatterText the format of logs is simple text
	FormatterText logrus.Formatter = &logrus.TextFormatter{}

	//FormatterJSON the format of logs is jSON
	FormatterJSON logrus.Formatter = &logrus.JSONFormatter{}

	//you can add custom formatter

)
