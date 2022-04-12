package container

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuexiangyou/code-art/config"
	"github.com/xuexiangyou/code-art/pkg/pulsar"
	"github.com/xuexiangyou/code-art/pkg/setting"
	"github.com/xuexiangyou/code-art/pkg/stores"
	"github.com/xuexiangyou/code-art/routers"
	"go.uber.org/fx"
	"net/http"
)

func fxProvideConfig() fx.Option {
	return fx.Provide(setting.NewConfig)
}

func fxProvideDb() fx.Option {
	return fx.Provide(stores.ConnectDatabase)
}

func fxProvideRedis() fx.Option {
	return fx.Provide(stores.ConnectRedis)
}

func fxProvideRouter() fx.Option {
	return fx.Provide(
		routers.InitGinRouter, //添加http接口路由
		routers.InitQueueRouter,
	)
}

//fxRegister 启动服务
func fxRegister() fx.Option {
	return fx.Invoke(func(lc fx.Lifecycle, config *config.Config, r *gin.Engine, pulsar *pulsar.PulsarQueues) {
		srv := &http.Server{
			Addr:    fmt.Sprintf("%s:%d", config.Host, config.Port),
			Handler: r,
		}

		lc.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				go srv.ListenAndServe()
				go pulsar.Start() //队列消费服务启动

				return nil
			},
			OnStop: func(ctx context.Context) error {
				pulsar.Stop() //队列退出
				return srv.Shutdown(ctx)
			},
		})
	})
}

func NewApp() *fx.App {
	app := fx.New(
		fxProvideConfig(),     //配置文件
		fxProvideDb(),         //数据库文件
		fxProvideRedis(),      //redis文件
		fxProvideRouter(),     //路由文件
		fxRegister(),          //http 服务启动
	)

	return app
}
