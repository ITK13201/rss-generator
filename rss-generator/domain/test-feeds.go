package domain

type TestFeedCreateInput struct {
	ScrapingSetting ScrapingSetting `json:"scraping_setting"`
}

type TestFeedCreateOutput struct {
	FeedID int `json:"feed_id"`
}
