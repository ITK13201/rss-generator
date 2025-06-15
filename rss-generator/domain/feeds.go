package domain

import (
	"time"

	"github.com/ITK13201/rss-generator/ent"
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

func ConvertFeedItemFromModelToDomain(feedItem *ent.FeedItem) *FeedItem {
	return &FeedItem{
		Title:       feedItem.Title,
		Description: feedItem.Description,
		Link:        &feedItem.Link,
		PublishedAt: feedItem.PublishedAt,
	}
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

// Public API output for getting a feed by ID
type FeedGetOutput struct {
	FeedID int `json:"feed_id"`
}

func ConvertFeedFromModelToDomain(feed *ent.Feed) *Feed {
	items := []*FeedItem{}
	for _, model := range feed.Edges.FeedItems {
		items = append(items, ConvertFeedItemFromModelToDomain(model))
	}
	return &Feed{
		Title:       feed.Title,
		Description: feed.Description,
		Link:        feed.Link,
		PublishedAt: feed.PublishedAt,
		Items:       items,
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

// Private API output for getting a feed by site slug
