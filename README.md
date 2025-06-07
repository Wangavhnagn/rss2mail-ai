# rss2mail-ai

⏰ 自动从多个 RSS 源抓取内容并通过邮件发送，支持去重与 AI 总结（可选）。

---

## 🚀 快速开始

### 1. 安装依赖并构建（Go 环境）

```bash
go mod tidy
go build -o rss2mail
./rss2mail
```

### 2. 或使用 Docker

```bash
docker build -t rss2mail .
docker run -v $(pwd)/config.yaml:/app/config.yaml rss2mail
```

---

## ⚙️ 配置指南（config.yaml）

以下是一个完整的 `config.yaml` 配置示例，可根据需要进行自定义：

```yaml
email:
  sender: "your@gmail.com"        # 发件邮箱地址
  password: "app_password"        # 邮箱授权码或密码
  smtp_host: "smtp.gmail.com"     # SMTP 服务器
  smtp_port: 465                  # SMTP 端口
  receivers:                      # 接收者邮箱列表（支持多个）
    - "target1@example.com"
    - "target2@example.com"

rss:
  feeds:
    - "https://example.com/rss"   # 支持多个 RSS 链接
    - "https://another.com/feed"
  fetch_interval_minutes: 60      # 每次抓取的间隔（分钟）
  enable_deduplication: true      # 是否去重（避免重复发送）

ai_summary:
  enabled: false                  # 是否启用 AI 总结（默认关闭）
  api_key: ""                     # OpenAI API Key（仅在启用时需要）
  api_url: "https://api.openai.com/v1/chat/completions"  # API 地址
  model: "gpt-3.5-turbo"          # 使用的模型
  max_tokens: 300
  prompt: "请总结这篇文章的关键内容，用简明扼要的中文描述。"                 # 最大生成长度
```

---

## 🧠 AI 总结说明（可选）

- `prompt` 字段可自定义提示词，决定 AI 总结风格与语言风格


- 默认关闭，如需启用请将 `enabled` 设置为 `true`
- 支持自定义 OpenAI API Key、模型、URL 等参数
- 总结内容将附加在每条 RSS 项后面

---

## 📧 邮件说明

邮件将按配置频率自动发送，内容包含：
- 每条 RSS 标题
- 链接
- 可选的 AI 总结内容

---

## 📦 编译为二进制

```bash
go build -o rss2mail
```

可编译为 `.exe` 或其他平台二进制文件后运行。

---

## 🐳 Docker 使用说明

构建镜像：

```bash
docker build -t rss2mail .
```

运行容器（挂载配置）：

```bash
docker run -v $(pwd)/config.yaml:/app/config.yaml rss2mail
```

---

## 🔒 建议

- 不要公开配置中的密码和 API 密钥
- 邮箱建议使用 App 授权码（如 Gmail）

---

MIT License · Created by You
