package shubao69

import (
	"bytes"
	"crawler/crawler"
	handlers "crawler/websites/shubao69/handlers"
	models "crawler/websites/shubao69/models"
	"github.com/PuerkitoBio/goquery"
	"log"
)

func ParseTitle(content []byte, _ map[string]string) crawler.ParseResult{
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(content))
	if err != nil {
		log.Fatal(err)
	}

	var Content []*models.Novel
	var Requests []crawler.Request

	doc.Find(".cover p").Each(func(i int, s *goquery.Selection) {
		a := s.Find("a")
		url, exists := a.Eq(1).Attr("href")
		if !exists {
			return
		}
		tag := a.Eq(0).Text()
		title := a.Eq(1).Text()
		author := a.Eq(2).Text()
		novel := models.Novel{
			Title:         title,
			AuthorName:    author,
			Url:           "http://XXXXX" + url,
			Tag:           tag,
		}
		Content = append(Content, &novel)
	})

	var url string
	doc.Find(".page a").Each(func(i int, s *goquery.Selection) {
		if s.Text() == "下页"{
			url, _ = s.Attr("href")
		}
	})

	if url != "" {
		Requests = []crawler.Request{
			crawler.Request{
				Url:       "http://XXXX" + url,
				ParseFunc: ParseTitle,
			},
		}
	}

	return crawler.ParseResult{
		Requests:   Requests,
		HandleFunc: handlers.HandleNovel,
		Content:    Content,
	}
}
