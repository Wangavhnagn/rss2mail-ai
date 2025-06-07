package summarizer

import (
    "bytes"
    "encoding/json"
    "net/http"
    "rss2mail-ai/fetcher"
)

type Config struct {
    AISummary struct {
        APIKey    string
        APIURL    string
        Model     string
        MaxTokens int
        Prompt    string
    }
}

func Summarize(cfg *Config, item fetcher.FeedItem) string {
    prompt := cfg.AISummary.Prompt
    if prompt == "" {
        prompt = "Summarize this article in 1-2 sentences."
    }

    reqBody := map[string]interface{}{
        "model": cfg.AISummary.Model,
        "messages": []map[string]string{
            {"role": "system", "content": prompt},
            {"role": "user", "content": item.Title + " - " + item.Link},
        },
        "max_tokens": cfg.AISummary.MaxTokens,
    }

    jsonBody, _ := json.Marshal(reqBody)
    req, _ := http.NewRequest("POST", cfg.AISummary.APIURL, bytes.NewBuffer(jsonBody))
    req.Header.Set("Authorization", "Bearer "+cfg.AISummary.APIKey)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil || resp.StatusCode != 200 {
        return ""
    }
    defer resp.Body.Close()

    var respData map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&respData)

    choices := respData["choices"].([]interface{})
    if len(choices) > 0 {
        msg := choices[0].(map[string]interface{})["message"].(map[string]interface{})["content"]
        return msg.(string)
    }
    return ""
}
