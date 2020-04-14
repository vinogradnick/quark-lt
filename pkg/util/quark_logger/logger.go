package quark_logger

import (
	"github.com/sirupsen/logrus"
	"log"
)

const (
	STDOUT_LOGGER   int = 1
	ROTATE_LOGGER   int = 2
	DATABASE_LOGGER int = 3
)

/**
Установка логгера в проект
*/
func SetupLogger(loggerType int) {
	logger := logrus.New()
	log.SetOutput(logger.Writer())
}
/*


lumberjackLogger := &lumberjack.Logger{
  Filename:   "/var/log/misc.log",
  MaxSize:    10,
  MaxBackups: 3,
  MaxAge:     3,
  LocalTime:  true,
}
logrus.SetOutput(lumberjackLogger)
 */