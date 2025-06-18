package domain

type SitesGetAllOutput struct {
	Slug              string             `json:"slug"`
	Title             string             `json:"title"`
	Description       *string            `json:"description"`
	URL               string             `json:"url"`
	EnableJSRendering bool               `json:"enable_js_rendering"`
	CreatedAt         string             `json:"created_at"`
	UpdatedAt         string             `json:"updated_at"`
	ScrapingSetting   *ScrapingSettingV2 `json:"scraping_setting"`
}

type SitesGetBySlugOutput struct {
	Slug              string             `json:"slug"`
	Title             string             `json:"title"`
	Description       *string            `json:"description"`
	URL               string             `json:"url"`
	EnableJSRendering bool               `json:"enable_js_rendering"`
	CreatedAt         string             `json:"created_at"`
	UpdatedAt         string             `json:"updated_at"`
	ScrapingSetting   *ScrapingSettingV2 `json:"scraping_setting"`
}

type SitesCreateInput struct {
	Slug              string  `json:"slug"`
	Title             string  `json:"title"`
	Description       *string `json:"description"`
	URL               string  `json:"url"`
	EnableJSRendering bool    `json:"enable_js_rendering" default:"false"`
}

type SitesCreateOutput struct {
	Slug              string  `json:"slug"`
	Title             string  `json:"title"`
	Description       *string `json:"description"`
	URL               string  `json:"url"`
	EnableJSRendering bool    `json:"enable_js_rendering"`
	CreatedAt         string  `json:"created_at"`
	UpdatedAt         string  `json:"updated_at"`
}

type SitesUpdateInput struct {
	Title             *string `json:"title"`
	Description       *string `json:"description"`
	URL               *string `json:"url"`
	EnableJSRendering *bool   `json:"enable_js_rendering"`
}

type SitesUpdateOutput struct {
	Slug              string  `json:"slug"`
	Title             string  `json:"title"`
	Description       *string `json:"description"`
	URL               string  `json:"url"`
	EnableJSRendering bool    `json:"enable_js_rendering"`
	CreatedAt         string  `json:"created_at"`
	UpdatedAt         string  `json:"updated_at"`
}
