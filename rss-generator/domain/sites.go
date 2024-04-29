package domain

type SitesCreateInput struct {
	Slug              string  `json:"slug"`
	Title             string  `json:"title"`
	Description       *string `json:"description"`
	URL               string  `json:"url"`
	EnableJSRendering bool    `json:"enable_js_rendering" default:"false"`
}

type SitesCreateOutput struct {
	Slug  string `json:"slug"`
	Title string `json:"title"`
}

type SitesUpdateInput struct {
	Title             *string `json:"title"`
	Description       *string `json:"description"`
	URL               *string `json:"url"`
	EnableJSRendering *bool   `json:"enable_js_rendering"`
}
