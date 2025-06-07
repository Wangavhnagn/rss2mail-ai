package fetcher

import (
    "github.com/mmcdole/gofeed"
)

type FeedItem struct {
    Title   string
    Link    string
    Summary string
}

func FetchFeeds(urls []string, dedupEnabled bool, seen map[string]bool) []FeedItem {
    parser := gofeed.NewParser()
    var results []FeedItem

    for _, url := range urls {
        feed, err := parser.ParseURL(url)
        if err != nil {
            continue
        }
        for _, item := range feed.Items {
            if dedupEnabled && seen[item.Link] {
                continue
            }
            seen[item.Link] = true
            results = append(results, FeedItem{
                Title: item.Title,
                Link:  item.Link,
            })
        }
    }
    return results
}
