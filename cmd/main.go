package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuexiangyou/code-art/infrastructure/container"
	"go.uber.org/dig"
)

func main() {
	digContainer := dig.New()

	//初始化依赖加载
	container.Init(digContainer)

	//加载路由配置
	err := digContainer.Invoke(func(r *gin.Engine) {
		r.Run(":"+ "8080")
	})

	if err != nil {
		fmt.Println(err)
		return
	}
}
