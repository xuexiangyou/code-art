package mq

import (
	"github.com/xuexiangyou/code-art/common"
	"github.com/xuexiangyou/code-art/pkg/pulsar"
)

func NewRouter(f common.FxCommonParams) map[string]pulsar.CallHandler {
	routes := make(map[string]pulsar.CallHandler)
	{
		NewTagController(routes, f)
	}
	return routes
}
