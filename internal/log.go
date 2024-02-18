package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

func init() {
  file, err := os.OpenFile("flappy.json", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
  if err == nil {
   Logger.Out = file
  } else {
   panic("Failed to log to file, using default stderr")
  }
  Logger.Formatter = &logrus.JSONFormatter{}
}
