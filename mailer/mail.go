package mailer

import (
    "crypto/tls"
    "fmt"
    "log"
    "net/smtp"
    "strings"
    "rss2mail-ai/fetcher"
)

func SendEmail(cfg *Config, items []fetcher.FeedItem) {
    body := ""
    for _, item := range items {
        body += fmt.Sprintf("• %s
%s

", item.Title, item.Link)
        if item.Summary != "" {
            body += fmt.Sprintf("Summary: %s

", item.Summary)
        }
    }

    auth := smtp.PlainAuth("", cfg.Email.Sender, cfg.Email.Password, cfg.Email.SMTPHost)

    msg := "From: " + cfg.Email.Sender + "
" +
        "To: " + strings.Join(cfg.Email.Receivers, ",") + "
" +
        "Subject: RSS Digest

" + body

    conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", cfg.Email.SMTPHost, cfg.Email.SMTPPort), nil)
    if err != nil {
        log.Println("Dial error:", err)
        return
    }
    c, err := smtp.NewClient(conn, cfg.Email.SMTPHost)
    if err != nil {
        log.Println("Client error:", err)
        return
    }
    defer c.Quit()
    c.Auth(auth)
    c.Mail(cfg.Email.Sender)
    for _, to := range cfg.Email.Receivers {
        c.Rcpt(to)
    }

    w, _ := c.Data()
    _, err = w.Write([]byte(msg))
    if err != nil {
        log.Println("Write error:", err)
    }
    w.Close()
}
