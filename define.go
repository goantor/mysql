package mysql

import (
	"gorm.io/gorm"
	"time"
)

// IOption connector 配置
type IOption interface {
	DataSourceName() string
	IsDebug() bool
	TakeLogMode() int
	TakeMaxIdleConn() int
	TakeMaxOpenConn() int
	TakeMaxLifeTime() time.Duration
}

// IMysql connector 接口
type IMysql interface {
	Connect() *gorm.DB
}
