package container

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuexiangyou/code-art/config"
	"github.com/xuexiangyou/code-art/infrastructure/setting"
	"github.com/xuexiangyou/code-art/infrastructure/stores"
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

/*func fxProvideController() fx.Option {
	return fx.Provide(
		v1.NewTagController,
		v1.NewArticleController,
	)
}*/

func fxProvideRouter() fx.Option {
	return fx.Provide(routers.InitRouter)
}

//fxRegister 启动服务
func fxRegister() fx.Option {
	return fx.Invoke(func(lc fx.Lifecycle, config *config.Config, r *gin.Engine) {
		srv := &http.Server{
			Addr:    fmt.Sprintf("%s:%d", config.Host, config.Port),
			Handler: r,
		}

		lc.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				go srv.ListenAndServe()
				return nil
			},
			OnStop: func(ctx context.Context) error {
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
		//fxProvideController(), //控制器文件
		fxProvideRouter(),     //路由文件
		fxRegister(),          //http 服务启动
	)

	return app
}
