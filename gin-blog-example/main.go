package main

import (
	"fmt"
	_ "gin-blog-example/docs"
	"gin-blog-example/pkg/setting"
	"gin-blog-example/routers"
	"github.com/fvbock/endless"
	"go.uber.org/zap"
	"syscall"
)

//func main() {
//	router := routers.InitRouter()
//
//	s := &http.Server{
//		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
//		Handler:        router,
//		ReadTimeout:    setting.ReadTimeout,
//		WriteTimeout:   setting.WriteTimeout,
//		MaxHeaderBytes: 1 << 20,
//	}
//
//	err := s.ListenAndServe()
//	if err != nil {
//		return
//	}
//}

func init() {
	replaceLogger()
}

func replaceLogger() {
	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)
}

// @title gin-blog API
// @version 1.0
// @description gin-blog的示例项目
func main() {
	endless.DefaultReadTimeOut = setting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.HTTPPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		zap.L().Info("Actual pid is %d", zap.Int("pid", syscall.Getpid()))
	}

	err := server.ListenAndServe()
	if err != nil {
		zap.S().Infof("Server err: %v", err)
	}
}
