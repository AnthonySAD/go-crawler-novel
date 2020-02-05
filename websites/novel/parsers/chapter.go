package shubao69

import (
	"bytes"
	"crawler/crawler"
	handlers "crawler/websites/shubao69/handlers"
	models "crawler/websites/shubao69/models"
	"github.com/PuerkitoBio/goquery"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func ParseChapter(content []byte, info map[string]string) crawler.ParseResult{
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(content))
	if err != nil {
		log.Fatal(err)
	}

	var Content *models.Chapter
	var Requests []crawler.Request

	if info["title"] == ""{
		chapterTitle := doc.Find("#nr_title").Text()
		chapterTitle = strings.TrimSpace(chapterTitle)
		titleRe := regexp.MustCompile("\\(\u7b2c\\d+/\\d+\u9875\\)")
		chapterTitle = titleRe.ReplaceAllString(chapterTitle, "")
		info["title"] = chapterTitle
	}


	chapterContent := doc.Find("#nr1").Text()
	chapterContent = strings.ReplaceAll(chapterContent, "    ", "\n    ")
	ignoreRe := regexp.MustCompile("-->>\\(\u7b2c\\d+/\\d+\u9875\\)\uff08\u672c\u7ae0\u672a\u5b8c\uff0c\u8bf7\u70b9\u51fb\u4e0b\u4e00\u9875\u7ee7\u7eed\u9605\u8bfb\uff09$")
	chapterContent = ignoreRe.ReplaceAllString(chapterContent, "")
	info["content"] += chapterContent

	nextPage := doc.Find("#pb_next")

	url, _ := nextPage.Attr("href")
	if strings.Count(url, "/") == 2{
		url = "http://XXXXX" + url
		Requests = append(Requests, crawler.Request{
			Url:       url,
			ParseFunc: ParseChapter,
			Info:      info,
		})
	}else{
		url = ""
	}

	if nextPage.Text() == "下一章" {
		id, _ := strconv.Atoi(info["novelId"])
		order, _ := strconv.Atoi(info["order"])
		order += 1
		Content = &models.Chapter{
			NovelID:    uint(id),
			Order:		uint(order),
			Title:      info["title"],
			Content:    info["content"],
			WordsTotal: len(info["content"]) / 3,
			NextUrl:    url,
		}
		info["title"] = ""
		info["content"] = ""
		info["order"] = strconv.Itoa(order)
	}

	return crawler.ParseResult{
		Requests:   Requests,
		HandleFunc: handlers.HandleChapter,
		Content:    Content,
	}
}
