package http

import (
	"fmt"
	"net/http"
	"os"

	"github.com/QingYu-Su/TG-Bot/config"
	"github.com/QingYu-Su/TG-Bot/log"
	"github.com/gin-gonic/gin"
)

type Message struct {
	Name    string
	Content map[string]interface{}
}

// HTTPServer 结构
type HTTPServer struct {
	engine *gin.Engine
	config *config.Config
	data   chan Message
}

// NewHTTPServer 初始化 HTTP 服务器
func NewHTTPServer(cfg *config.Config) *HTTPServer {
	// 禁用 Gin 默认日志
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = os.NewFile(0, os.DevNull)
	gin.DefaultErrorWriter = os.NewFile(0, os.DevNull)

	// 初始化服务器 & 数据通道
	server := &HTTPServer{
		engine: gin.New(),
		config: cfg,
		data:   make(chan Message, 100), // 缓冲区大小 100，避免阻塞
	}

	// 遍历 Receivers 注册路由
	for _, receiver := range cfg.Receivers {
		name := receiver.Name
		path := receiver.Path
		parts := receiver.Parts

		// 注册 POST 请求
		server.engine.POST(path, func(c *gin.Context) {
			var jsonData map[string]interface{}
			if err := c.ShouldBindJSON(&jsonData); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
				return
			}

			// 仅提取 parts 里指定的字段
			filteredData := make(map[string]interface{})
			for _, part := range parts {
				if value, exists := jsonData[part]; exists {
					filteredData[part] = value
				}
			}

			message := Message{name, filteredData}

			// 将数据写入通道（非阻塞）
			select {
			case server.data <- message:
				log.Infof("[%s] Received at [%s]: %+v\n", name, path, filteredData)
			default:
				log.Error("数据通道已满，丢弃数据！")
			}
		})

		log.Infof("Registered POST route: %s\n", path)
	}

	return server
}

// Start 运行 HTTP 服务器
func (s *HTTPServer) Start() {
	addr := fmt.Sprintf("0.0.0.0:%d", s.config.Port)
	log.Infof("Starting server on [%s]...\n", addr)
	go s.engine.Run(addr)
}

// GetDataChannel 获取数据
func (s *HTTPServer) GetData() Message {
	return <-s.data
}
