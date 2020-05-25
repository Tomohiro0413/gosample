package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {
	wait := new(sync.WaitGroup)
	urls := []string{
		"http://example.com",
		"http://example.net",
		"http://example.org",
	}
	// for _, url := range urls {
	// wait groupに追加
	// wait.Add(1)
	statusChan := make(chan string)
	for _, url := range urls {
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			fmt.Println(url, res.Status)
			wait.Done()
		}(url)
	}
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-statusChan)
	}

	// wait.Wait()
	// }
}
