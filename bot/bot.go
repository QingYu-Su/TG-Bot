package bot

import (
	"context"
	"fmt"
	"strings"

	"github.com/QingYu-Su/TG-Bot/config"
	"github.com/QingYu-Su/TG-Bot/http"
	"github.com/QingYu-Su/TG-Bot/log"
	"github.com/go-telegram/bot" // 导入 Telegram Bot API 的 Go 实现
	"github.com/pkg/errors"
)

type BotServer struct {
	bot *bot.Bot
	cfg *config.Config
	hs  *http.HTTPServer
}

func NewBotServer(cfg *config.Config, hs *http.HTTPServer) (*BotServer, error) {
	s := &BotServer{
		cfg: cfg,
		hs:  hs,
	}

	b, err := bot.New(cfg.Token)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create bot")
		// 创建 Telegram Bot 实例，如果失败则返回错误
	}
	s.bot = b

	return s, nil
}

func (b *BotServer) Start() {
	for {
		message := b.hs.GetData()
		b.sendMessage(message)
	}
}

func (b *BotServer) sendMessage(message http.Message) {
	ctx := context.Background() // 创建一个默认的 Context

	// 格式化 message 为字符串
	msgText := b.formatMessage(message)

	// 遍历配置中的用户列表，发送消息
	for _, user := range b.cfg.User {
		_, err := b.bot.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: user,
			Text:   msgText,
		})
		if err != nil {
			log.Errorf("发送消息失败: %v\n", err)
		} else {
			log.Infof("成功发送消息给 %s\n", user)
		}
	}
}

func (b *BotServer) formatMessage(message http.Message) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("📩 收到新消息\n\n📝 Name: %s\n", message.Name))
	sb.WriteString("📜 内容:\n")

	for key, value := range message.Content {
		sb.WriteString(fmt.Sprintf("🔹 %s: %v\n", key, value))
	}

	return sb.String()
}
