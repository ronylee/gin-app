package main

import (
	"context"
	"gin-app/config"
	"gin-app/model"
	"gin-app/pkg/logging"
	"gin-app/pkg/redis"
	"gin-app/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	config.InitConfig() // 初始化配置
	model.InitDB()      // 初始化数据库
	logging.InitLog()   // 初始化日志
	redis.InitRedis()   // 初始化redis

	routers := router.InitRouter() // 初始化路由，中间件

	srv := &http.Server{
		Addr:    ":" + config.Bases.Server.Port,
		Handler: routers,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
