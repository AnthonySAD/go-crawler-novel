# 简介

该项目是小说爬虫，爬取小说网站的所有小说及小说章节，并下载所有小说，因为尝试爬取的是盗版网站，所以代码中，我把网址隐去了。
该项目是我第一次用glang写项目，也是第一次写爬虫。该项目没有经过优化，目前仅仅是简约版爬虫。

# 功能介绍

实现了高并发异步爬取，但是没有很好的监控并处理请求情况，及goroutine的运行情况。

该项目的总体构架是，httpclient复杂爬取网页，parser负责分析页面数据，handler负责储存数据及统计工作。crawler负责调用这3个模块以实现功能。

# 运行

```
git clone github.com/AnthonySAD/go-crawler-novel
cd go-crawler-novel
go run main.go
```

