package main

import (
	"context"
	"github.com/xuexiangyou/code-art/pkg/container"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	//容器初始化、并加载相关对象
	app := container.NewApp()
	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	stopCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Stop(stopCtx); err != nil {
		log.Fatal(err)
	}

	/*digContainer := dig.New()

	//初始化依赖加载
	container.Init(digContainer)

	var srv *http.Server

	//启动服务并支持优雅关闭
	go func() {
		err := digContainer.Invoke(func(r *gin.Engine) {
			srv = &http.Server{
				Addr:    ":8080",
				Handler: r,
			}
			// service connections
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("listen: %s\n", err)
			}
		})
		if err != nil {
			log.Fatalf("service start: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}*/
}
