package scraper

import (
	"encoding/json"
	"fmt"
	"github.com/ITK13201/rss-generator/domain"
	"github.com/PuerkitoBio/goquery"
	"github.com/sirupsen/logrus"
	"github.com/tebeka/selenium"
	"io"
	"net/http"
	"strings"
	"time"
)

type Scraper struct {
	cfg    *domain.Config
	logger *logrus.Logger
}

func NewScraper(cfg *domain.Config, logger *logrus.Logger) *Scraper {
	return &Scraper{
		cfg:    cfg,
		logger: logger,
	}
}

func (scraper *Scraper) fetchHTML(siteURL string) (*string, error) {
	res, err := http.Get(siteURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	htmlByte, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	html := string(htmlByte)
	return &html, nil
}

func (scraper *Scraper) fetchHTMLWithJS(siteURL string) (*string, error) {
	caps := selenium.Capabilities{"browserName": "chrome"}
	driver, err := selenium.NewRemote(caps, fmt.Sprintf(scraper.cfg.SeleniumServerURL))
	if err != nil {
		return nil, err
	}
	defer driver.Quit()

	err = driver.MaximizeWindow("")
	if err != nil {
		return nil, err
	}

	err = driver.Get(siteURL)
	if err != nil {
		return nil, err
	}
	html, err := driver.PageSource()
	if err != nil {
		return nil, err
	}
	return &html, nil
}

func (scraper *Scraper) formatString(str string) string {
	str = strings.Trim(str, "\n")
	str = strings.Trim(str, "\r")
	str = strings.TrimSpace(str)
	return str
}

func (scraper *Scraper) selectFeedObjects(siteURL string, html *string, scrapingSetting *domain.ScrapingSetting) (*domain.Feed, error) {
	now := time.Now()

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(*html))
	if err != nil {
		return nil, err
	}

	siteTitle := doc.Find("title").Text()
	siteDescription, _ := doc.Find("meta[name='description']").Attr("content")
	siteLink := siteURL

	selector := scrapingSetting.Selector
	innerSelector := scrapingSetting.InnerSelector
	titleSelector := scrapingSetting.TitleSelector
	descriptionSelector := scrapingSetting.DescriptionSelector
	linkSelector := scrapingSetting.LinkSelector

	selection := doc.Find(selector)

	feedItems := []*domain.FeedItem{}
	selection.Find(innerSelector).Each(func(i int, s *goquery.Selection) {
		var title string
		s.Find(titleSelector).Each(func(i int, s *goquery.Selection) {
			tmp := s.Clone()
			tmp.Find("*").Remove()
			title = tmp.Text()
		})
		var description string
		s.Find(descriptionSelector).Each(func(i int, s *goquery.Selection) {
			tmp := s.Clone()
			tmp.Find("*").Remove()
			description = tmp.Text()
		})

		var link *string
		if linkSelector != nil {
			linkStr, _ := s.Find(*linkSelector).Attr("href")
			link = &linkStr
		}

		formattedTitle := scraper.formatString(title)
		formattedDescription := scraper.formatString(description)

		feedItem := &domain.FeedItem{
			Title:       formattedTitle,
			Description: formattedDescription,
			Link:        link,
			PublishedAt: now,
		}
		feedItems = append(feedItems, feedItem)
	})

	feed := &domain.Feed{
		Title:       siteTitle,
		Description: siteDescription,
		Link:        siteLink,
		Items:       feedItems,
		PublishedAt: now,
	}

	return feed, nil
}

func (scraper *Scraper) FetchFeedElements(siteURL string, enableJSRendering bool, scrapingSetting *domain.ScrapingSetting) (*domain.Feed, error) {
	var html *string
	var err error
	if enableJSRendering {
		scraper.logger.Infof("fetching HTML with JS rendering from: %s", siteURL)
		html, err = scraper.fetchHTMLWithJS(siteURL)
	} else {
		scraper.logger.Infof("fetching HTML simply from: %s", siteURL)
		html, err = scraper.fetchHTML(siteURL)
	}
	if err != nil {
		return nil, err
	}
	scraper.logger.Infof("fetched HTML from: %s", siteURL)
	scraper.logger.Infof("selecting feed objects: %s", siteURL)
	feed, err := scraper.selectFeedObjects(siteURL, html, scrapingSetting)
	if err != nil {
		return nil, err
	}
	scraper.logger.Infof("selected feed objects: %s", siteURL)
	feedJson, _ := json.Marshal(feed)
	scraper.logger.Infof("Generated Latest RSS feed: %s", feedJson)
	return feed, nil
}
