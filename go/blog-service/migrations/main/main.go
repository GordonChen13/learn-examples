package main

import (
	"log"
	"time"

	"github.com/GordonChen13/learn-examples/go/blog-service/global"
	"github.com/GordonChen13/learn-examples/go/blog-service/internal/model"
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
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := global.DBEngine.AutoMigrate(&model.Article{}, &model.Tag{}, &model.ArticleTag{})
	if err != nil {
		panic(err)
	}
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
