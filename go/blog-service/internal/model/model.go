package model

import (
	"fmt"

	"github.com/GordonChen13/learn-examples/go/blog-service/global"
	"github.com/GordonChen13/learn-examples/go/blog-service/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Model struct {
	*gorm.Model
	CreatedBy string `gorm:"type:VARCHAR(32) DEFAULT '' COMMENT '创建人'"`
	UpdatedBy string `gorm:"type:VARCHAR(32) DEFAULT '' COMMENT '更新人'"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if global.ServerSetting.RunMode == "debug" {
		db = db.Debug()
	}
	sqlDb, err := db.DB()
	if err != nil {
		return db, err
	}
	sqlDb.SetMaxIdleConns(databaseSetting.MaxIdleConns)
	sqlDb.SetMaxOpenConns(databaseSetting.MaxOpenConns)

	return db, nil
}
