package config

import (
	"encoding/json"
	"log"
	"os"
)

// Database 数据库配置对象
type Database struct {
	Type        string `json:"type"`
	User        string `json:"user"`
	Password    string `json:"password"`
	Host        string `json:"host"`
	Port        string `json:"port"`
	Name        string `json:"name"`
	TablePrefix string `json:"table_prefix"`
}

// Config 配置对象
type Config struct {
	Database *Database `json:"database"`
}

var (
	// DatabaseSetting 数据库配置对象实例.
	DatabaseSetting = &Database{}

	// GlobalConfigSetting 配置实例.
	GlobalConfigSetting = &Config{}
)

// GetConfig 读取配置
func GetConfig() {
	confFile := "config.json"
	filePtr, err := os.Open(confFile) // config的文件目录

	if err != nil {
		log.Fatalf("open config file from '%s' is error:%s\n", confFile, err.Error())
		return
	}
	defer filePtr.Close()
	// 创建json解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(GlobalConfigSetting)
	DatabaseSetting = GlobalConfigSetting.Database

	if err != nil {
		log.Fatalf("decode config file error:%s\n", err.Error())
	}
}
