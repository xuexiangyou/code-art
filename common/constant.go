package common

import (
	"github.com/go-redis/redis/v8"
	"github.com/xuexiangyou/code-art/config"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

const (
	AppEnvDev 	= "dev"
	AppEnvTest 	= "test"
	AppEnvPrd   = "prd"
)

type FxCommonParams struct {
	fx.In
	Db                *gorm.DB
	Redis    		  *redis.Client
	Config            *config.Config
}