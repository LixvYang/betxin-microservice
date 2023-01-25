package repository

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitDB() {
	DbHost := viper.GetString("database.DbHost")
	DbPort := viper.GetString("database.DbPort")
	DbName := viper.GetString("database.DbName")
	DbUser := viper.GetString("database.DbUser")
	DbPassWord := viper.GetString("database.DbPassWord")

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DbUser,
		DbPassWord,
		DbHost,
		DbPort,
		DbName,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		// gorm日志模式：Warn
		Logger: logger.Default.LogMode(logger.Warn),
		// 外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务（提高运行速度）
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
	})
	if err != nil {
		log.Panic("连接数据库失败,请检查参数:", err)
	}

	// DB.AutoMigrate(
	// 	&User{},
	// )

	sqlDB, _ := DB.DB()
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	// SetMaxOpenCons 设置数据库的最大连接数量。
	// SetConnMaxLifetiment 设置连接的最大可复用时间
	sqlDB.SetMaxIdleConns(1000)
	sqlDB.SetMaxOpenConns(5000)
	sqlDB.SetConnMaxLifetime(time.Second * 30)
}
