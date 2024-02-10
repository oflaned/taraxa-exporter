package utils

import (
	"os"

	"github.com/gowizzard/dotenv"
	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

func InitLogger () {
	Logger.SetOutput(os.Stdout)
	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	if dotenv.String("STAGE") == "development" {
		Logger.SetLevel(logrus.DebugLevel)
		Logger.Info("Log level: ", Logger.GetLevel())
	} else {
		Logger.SetLevel(logrus.InfoLevel)
		Logger.Info("Log level: ", Logger.GetLevel())
	}
}