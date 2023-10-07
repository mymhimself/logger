package logger

import (
	"fmt"
	"os"
	"runtime"
	"strconv"

	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
	glogrus "github.com/sirupsen/logrus"
)

type myLogger struct {
	l      *glogrus.Logger
	s      *sentry.Client
	config config
	mdc    map[string]interface{}
}

var _ Logger = &myLogger{}

// New will cerate a new logger based on the options. it will not update the Default
func New(ops ...Option) (Logger, error) {
	lg := &myLogger{
		l:      glogrus.New(),
		mdc:    make(map[string]interface{}),
		config: config{printCaller: -1},
	}
	for _, fn := range ops {
		err := fn(lg)
		if err != nil {
			return nil, err
		}
	}
	if lg.l.Out == nil {
		lg.l.SetOutput(os.Stdout)
	}
	if lg.config.level != 0 {
		lg.l.SetLevel(lg.config.level)
	}
	return lg, nil
}

// Init will initialize the logger and update the default value.
func Init(ops ...Option) error {
	d, err := New(ops...)
	if err != nil {
		return err
	}
	Default = d
	return nil
}

// ─── FUNCTIONS ──────────────────────────────────────────────────────────────────

// ─── External FUNCTIONS ──────────────────────────────────────────────────────────────────
// func (l *myLogger) InfofWithFields(fs map[string]interface{},format string, args ...interface{})
// func (l *myLogger) WarnWithFields(fs map[string]interface{},args ...interface{})
// func (l *myLogger) WarnfWithFields(fs map[string]interface{},format string, args ...interface{})
// func (l *myLogger) DebugWithFields(fs map[string]interface{},args ...interface{})
// func (l *myLogger) DebugfWithFields(fs map[string]interface{},format string, args ...interface{})
// func (l *myLogger) FatalWithFields(fs map[string]interface{},v ...interface{})
// func (l *myLogger) ErrorWithFields(fs map[string]interface{},args ...interface{})
// func (l *myLogger) ErrorfWithFields(fs map[string]interface{},format string, args ...interface{})
// Info like log.Println()
func (l *myLogger) Info(args ...interface{}) {
	name, ok := l.getCallerName()

	switch {
	case len(l.mdc) == 0 && !ok:
		l.l.Info(args...)
	case len(l.mdc) == 0 && ok:
		l.l.WithField("callerName", name).Info(args...)
	default:
		l.l.WithField("callerName", name).WithFields(l.getFields()).Info(args...)
	}
	if l.s != nil && l.config.level <= LevelInfo {
		l.s.CaptureMessage(fmt.Sprint(args...), nil, nil)
	}

}
func (l *myLogger) InfoWithFields(fs map[string]interface{}, args ...interface{}) {
	name, ok := l.getCallerName()
	switch {
	case !ok:
		l.l.WithFields(l.getFields()).WithFields(fs).Info(args...)
	default:
		l.l.WithField("callerName", name).WithFields(l.getFields()).WithFields(fs).Info(args...)
	}
}

// Infof like log.Printf()
func (l *myLogger) Infof(format string, args ...interface{}) {
	name, ok := l.getCallerName()

	switch {
	case len(l.mdc) == 0 && !ok:
		l.l.Infof(format, args...)
	case len(l.mdc) == 0 && ok:
		l.l.WithField("callerName", name).Infof(format, args...)
	default:
		l.l.WithField("callerName", name).WithFields(l.getFields()).
			Infof(format, args...)
	}

	if l.s != nil && l.config.level <= LevelInfo {
		l.s.CaptureMessage(fmt.Sprintf(format, args...), nil, nil)
	}

}

func (l *myLogger) InfofWithFields(fs map[string]interface{}, format string, args ...interface{}) {
	name, ok := l.getCallerName()
	switch {
	case !ok:
		l.l.WithFields(l.getFields()).WithFields(fs).Infof(format, args...)
	default:
		l.l.WithField("callerName", name).WithFields(l.getFields()).WithFields(fs).Infof(format, args...)
	}
}

// Warn like log.Println()
func (l *myLogger) Warn(args ...interface{}) {
	name, ok := l.getCallerName()

	switch {
	case len(l.mdc) == 0 && !ok:
		l.l.Warn(args...)
	case len(l.mdc) == 0 && ok:
		l.l.WithField("callerName", name).Warn(args...)
	default:
		l.l.WithField("callerName", name).WithFields(l.getFields()).Warn(args...)
	}
	if l.s != nil && l.config.level <= LevelInfo {
		l.s.CaptureMessage(fmt.Sprint(args...), nil, nil)
	}
}

func (l *myLogger) WarnWithFields(fs map[string]interface{}, args ...interface{}) {
	name, ok := l.getCallerName()
	switch {
	case !ok:
		l.l.WithFields(l.getFields()).WithFields(fs).Warn(args...)
	default:
		l.l.WithField("callerName", name).WithFields(l.getFields()).WithFields(fs).Warn(args...)
	}
}

// Warnf like log.Printf()
func (l *myLogger) Warnf(format string, args ...interface{}) {
	name, ok := l.getCallerName()

	switch {
	case len(l.mdc) == 0 && !ok:
		l.l.Warnf(format, args...)
	case len(l.mdc) == 0 && ok:
		l.l.WithField("callerName", name).Warnf(format, args...)
	default:
		l.l.WithField("callerName", name).WithFields(l.getFields()).
			Warnf(format, args...)
	}
}

func (l *myLogger) WarnfWithFields(fs map[string]interface{}, format string, args ...interface{}) {
	name, ok := l.getCallerName()
	switch {
	case !ok:
		l.l.WithFields(l.getFields()).WithFields(fs).Warnf(format, args...)
	default:
		l.l.WithField("callerName", name).WithFields(l.getFields()).WithFields(fs).Warnf(format, args...)
	}
}

// Fatal like log.Println()
func (l *myLogger) Fatal(v ...interface{}) {
	l.l.Fatal(v...)
}
func (l *myLogger) FatalWithFields(fs map[string]interface{}, args ...interface{}) {
	name, ok := l.getCallerName()
	switch {
	case !ok:
		l.l.WithFields(l.getFields()).WithFields(fs).Fatal(args...)
	default:
		l.l.WithField("callerName", name).WithFields(l.getFields()).WithFields(fs).Fatal(args...)
	}
}

// Debug like log.Println()
func (l *myLogger) Debug(args ...interface{}) {
	name, ok := l.getCallerName()

	switch {
	case len(l.mdc) == 0 && !ok:
		l.l.Debug(args...)
	case len(l.mdc) == 0 && ok:
		l.l.WithField("callerName", name).Debug(args...)
	default:
		l.l.WithField("callerName", name).WithFields(l.getFields()).
			Debug(args...)

	}
}
func (l *myLogger) DebugWithFields(fs map[string]interface{}, args ...interface{}) {
	name, ok := l.getCallerName()
	switch {
	case !ok:
		l.l.WithFields(l.getFields()).WithFields(fs).Debug(args...)
	default:
		l.l.WithField("callerName", name).WithFields(l.getFields()).WithFields(fs).Debug(args...)
	}
}

// Debugf like log.Printf()
func (l *myLogger) Debugf(format string, args ...interface{}) {
	name, ok := l.getCallerName()

	switch {
	case len(l.mdc) == 0 && !ok:
		l.l.Debugf(format, args...)
	case len(l.mdc) == 0 && ok:
		l.l.WithField("callerName", name).Debugf(format, args...)
	default:
		l.l.WithField("callerName", name).WithFields(l.getFields()).
			Debugf(format, args...)

	}
}
func (l *myLogger) DebugfWithFields(fs map[string]interface{}, format string, args ...interface{}) {
	name, ok := l.getCallerName()
	switch {
	case !ok:
		l.l.WithFields(l.getFields()).WithFields(fs).Debugf(format, args...)
	default:
		l.l.WithField("callerName", name).WithFields(l.getFields()).WithFields(fs).Debugf(format, args...)
	}
}

// Error like log.Println()
func (l *myLogger) Error(args ...interface{}) {
	name, ok := l.getCallerName()

	switch {
	case len(l.mdc) == 0 && !ok:
		l.l.Error(args...)
	case len(l.mdc) == 0 && ok:
		l.l.WithField("callerName", name).Error(args...)
	default:
		l.l.WithField("callerName", name).WithFields(l.getFields()).
			Error(args...)

	}
}
func (l *myLogger) ErrorWithFields(fs map[string]interface{}, args ...interface{}) {
	name, ok := l.getCallerName()
	switch {
	case !ok:
		l.l.WithFields(l.getFields()).WithFields(fs).Error(args...)
	default:
		l.l.WithField("callerName", name).WithFields(l.getFields()).WithFields(fs).Error(args...)
	}
}

// Errorf like log.Printf()
func (l *myLogger) Errorf(format string, args ...interface{}) {
	name, ok := l.getCallerName()

	switch {
	case len(l.mdc) == 0 && !ok:
		l.l.Errorf(format, args...)
	case len(l.mdc) == 0 && ok:
		l.l.WithField("callerName", name).Errorf(format, args...)
	default:
		l.l.WithField("callerName", name).WithFields(l.getFields()).
			Errorf(format, args...)

	}
}
func (l *myLogger) ErrorfWithFields(fs map[string]interface{}, format string, args ...interface{}) {
	name, ok := l.getCallerName()
	switch {
	case !ok:
		l.l.WithFields(l.getFields()).WithFields(fs).Errorf(format, args...)
	default:
		l.l.WithField("callerName", name).WithFields(l.getFields()).WithFields(fs).Errorf(format, args...)
	}
}

// ─── UTILS FUNCTIONS ────────────────────────────────────────────────────────────

func (l *myLogger) getCallerName() (string, bool) {
	if l.config.printCaller < 0 {
		return "", false
	}
	pc, _, line, ok := runtime.Caller(l.config.printCaller)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		return details.Name() + ":" + strconv.Itoa(line), ok
	}

	return "", ok
}

func (l *myLogger) getFields() logrus.Fields {
	fields := logrus.Fields{}
	for key, val := range l.mdc {
		fields[key] = val
	}
	return fields
}

// MDCPut Sets the key-value in the mdc of the log
func (l *myLogger) MDCPut(key string, value interface{}) {
	l.mdc[key] = value
}

// ClearMDC Clears the key-value in the mdc of the log
func (l *myLogger) ClearMDC() {
	l.mdc = make(map[string]interface{})
}

// MDCRemove Removes a single key from mdc
func (l *myLogger) MDCRemove(key string) {
	delete(l.mdc, key)
}

func (l *myLogger) Clone() Logger {
	return &myLogger{
		l: l.l,
		s: l.s,
		config: config{
			serviceName: l.config.serviceName,
			logFileName: l.config.logFileName,
			level:       l.config.level,
			printCaller: l.config.printCaller,
		},
		mdc: make(map[string]interface{}),
	}
}
