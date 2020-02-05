package shubao69

import (
	"bytes"
	"crawler/crawler"
	"crawler/db"
	"crawler/httpclient"
	handlers "crawler/websites/shubao69/handlers"
	models "crawler/websites/shubao69/models"
	"github.com/PuerkitoBio/goquery"
	"log"
)

func ParseNovelDetail(content []byte, info map[string]string) crawler.ParseResult{
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(content))
	if err != nil {
		log.Fatal(err)
	}

	var Requests []crawler.Request

	a := doc.Find(".ablum_read a").Eq(1)
	url, _ := a.Attr("href")

	filePath := "download/" + info["novelId"] + "." + info["title"] + ".txt"
	err2 := httpclient.Download(url, filePath)

	if err2 == nil{
		DB := db.Db()
		var novel models.Novel
		DB.First(&novel, info["novelId"])
		DB.Model(&novel).Update("status", 1)
	}

	return crawler.ParseResult{
		Requests:   Requests,
		HandleFunc: handlers.HandleNovelDetail,
		Content:    url,
	}
}

