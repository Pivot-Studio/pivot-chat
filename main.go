package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	err := api.Engine.run()
	if err != nil {
		logrus.Fatalf("[main] engine run err:%+v", err)
	}
}