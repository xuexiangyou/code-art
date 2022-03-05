package container

import (
	"fmt"
	"github.com/xuexiangyou/code-art/controllers"
	"github.com/xuexiangyou/code-art/domain/strategy"
	"github.com/xuexiangyou/code-art/infrastructure/setting"
	"github.com/xuexiangyou/code-art/infrastructure/stores"
	"github.com/xuexiangyou/code-art/middleware/log"
	"github.com/xuexiangyou/code-art/routers"
	"github.com/xuexiangyou/code-art/services"
	"go.uber.org/dig"
)

func initConfig(container *dig.Container) {
	container.Provide(setting.NewConfig)
}

func initLogs(container *dig.Container) {
	container.Provide(log.NewLogs)
}

func initDb(container *dig.Container) {
	container.Provide(stores.ConnectDatabase)
}

func initRedis(container *dig.Container) {
	container.Provide(stores.ConnectRedis)
}

func initStrategy(container *dig.Container) {
	container.Provide(strategy.NewTagStrategy)
}

func initService(container *dig.Container) {
	container.Provide(services.NewTagService)
}

func initController(container *dig.Container) {
	err := container.Provide(controllers.NewTagController)
	fmt.Println("??????", err)
}

func initRouter(container *dig.Container) {
	container.Provide(routers.InitRouter)
}

func Init(c *dig.Container) {
	initConfig(c)
	initLogs(c)
	initDb(c)
	initRedis(c)
	initStrategy(c)
	initService(c)
	initController(c)
	initRouter(c)
}


