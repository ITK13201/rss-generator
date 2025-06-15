package domain

type SiteScrapingSettingOutput struct {
	Slug              string                                   `json:"slug"`
	Title             string                                   `json:"title"`
	Description       *string                                  `json:"description"`
	URL               string                                   `json:"url"`
	EnableJSRendering bool                                     `json:"enable_js_rendering"`
	ScrapingSetting   SiteScrapingSettingOutputScrapingSetting `json:"scraping_setting"`
}

type SiteScrapingSettingOutputScrapingSetting struct {
	Selector            string  `json:"selector"`
	InnerSelector       string  `json:"inner_selector"`
	TitleSelector       string  `json:"title_selector"`
	DescriptionSelector string  `json:"description_selector"`
	LinkSelector        *string `json:"link_selector"`
	CreatedAt           string  `json:"created_at"`
	UpdatedAt           string  `json:"updated_at"`
}
