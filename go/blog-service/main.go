package main

import (
	"log"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/GordonChen13/learn-examples/go/blog-service/internal/model"

	"github.com/gin-gonic/gin"

	"github.com/GordonChen13/learn-examples/go/blog-service/global"
	"github.com/GordonChen13/learn-examples/go/blog-service/internal/routers"
	"github.com/GordonChen13/learn-examples/go/blog-service/pkg/setting"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init setupSetting err: %s", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init setupDBEngine err: %s", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init setupDBEngine err: %s", err)
	}
}

// @title 博客系统
// @version 1.0
// @description Go learning blog
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeOut,
		WriteTimeout:   global.ServerSetting.WriteTimeOut,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func setupSetting() error {
	settings, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = settings.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = settings.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = settings.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeOut *= time.Second
	global.ServerSetting.WriteTimeOut *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	var err error
	if global.ServerSetting.RunMode == "debug" {
		global.Logger, err = zap.NewDevelopment()
	} else {
		global.Logger, err = zap.NewProduction()
	}
	if err != nil {
		return err
	}
	return nil
}
