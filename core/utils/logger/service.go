package logger

import (
	"core-system/core/entities/logger"

	"github.com/sirupsen/logrus"
)

func NewLoggerService(cfg *Config) logger.LoggerInterface {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
	return &config{
		log:   logger,
		level: cfg.Level,
		color: cfg.Color,
	}
}

func (s *config) Println(args ...interface{}) {
	s.log.Println(args...)
}

func (s *config) Errorln(args ...interface{}) {
	s.log.Errorln(args...)
}
