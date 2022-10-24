package conf

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

var C *Config

type Config struct {
	DB struct {
		UserName string
		Password string
		Host     string
		Schema   string
	}
	Redis struct {
		Host     string
		Password string
	}
	EmailServer struct {
		Email    string
		Port     int
		Host     string
		Password string
	}
}
var K8SConfig = "/etc/chat/config.json"

func init() {
	C = &Config{}
	data, err := ioutil.ReadFile(K8SConfig)
	if err != nil {
		logrus.Fatal("[init] init config error:", err, " path:", K8SConfig)
	}
	err = json.Unmarshal(data, C)
	if err != nil {
		logrus.Fatal("[init] init json parse %v", err)
	}
}
