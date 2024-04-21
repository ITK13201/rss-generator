package domain

type FeedCreateInput struct {
	SiteID int `json:"site_id"`
}

type FeedCreateOutput struct {
	FeedID int `json:"feed_id"`
}

type LatestFeedItem struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
}

type LatestFeed struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Items       []*LatestFeedItem
}
