package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lawtech0902/go_gin_example/models"
	"github.com/lawtech0902/go_gin_example/pkg/gredis"
	"github.com/lawtech0902/go_gin_example/pkg/logging"
	"github.com/lawtech0902/go_gin_example/pkg/setting"
	"github.com/lawtech0902/go_gin_example/routers"
	"log"
	"net/http"
)

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)
	
	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20
	
	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	
	log.Printf("[info] start http server listening %s", endPoint)
	
	_ = server.ListenAndServe()
	
	// sig 热更新
	// endless.DefaultReadTimeOut = setting.ReadTimeout
	// endless.DefaultWriteTimeOut = setting.WriteTimeout
	// endless.DefaultMaxHeaderBytes = 1 << 20
	// endpoint := fmt.Sprintf(":%d", setting.HTTPPort)
	//
	// server := endless.NewServer(endpoint, router)
	// server.BeforeBegin = func(add string) {
	// 	log.Printf("Actual pid is %d", syscall.Getpid())
	// }
	//
	// if err := server.ListenAndServe(); err != nil {
	// 	log.Printf("Server err: %v", err)
	// }
}
