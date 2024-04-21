package domain

type TestFeedCreateInput struct {
	SiteID    int               `json:"site_id"`
	Selectors ScrapingSelectors `json:"selectors"`
}

type TestFeedCreateOutput struct {
	FeedID int `json:"feed_id"`
}
