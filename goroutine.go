package main

import (
	"fmt"
	"log"
	"net/http"
)

// func main() {
// 	// wait := new(sync.WaitGroup)
// 	urls := []string{
// 		"http://example.com",
// 		"http://example.net",
// 		"http://example.org",
// 	}
// 	// for _, url := range urls {
// 	// wait groupに追加
// 	// wait.Add(1)
// 	statusChan := make(chan string)
// 	for _, url := range urls {
// 		go func(url string) {
// 			res, err := http.Get(url)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			defer res.Body.Close()
// 			// fmt.Println(url, res.Status)
// 			// wait.Done()
// 			statusChan <- res.Status
// 		}(url)
// 	}
// 	for i := 0; i < len(urls); i++ {
// 		fmt.Println(<-statusChan)
// 	}
// 	// wait.Wait()
// 	// }
// }

func getStatus(urls []string) <-chan string {
	//getstatusないで結果を渡すためのstatusChanを生成
	//非同期で行う処理を匿名関数としてリクエストをそれぞれ別々のごルーチンで実行する
	statusChan := make(chan string)
	for _, url := range urls {
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			statusChan <- res.Status
		}(url)
	}
	return statusChan
}

func main() {
	urls := []string{
		"http://example.com",
		"http://example.net",
		"http://example.org",
	}
	statusChan := getStatus(urls)
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-statusChan)
	}
}
