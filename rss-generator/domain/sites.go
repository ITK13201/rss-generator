package domain

type SitesCreateInput struct {
	Slug        string  `json:"slug"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	URL         string  `json:"url"`
}

type SitesCreateOutput struct {
	Slug  string `json:"slug"`
	Title string `json:"title"`
}
