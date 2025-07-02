package config

import (
	"encoding/json"
	"os"
)

// Config 应用程序配置
type Config struct {
	Server   ServerConfig   `json:"server"`
	Database DatabaseConfig `json:"database"`
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Address string `json:"address"`
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Driver   string `json:"driver"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
}

// LoadConfig 从配置文件加载配置
func LoadConfig() (*Config, error) {
	// 默认配置
	config := &Config{
		Server: ServerConfig{
			Address: ":8080",
		},
		Database: DatabaseConfig{
			Driver:   "mysql",
			Host:     "localhost",
			Port:     3306,
			User:     "root",
			Password: "password",
			DBName:   "recipe_app",
		},
	}

	// 尝试从文件加载配置
	file, err := os.Open("config.json")
	if err == nil {
		defer file.Close()
		decoder := json.NewDecoder(file)
		if err := decoder.Decode(config); err != nil {
			return nil, err
		}
	}

	return config, nil
}
