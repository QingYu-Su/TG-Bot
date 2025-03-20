# 项目说明

1个简单的转发web post消息到TG机器人的框架，功能简陋，后续有时间会重构并且增加功能

## 部署
1. 创建目录，假设创建为`~/tgbot`
2. 进入tgbot目录，执行以下命令
    ```
    git clone https://github.com/QingYu-Su/TG-Bot.git
    ```
3. 复制`~/tgbot/TG-Bot/config.yaml`文件到tgbot目录中
4. 按需修改`config.yaml`
   1. TOKEN：TG机器人token（你想用哪个机器人）
   2. USER：TG用户ID（你想发给哪个用户）
   3. LOG_LEVEL：日志等级，info打印全部，error只打印错误日志，disabled不打印日志
   4. PORT：web服务要监听哪个端口
   5. RECEIVERS：web服务的接收者
      1. NAME：名称
      2. PATH：要用哪个url路径来接收post请求
      3. PARTS：要接收post请求数据中的哪个字段
5. 在tgbot目录创建`docker-compose.yml`文件，写入以下内容
    ```
    services:
        tgbot:
            build: TG-Bot
            container_name: tgbot
            volumes:
                - ./config.yaml:/app/config.yaml
            ports:
                - "9999:9999"
    ```
6. 执行`docker compose up -d`

# 参考资料
> [memogram](https://github.com/usememos/telegram-integration)
