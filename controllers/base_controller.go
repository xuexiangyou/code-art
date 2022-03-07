package controllers

import (
	"github.com/xuexiangyou/code-art/middleware/log"
	"go.uber.org/fx"
)

type BaseCtlParams struct {
	fx.In
	Logs *log.Logs
}

type BaseController struct {
	logs *log.Logs
}



