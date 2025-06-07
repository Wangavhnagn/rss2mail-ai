package main

import (
    "log"
    "time"
    "rss2mail-ai/fetcher"
    "rss2mail-ai/mailer"
    "rss2mail-ai/summarizer"
    "gopkg.in/yaml.v2"
    "os"
    "io/ioutil"
)

type Config struct {
    Email struct {
        Sender    string   `yaml:"sender"`
        Password  string   `yaml:"password"`
        SMTPHost  string   `yaml:"smtp_host"`
        SMTPPort  int      `yaml:"smtp_port"`
        Receivers []string `yaml:"receivers"`
    } `yaml:"email"`
    RSS struct {
        Feeds                []string `yaml:"feeds"`
        FetchIntervalMinutes int      `yaml:"fetch_interval_minutes"`
        EnableDeduplication  bool     `yaml:"enable_deduplication"`
    } `yaml:"rss"`
    AISummary struct {
        Enabled   bool   `yaml:"enabled"`
        APIKey    string `yaml:"api_key"`
        APIURL    string `yaml:"api_url"`
        Model     string `yaml:"model"`
        MaxTokens int    `yaml:"max_tokens"`
    } `yaml:"ai_summary"`
}

func loadConfig(path string) (*Config, error) {
    data, err := ioutil.ReadFile(path)
    if err != nil {
        return nil, err
    }
    var cfg Config
    if err := yaml.Unmarshal(data, &cfg); err != nil {
        return nil, err
    }
    return &cfg, nil
}

func main() {
    cfg, err := loadConfig("config.yaml")
    if err != nil {
        log.Fatal("Error loading config:", err)
    }

    dedup := make(map[string]bool)

    for {
        contents := fetcher.FetchFeeds(cfg.RSS.Feeds, cfg.RSS.EnableDeduplication, dedup)

        if cfg.AISummary.Enabled {
            for i, item := range contents {
                summary := summarizer.Summarize(cfg, item)
                contents[i].Summary = summary
            }
        }

        mailer.SendEmail(cfg, contents)
        log.Println("Sleeping for", cfg.RSS.FetchIntervalMinutes, "minutes...")
        time.Sleep(time.Duration(cfg.RSS.FetchIntervalMinutes) * time.Minute)
    }
}
