package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type SqlConfig struct {
	Sql struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Database string `json:"database"`
		Charset  string `json:"charset"`
	} `json:"Sql"`
}

type IConfig interface {
	GetSqlByConfig(filePath string) string
}

type Config struct {
	SqlConfig
}

func NewConfig(sqlConfig SqlConfig) *Config {
	return &Config{SqlConfig: sqlConfig}
}

func (sqlConfig SqlConfig) GetSqlByConfig(filePath string) string {
	// 读取JSON文件内容
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return ""
	}

	// 将JSON数据解码为结构体
	var config SqlConfig
	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return ""
	}

	// 根据结构体信息生成SQL连接字符串
	sqlUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		config.Sql.Username, config.Sql.Password, config.Sql.Host, config.Sql.Port, config.Sql.Database, config.Sql.Charset)

	return sqlUrl
}
