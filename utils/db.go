package utils

import (
	"fmt"
	"jianji-server/config"
	"jianji-server/entity"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/samber/lo"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB gorm.DB

type DBLogger struct {
	logger *zap.Logger
}

func (c *DBLogger) Write(p []byte) (n int, err error) {
	// parse
	// \d+.\d+[a-z]+   		--->  latency
	// /.*:\d+ 				--->  file
	// \[rows:(.*)?]		--->  rows
	// (?s)\\[rows:.*](.*)	--->  syntax (FindAllStringSubmatch)
	raw := string(p)
	lines := strings.Split(raw, "\n")
	latency := regexp.MustCompile("\\d+.\\d+[a-z]+").FindString(GetArrayItemByIndex(lines, 1, ""))
	stack := GetArrayItemByIndex(lines, 0, "")
	sql := strings.Join(lo.Drop(lines, 1), "\n")
	latency = strings.Trim(latency, " \n")
	stack = strings.Trim(stack, " \n")
	sql = strings.Trim(sql, " \n")

	c.logger.Info("Gorm log", zap.String("latency", latency), zap.String("stack", stack), zap.String("sql", sql))
	if regexp.MustCompile(`\D$`).MatchString(stack) {
		fmt.Printf("%sBAD SQL%s %s\n", logger.Red, logger.Reset, stack)
		fmt.Println(sql)
	}

	return len(p), nil
}

func AutoMigrateDB() {

	dbEntities := []any{
		&entity.User{},
		&entity.UserPassword{},
		&entity.ResponseLog{},
		&entity.UserToken{},
		&entity.Categories{},
		&entity.Tag{},
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
		Logger: logger.New(log.New(&DBLogger{
			logger: zap.New(zapcore.NewCore(LoggerEncoder, LoggerFileSyncer, LoggerLevelEnabler)),
		}, "", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Log level
				IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
				Colorful:                  false,       // Disable color
			},
		),
	})

	if err != nil {
		log.Panicln("数据库连接失败", err)
	}

	DB = *db

	AutoMigrateDB()
}
