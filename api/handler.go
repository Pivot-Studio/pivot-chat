package api

import (
	"encoding/json"
	"github.com/Pivot-Studio/pivot-chat/model"
	"github.com/Pivot-Studio/pivot-chat/service"
	"github.com/sirupsen/logrus"
)

func (c *WsConnContext) Message(data []byte) {
	meg := model.Message{}
	err := json.Unmarshal(data, &meg)
	if err != nil {
		logrus.Errorf("[Message] json unmarshal %+v", err)
		return
	}

	service.HandleGroupMessage(&meg)
}
