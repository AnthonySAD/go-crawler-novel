package crawler

import (
	"crawler/db"
	"crawler/httpclient"
	"log"
	"sync"
)

func Run(requests []Request){
	db.Init()
	for len(requests) > 0 {
		request := requests[0]
		requests = requests[1:]

		content, err := httpclient.Get(request.Url)

		if err != nil {
			log.Printf("Fetch error, Url: %s %v\n", request.Url, err)
			continue
		}

		result := request.ParseFunc(content, request.Info)

		requests = append(requests, result.Requests...)

		result.HandleFunc(result.Content)
	}

}

func Worker(requestChan chan Request, wg *sync.WaitGroup){

	for {
		request := <- requestChan
		wg.Add(1)
		content, err := httpclient.Get(request.Url)

		if err != nil {
			log.Printf("Fetch error, Url: %s %v\n", request.Url, err)
			return
		}

		result := request.ParseFunc(content, request.Info)

		result.HandleFunc(result.Content)

		for _, value := range result.Requests{
			requestChan <- value
		}

		wg.Done()
	}

}