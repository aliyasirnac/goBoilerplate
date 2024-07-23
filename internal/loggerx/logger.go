package loggerx

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Formatter string

const (
	FormatterText Formatter = "text"
	FormatterJson Formatter = "json"
)

type Config struct {
	Level         string
	DisableColors bool
	Formatter     Formatter
}

func New(cfg Config) logrus.FieldLogger {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)

	logLevel, err := logrus.ParseLevel(cfg.Level)
	if err != nil {
		logLevel = logrus.InfoLevel
	}
	logger.SetLevel(logLevel)

	if cfg.Formatter == FormatterJson {
		logger.Formatter = &logrus.JSONFormatter{}
	} else {
		logger.Formatter = &logrus.TextFormatter{FullTimestamp: true, DisableColors: cfg.DisableColors}
	}

	return logger
}

func ExitOnError(err error, content string) {
	if err == nil {
		return
	}

	log := &logrus.Logger{
		Out:          os.Stdout,
		Formatter:    &logrus.JSONFormatter{},
		Hooks:        make(logrus.LevelHooks),
		Level:        logrus.InfoLevel,
		ExitFunc:     os.Exit,
		ReportCaller: false,
	}

	log.Fatalf("%s %s", content, err)
}

