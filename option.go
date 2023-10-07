package logger

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/getsentry/sentry-go"
)

type Option func(*myLogger) error

func OptionServiceName(str string) Option {
	return func(ml *myLogger) error {
		ml.config.serviceName = str
		return nil
	}
}

func OptionPrintToFile(filename string) Option {
	return func(ml *myLogger) error {
		ml.config.logFileName = filename

		dir := filepath.Dir(filename)
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			err := os.Mkdir(dir, 0755)
			if err != nil {
				fmt.Printf("Not able to create directory. error: %v", err.Error())
				return err
			}
		}

		f, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755)
		if err != nil {
			fmt.Printf("Not able to open logfile for writing. error: %v", err.Error())
			return err
		}
		ml.l.SetOutput(f)

		return nil
	}

}

func OptionSetLevel(level Level) Option {
	return func(ml *myLogger) error {
		ml.config.level = level

		return nil
	}
}

func OptionSetFormatter(formatter Formatter) Option {
	return func(ml *myLogger) error {
		ml.l.SetFormatter(formatter)
		return nil
	}
}

func OptionWithHook(h *hook) Option {
	return func(ml *myLogger) error {
		ml.l.AddHook(h)
		return nil
	}
}

// OptionReportCaller set the report caller, the offset or skip is depends on the development pattern.
// mostly set it to 3
func OptionReportCaller(offset int) Option {
	return func(ml *myLogger) error {
		ml.config.printCaller = offset
		// ml.l.SetReportCaller()
		return nil
	}
}

// OptionWithSentry is a small integration between the Logger and Sentry.io, if you want you can define a specific hooks for that. see OptionWithHooks
func OptionWithSentry(DSN string, Environment string, Release string) Option {

	return func(ml *myLogger) error {
		sDebug := false
		if ml.config.level == LevelDebug {
			sDebug = true
		}
		c, err := sentry.NewClient(sentry.ClientOptions{
			Dsn:         DSN,
			Environment: Environment,
			Release:     Release,
			Debug:       sDebug,
		})
		if err != nil {
			return err
		}
		ml.s = c
		func() {
			defer sentry.Recover()
			// do all of the scary things here
		}()
		ml.l.AddHook(HookCaptureMessage(ml.s))
		ml.l.AddHook(HookCaptureError(ml.s))
		return nil
	}

}

// OptionWithHooks can be use to define custom Hooks, Hooks will fire once a log in the levels being triaged
func OptionWithHooks(hks ...Hook) Option {
	return func(ml *myLogger) error {
		for _, h := range hks {
			ml.l.AddHook(h)
		}
		return nil
	}
}
