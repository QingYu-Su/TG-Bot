package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Receiver struct {
	Name  string   `yaml:"NAME"`
	Path  string   `yaml:"PATH"`
	Parts []string `yaml:"PARTS"`
}

type Config struct {
	Token     string     `yaml:"TOKEN"`
	User      []string   `yaml:"USER"`
	LogLevel  string     `yaml:"LOG_LEVEL"`
	Port      int        `yaml:"PORT"`
	Receivers []Receiver `yaml:"RECEIVERS"`
}

// LoadConfig 读取并解析 YAML 配置文件
func LoadConfig(path string) (*Config, error) {
	// 读取文件内容
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("无法读取配置文件: %w", err)
	}

	// 解析 YAML 数据
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("解析 YAML 失败: %w", err)
	}

	return &config, nil
}

// Print 打印所有配置信息
func (c *Config) Print() {
	fmt.Println("配置详情：")
	fmt.Println("TOKEN:", c.Token)
	fmt.Println("用户列表:", c.User)
	fmt.Println("日志级别:", c.LogLevel)
	fmt.Println("端口:", c.Port)
	fmt.Println("接收者列表:")
	for _, receiver := range c.Receivers {
		fmt.Printf("  - 名称: %s, 路径: %s, 解析项: %v\n", receiver.Name, receiver.Path, receiver.Parts)
	}
}
