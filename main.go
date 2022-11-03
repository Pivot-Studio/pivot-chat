package main

import (
	"github.com/Pivot-Studio/pivot-chat/api"
	"github.com/sirupsen/logrus"
)

func main() {
	err := api.Engine.Run()
	if err != nil {
		logrus.Fatalf("[main] engine run err:%+v", err)
	}

}
