package logger

import (
	"github.com/sirupsen/logrus"
)

// A hook to be fired when logging on the logging levels returned from
// `Levels()` on your implementation of the interface. Note that this is not
// fired in a goroutine or a channel with workers, you should handle such
// functionality yourself if your call is non-blocking and you don't wish for
// the logging calls for levels returned from `Levels()` to block.
type Hook = logrus.Hook

type hook struct {
	levels []Level
	cb     func(*logrus.Entry) error
}

var _ logrus.Hook = &hook{}

func (h *hook) Levels() []Level {
	return h.levels
}

// Fire all the hooks for the passed level. Used by `entry.log` to fire
// appropriate hooks for a log entry.
func (h *hook) Fire(e *logrus.Entry) error {
	return h.cb(e)
}

// NewHook create new Hook to be used in logrus
func NewHook(levels []Level, fn func(*logrus.Entry) error) Hook {
	return &hook{
		levels: levels,
		cb:     fn,
	}
}
