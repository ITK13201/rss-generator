package services

import (
	"github.com/ITK13201/rss-generator/domain"
	"github.com/sirupsen/logrus"
	"os"
)

func NewLogger(cfg *domain.Config) *logrus.Logger {
	logger := logrus.New()
	{
		if (*cfg).Debug {
			logger.Level = logrus.DebugLevel
		} else {
			logger.Level = logrus.InfoLevel
		}
		logger.Formatter = &logrus.TextFormatter{}
		logger.Out = os.Stdout
	}
	return logger
}
