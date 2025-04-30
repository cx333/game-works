package config

/**
 * @Author: wgl
 * @Description:
 * @File: config
 * @Version: 1.0.0
 * @Date: 2025/4/30 18:55
 */

import (
	"log"

	"github.com/spf13/viper"
)

// Config 是全局配置结构体
var Config *AppConfig

// AppConfig 定义yaml配置结构体
type AppConfig struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
	Log struct {
		Level string `yaml:"level"`
		Path  string `yaml:"path"`
	} `yaml:"log"`
	Database struct {
		Pgsql struct {
			Host     string `yaml:"host"`
			Port     int    `yaml:"port"`
			User     string `yaml:"user"`
			Password string `yaml:"password"`
			DbName   string `yaml:"dbname"`
			SslMode  string `yaml:"sslmode"`
		} `yaml:"pgsql"`
		Mysql struct {
			Host     string `yaml:"host"`
			Port     int    `yaml:"port"`
			User     string `yaml:"user"`
			Password string `yaml:"password"`
			DbName   string `yaml:"dbname"`
		} `yaml:"mysql"`
		Mongo struct {
			Uri    string `yaml:"uri"`
			DbName string `yaml:"dbname"`
		} `yaml:"mongo"`
	} `yaml:"database"`
	Redis struct {
		Addr     string `yaml:"addr"`
		Password string `yaml:"password"`
		Db       int    `yaml:"db"`
	} `yaml:"redis"`
	Jwt struct {
		Secret  string `yaml:"secret"`
		Expires int    `yaml:"expires"`
	} `yaml:"jwt"`
}

// InitConfig 初始化配置
func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}

	var c AppConfig
	if err := viper.Unmarshal(&c); err != nil {
		log.Fatalf("解析配置文件失败: %v", err)
	}
	Config = &c
}
