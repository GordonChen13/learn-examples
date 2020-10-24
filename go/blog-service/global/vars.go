package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DBEngine *gorm.DB
	Logger   *zap.Logger
)
