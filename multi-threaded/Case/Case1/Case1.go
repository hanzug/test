package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type urlStatus struct {
	Status   int
	Duration time.Duration
}

func UrlStatusCheck() {
	concurrency := 5
	links := []string{
		"https://www.baidu.com",
		"https://www.yahoo.com",
		"https://www.amazon.com",
		"https://www.wikipedia.org",
		"https://www.qq.com",
		"https://www.live.com",
		"https://www.taobao.com",
		"https://www.bing.com",
		"https://www.instagram.com",
		"https://www.weibo.com",
		"https://www.sina.com.cn",
		"https://www.linkedin.com",
		"https://www.yahoo.co.jp",
		"https://www.msn.com",
		"https://www.vk.com",
		"https://www.google.de",
	}

	urlStatusCheck(links, concurrency)
}

func urlStatusCheck(links []string, concurrency int) {
	var wg sync.WaitGroup
	results := make(map[string]urlStatus)
	var mutex sync.Mutex
	semaphora := make(chan struct{}, concurrency)

	for _, link := range links {
		wg.Add(1)
		semaphora <- struct{}{}

		go func(url string) {
			defer wg.Done()
			start := time.Now()
			resp, err := http.Get(url)
			duration := time.Since(start)

			if err != nil {
				fmt.Printf("Error fetching %s: %s\n", url, err)
				mutex.Lock()
				results[url] = urlStatus{Status: 0, Duration: duration}
				mutex.Unlock()

				<-semaphora
				return
			}
			resp.Body.Close()

			mutex.Lock()
			results[url] = urlStatus{Status: resp.StatusCode, Duration: duration}
			mutex.Unlock()

			<-semaphora
		}(link)
	}
	wg.Wait()

	for url, status := range results {
		fmt.Printf("%s: %+v\n", url, status)
	}
}

func main() {
	UrlStatusCheck()
}
