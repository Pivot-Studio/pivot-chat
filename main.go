package main

import (
	"github.com/sirupsen/logrus"
	"pivot-chat/api"
)

func main() {
	err := api.Engine.Run()
	if err != nil {
		logrus.Fatalf("[main] engine run err:%+v", err)
	}

}
