package mysqlDB

import (
	"database/sql"
	"fmt"
	setting "gindome/config"
	"time"

	mySql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var dbConn *gorm.DB

func Init(cfg *setting.MySQLConfig) {
	var dsn string = mysqlDsn(cfg)
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("sql.Open err, %v", err))
	}
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns) //最大连接数
	sqlDB.SetMaxOpenConns(cfg.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Hour)

	gormDB, err := gorm.Open(mySql.New(mySql.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		SkipDefaultTransaction:                   false,
		DisableForeignKeyConstraintWhenMigrating: true, //禁用外键生成
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(fmt.Sprintf("链接数据库失败%s", err.Error()))
	}
	dbConn = gormDB
}

func mysqlDsn(cfg *setting.MySQLConfig) string {
	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s&timeout=%s&readTimeout=%s&writeTimeout=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB, cfg.Charset,
		cfg.ParseTime, cfg.Loc, cfg.Timeout, cfg.ReadTimeout, cfg.WriteTimeout)
	return dsn
}

func GetDB() *gorm.DB {
	return dbConn.Debug()
}
