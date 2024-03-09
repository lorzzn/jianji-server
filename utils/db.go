package utils

import (
	"fmt"
	"jianji-server/config"
	"jianji-server/entity"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB gorm.DB

func AutoMigrateDB() {

	dbEntities := []any{
		&entity.User{},
		&entity.UserPassword{},
		&entity.ResponseLog{},
		&entity.UserToken{},
	}

	for _, dbEntity := range dbEntities {
		if err := DB.AutoMigrate(dbEntity); err != nil {
			log.Panicln("数据库自动迁移失败", err)
		}
	}
}

func SetupDB() {
	postgresConfig := config.Postgres

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		postgresConfig.Host,
		postgresConfig.Username,
		postgresConfig.Password,
		postgresConfig.DBName,
		postgresConfig.Port,
		postgresConfig.SSLMode,
		postgresConfig.TimeZone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，不自动加（s）
		},
	})

	if err != nil {
		log.Panicln("数据库连接失败", err)
	}

	DB = *db

	AutoMigrateDB()
}
