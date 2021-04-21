package model

import (
	"fmt"
	"gin-app/config"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

// gorm结构体，可以嵌入到其它模型结构体中
type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// 根据数据库类型，初始化对应的链接
func InitDB() {
	switch config.Bases.App.Dbtype {
	case "mysql":
		GormMysql()
	default:
		GormMysql()
	}
}

// gorm mysql 连接
func GormMysql() {
	var err error

	mysqlConfig := config.Bases.Mysql

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		mysqlConfig.Username,
		mysqlConfig.Password,
		mysqlConfig.Host,
		mysqlConfig.Port,
		mysqlConfig.Dbname)

	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "tb_", // 表前缀
			SingularTable: true, // 禁用表复数
		},
	})

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(10)
}

func GetDB() *gorm.DB {
	return db
}
