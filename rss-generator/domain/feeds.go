package domain

import (
	"github.com/ITK13201/rss-generator/ent"
	"time"
)

type ScrapingSetting struct {
	Selector            string  `json:"selector"`
	InnerSelector       string  `json:"inner_selector"`
	TitleSelector       string  `json:"title_selector"`
	DescriptionSelector string  `json:"description_selector"`
	LinkSelector        *string `json:"link_selector"`
}

type FeedItem struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Link        *string   `json:"link"`
	PublishedAt time.Time `json:"published_at"`
}

type Feed struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Items       []*FeedItem
	PublishedAt time.Time `json:"published_at"`
}

type FeedUpsertInput struct {
	ScrapingSetting ScrapingSetting `json:"scraping_setting"`
}

type FeedGetOutput struct {
	FeedID int `json:"feed_id"`
}

func ConvertFeedFromModelToDomain(feed *ent.Feed) *Feed {
	return &Feed{
		Title:       feed.Title,
		Description: feed.Description,
		Link:        feed.Link,
		PublishedAt: feed.PublishedAt,
	}
}

func ConvertScrapingSettingFromModelToDomain(scrapingSetting *ent.ScrapingSetting) *ScrapingSetting {
	return &ScrapingSetting{
		Selector:            scrapingSetting.Selector,
		InnerSelector:       scrapingSetting.InnerSelector,
		TitleSelector:       scrapingSetting.TitleSelector,
		DescriptionSelector: scrapingSetting.DescriptionSelector,
		LinkSelector:        &scrapingSetting.LinkSelector,
	}
}
