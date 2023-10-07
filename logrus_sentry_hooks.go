package logger

import (
	"errors"

	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
)

//HookCaptureMessage is a default hook for sentry to capture the message for Info, Debug, Trace.
var HookCaptureMessage func(s *sentry.Client) logrus.Hook = func(s *sentry.Client) logrus.Hook {
	return NewHook([]Level{LevelInfo, LevelDebug, LevelTrace}, func(e *logrus.Entry) error {
		s.CaptureMessage(e.Message, nil, nil)
		return nil
	})
}

//HookCaptureError default hooks for sentry to capture errors in Log.Error
var HookCaptureError func(s *sentry.Client) logrus.Hook = func(s *sentry.Client) logrus.Hook {
	return NewHook([]Level{LevelError, LevelFatal}, func(e *logrus.Entry) error {
		s.CaptureException(errors.New(e.Message), nil, nil)
		return nil
	})
}
