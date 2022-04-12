package mq

import (
	clientPulsar "github.com/apache/pulsar-client-go/pulsar"
	"github.com/xuexiangyou/code-art/common"
	"github.com/xuexiangyou/code-art/controllers"
	"github.com/xuexiangyou/code-art/pkg/pulsar"
	"time"
)

type TagController struct {
	controllers.BaseController
}

func NewTagController(routes map[string]pulsar.CallHandler, f common.FxCommonParams) {
	t := TagController{BaseController: controllers.NewBaseController(f.Db, f.Redis)}
	{
		routes["getTag"] = t.getTag()
		routes["getTagList"] = t.getTagList()
	}
}

func (t TagController) getTag() pulsar.CallHandler {
	return func(message clientPulsar.Message) (interface{}, error) {
		time.Sleep(time.Second)
		return "getTag", nil
	}
}

func (t TagController) getTagList() pulsar.CallHandler {
	return func(message clientPulsar.Message) (interface{}, error) {
		time.Sleep(time.Second)
		return "getTagList", nil
	}
}


