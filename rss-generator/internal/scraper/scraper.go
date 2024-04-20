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
)

type Util struct {
	cfg    *domain.Config
	logger *logrus.Logger
}

func NewUtil(cfg *domain.Config, logger *logrus.Logger) *Util {
	return &Util{
		cfg:    cfg,
		logger: logger,
	}
}

func (u *Util) fetchHTML(siteURL string) (*string, error) {
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

func (u *Util) fetchHTMLWithJS(siteURL string) (*string, error) {
	caps := selenium.Capabilities{"browserName": "chrome"}
	driver, err := selenium.NewRemote(caps, fmt.Sprintf(u.cfg.SeleniumServerURL))
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

func (u *Util) formatString(str string) string {
	str = strings.Trim(str, "\n")
	str = strings.Trim(str, "\r")
	str = strings.TrimSpace(str)
	return str
}

func (u *Util) selectFeedObjects(siteURL string, html *string) (*domain.LatestFeed, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(*html))
	if err != nil {
		return nil, err
	}

	siteTitle := doc.Find("title").Text()
	siteDescription, _ := doc.Find("meta[name='description']").Attr("content")
	siteLink := siteURL

	selector := "#container > main > div.in_news.fade_b.isPlay_b > div > ul"
	innerSelector := "li"
	titleSelector := "a > dl > dt"
	descriptionSelector := "a > dl > dd"
	linkSelector := "a"

	selection := doc.Find(selector)

	feedItems := []*domain.LatestFeedItem{}
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
		link, _ := s.Find(linkSelector).Attr("href")

		feedItem := &domain.LatestFeedItem{
			Title:       u.formatString(title),
			Description: u.formatString(description),
			Link:        link,
		}
		feedItems = append(feedItems, feedItem)
	})

	feed := &domain.LatestFeed{
		Title:       siteTitle,
		Description: siteDescription,
		Link:        siteLink,
		Items:       feedItems,
	}

	return feed, nil
}

func (u *Util) FetchFeedElements(siteURL string, enableJSRendering bool) (*domain.LatestFeed, error) {
	var html *string
	var err error
	if enableJSRendering {
		u.logger.Infof("fetching HTML with JS rendering from: %s", siteURL)
		html, err = u.fetchHTMLWithJS(siteURL)
	} else {
		u.logger.Infof("fetching HTML simply from: %s", siteURL)
		html, err = u.fetchHTML(siteURL)
	}
	if err != nil {
		return nil, err
	}
	u.logger.Infof("fetched HTML from: %s", siteURL)
	u.logger.Infof("selecting feed objects: %s", siteURL)
	feed, err := u.selectFeedObjects(siteURL, html)
	if err != nil {
		return nil, err
	}
	u.logger.Infof("selected feed objects: %s", siteURL)
	feedJson, _ := json.Marshal(feed)
	u.logger.Infof("Generated Latest RSS feed: %s", feedJson)
	return feed, nil
}
