package mysqlDB

import (
	"database/sql"
	"fmt"
	setting "gindome/config"
	"time"

	mySql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	_ "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	dbConn *gorm.DB
	sqlDB  *sql.DB
)

/*
Init

@description: 初始化数据库连接

@param: cfg *setting.MySQLConfig 数据库配置.
*/
func Init(cfg *setting.MySQLConfig) {
	// 1. 拼接数据库连接字符串
	var dsn string = mysqlDsn(cfg)
	// 2. 打开数据库连接
	var err error
	sqlDB, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("sql.Open err, %v", err))
	}
	// 3. 设置数据库连接池参数
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns) // 最大连接数
	sqlDB.SetMaxOpenConns(cfg.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Hour)
	// 4. 创建 GORM 数据库连接
	gormDB, err := gorm.Open(mySql.New(mySql.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		SkipDefaultTransaction:                   false,
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用外键生成
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		// Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(fmt.Sprintf("链接数据库失败%s", err.Error()))
	}
	// 5. 保存 GORM 数据库连接
	dbConn = gormDB
}

/*
@description: 拼接数据库连接字符串

@param: cfg *setting.MySQLConfig 数据库配置

@return: string 数据库连接字符串.
*/
func mysqlDsn(cfg *setting.MySQLConfig) string {
	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s&timeout=%s&readTimeout=%s&writeTimeout=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB, cfg.Charset,
		cfg.ParseTime, cfg.Loc, cfg.Timeout, cfg.ReadTimeout, cfg.WriteTimeout)
	return dsn
}

/*
GetDB

@description: 获取数据库连接

@return: *gorm.DB 数据库连接.
*/
func GetDB() *gorm.DB {
	return dbConn
}

/*
CloseDB

@description: 关闭数据库连接

@return: error 错误信息.
*/
func CloseDB() error {
	return sqlDB.Close()
}
