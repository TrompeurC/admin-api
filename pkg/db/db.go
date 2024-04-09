package db

import (
	"admin-api/common/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var Db *gorm.DB

// 数据库初始化
func SetupDBLink() error {
	dbConfig := config.Config.DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Db, dbConfig.Charset)
	Db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		// 在 AutoMigrate 或 CreateTable 时，GORM 会自动创建外键约束，若要禁用该特性，可将其设置为 true
		DisableForeignKeyConstraintWhenMigrating: false,
	})

	if Db.Error != nil {
		panic(Db.Error)
	}

	sqlDb, _ := Db.DB()
	sqlDb.SetMaxIdleConns(dbConfig.MaxIdle)
	sqlDb.SetMaxIdleConns(dbConfig.MinOpen)

	return nil
}
