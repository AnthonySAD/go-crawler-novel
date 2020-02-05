package main

import (
	"crawler/crawler"
	"crawler/db"
	models "crawler/websites/shubao69/models"
	shubao69 "crawler/websites/shubao69/parsers"
	"strconv"
	"sync"
	"time"
)

func main() {

	workerCount := 20

	DB := db.Db()
	defer DB.Close()
	var novel []models.Novel

	DB.Not("status", 1).Find(&novel)
	requestChan := make(chan crawler.Request)
	var wg sync.WaitGroup
	go func() {
		for _, value := range novel{
			requestChan <- crawler.Request{
				Url:       value.Url,
				ParseFunc: shubao69.ParseNovelDetail,
				Info: 	   map[string]string{
					"novelId":strconv.Itoa(int(value.ID)),
					"title":value.Title,
				},
			}
		}
	}()

	go func() {
		for i := 0; i < workerCount; i ++ {
			go crawler.Worker(requestChan, &wg)
		}
	}()

	time.Sleep(time.Minute)
	wg.Wait()
}

//func main() {
//	DB := db.Db()
//	defer DB.Close()
//	var novel []models.Novel
//
//	DB.Not("status", 1).Find(&novel)
//	var requests []crawler.Request
//	for _, value := range novel{
//		requests = append(requests, crawler.Request{
//			Url:       value.Url,
//			ParseFunc: shubao69.ParseNovelDetail,
//			Info: 	   map[string]string{
//				"novelId":strconv.Itoa(int(value.ID)),
//				"title":value.Title,
//			},
//		})
//	}
//
//	crawler.Run(requests)
//}
