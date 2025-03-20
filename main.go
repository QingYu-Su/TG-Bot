package main

import (
	"fmt"
	"os"

	"github.com/QingYu-Su/TG-Bot/bot"
	"github.com/QingYu-Su/TG-Bot/config"
	"github.com/QingYu-Su/TG-Bot/http"
	"github.com/QingYu-Su/TG-Bot/log"
)

func main() {
	// 配置文件路径
	configPath := "./config.yaml"

	// 读取配置文件
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		fmt.Println("加载配置失败:", err)
		os.Exit(1)
	}

	// 初始化日志
	log.InitLogger(cfg)

	// 开启web服务，接收外部通知
	http_server := http.NewHTTPServer(cfg)
	http_server.Start()

	b, err := bot.NewBotServer(cfg, http_server)
	if err != nil {
		log.Error("机器人初始化失败")
		panic(err)
	}
	b.Start()
}
