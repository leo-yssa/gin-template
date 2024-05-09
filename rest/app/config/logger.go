package config

import (
	"os"
	"strings"

	nested "github.com/antonfisher/nested-logrus-formatter"
	logger "github.com/sirupsen/logrus"
)

func InitLogger() {
	logger.SetLevel(getLoggerLevel(os.Getenv("LOG_LEVEL")))
	logger.SetReportCaller(true)
	logger.SetFormatter(&nested.Formatter{
		HideKeys:        true,
		FieldsOrder:     []string{"component", "category"},
		TimestampFormat: "2006-01-02 15:04:05",
		ShowFullLevel:   true,
		CallerFirst:     true,
	})
}

func getLoggerLevel(value string) logger.Level {
	switch strings.ToUpper(value) {
	case "DEBUG":
		return logger.DebugLevel
	case "TRACE":
		return logger.TraceLevel
	default:
		return logger.InfoLevel
	}
}