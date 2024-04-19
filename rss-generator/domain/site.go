package domain

type SiteCreateInput struct {
	Slug        string  `json:"slug"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	URL         string  `json:"url"`
}

type SiteCreateOutput struct {
	Slug  string `json:"slug"`
	Title string `json:"title"`
}
