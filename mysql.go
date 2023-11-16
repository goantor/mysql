package mysql

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type connector struct {
	opt IOption
}

func NewConnector(opt IOption) IMysql {
	return &connector{opt: opt}
}

// todo:  暂时其他配置用不到, 有时间再说
func (m connector) makeConfig() (config *gorm.Config) {
	config = &gorm.Config{}
	if m.opt.IsDebug() {
		config.Logger = logger.Default.LogMode(logger.LogLevel(m.opt.TakeLogMode()))
	}

	return
}

func (m connector) Connect() (db *gorm.DB) {
	var err error
	if db, err = gorm.Open(mysql.Open(m.opt.DataSourceName()), m.makeConfig()); err != nil {
		panic(fmt.Sprintf("connect mysql failed: %s\n", err.Error()))
	}

	if err = m.option(db); err != nil {
		panic(fmt.Sprintf("mysql option pool failed: %s\n", err.Error()))
	}

	return
}

func (m connector) option(db *gorm.DB) (err error) {
	var (
		pool *sql.DB
	)

	if pool, err = db.DB(); err != nil {
		return
	}

	pool.SetMaxIdleConns(m.opt.TakeMaxIdleConn())
	pool.SetMaxOpenConns(m.opt.TakeMaxOpenConn())
	pool.SetConnMaxLifetime(m.opt.TakeMaxLifeTime())

	if m.opt.IsDebug() {
		db.Debug()
	}

	return

}
