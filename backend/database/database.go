package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"recipe-app/backend/config"
)

// DB 全局数据库连接
var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB(cfg config.DatabaseConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return nil, err
	}

	// 尝试读取SQL初始化脚本
	sqlBytes, err := os.ReadFile("init_data.sql")
	if err != nil {
		log.Printf("Warning: Could not read init_data.sql: %v", err)
	} else {
		// 执行SQL脚本
		sqlStatements := strings.Split(string(sqlBytes), ";")
		for _, stmt := range sqlStatements {
			stmt = strings.TrimSpace(stmt)
			if stmt == "" {
				continue
			}
			if err := db.Exec(stmt).Error; err != nil {
				log.Printf("Warning: Error executing SQL statement: %v", err)
			}
		}
	}

	DB = db

	// 获取底层SQL数据库连接以便后续关闭
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	return sqlDB, nil
}
