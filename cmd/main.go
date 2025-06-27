package main

import (
	"fmt"

	"github.com/polyk005/locale"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

func main() {
	fmt.Println("Hello locale")
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := gotenv.Load(); err != nil {
		logrus.Fatalf("Error loading env: %s", err.Error())
	}

	if err := InitConfig(); err != nil {
		logrus.Fatalf("Error loading config: %s", err.Error())
	}

	srv := new(locale.Server)
	if err := srv.Run(viper.GetString("port")); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
