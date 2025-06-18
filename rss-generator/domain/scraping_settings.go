package domain

type ScrapingSettingV2 struct {
	SiteSlug            string  `json:"site_slug"`
	Selector            string  `json:"selector"`
	InnerSelector       string  `json:"inner_selector"`
	TitleSelector       string  `json:"title_selector"`
	DescriptionSelector string  `json:"description_selector"`
	LinkSelector        *string `json:"link_selector"`
	CreatedAt           string  `json:"created_at"`
	UpdatedAt           string  `json:"updated_at"`
}

type SiteScrapingSettingOutputScrapingSetting struct {
	SiteSlug            string  `json:"site_slug"`
	Selector            string  `json:"selector"`
	InnerSelector       string  `json:"inner_selector"`
	TitleSelector       string  `json:"title_selector"`
	DescriptionSelector string  `json:"description_selector"`
	LinkSelector        *string `json:"link_selector"`
	CreatedAt           string  `json:"created_at"`
	UpdatedAt           string  `json:"updated_at"`
}
