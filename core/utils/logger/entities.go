package logger

import (
	"github.com/sirupsen/logrus"
)

type Config struct {
	Level string
	Color bool
}

type config struct {
	log   *logrus.Logger
	level string
	color bool
}
