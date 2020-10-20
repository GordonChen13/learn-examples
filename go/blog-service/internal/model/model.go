package model

import (
	"fmt"

	"github.com/GordonChen13/learn-examples/go/blog-service/global"
	"github.com/GordonChen13/learn-examples/go/blog-service/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Model struct {
	Id         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
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
