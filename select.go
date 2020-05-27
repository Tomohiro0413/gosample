package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

//  ch1 := make (chan string)
//  ch2 := make (chan string)

//  for {
// 	 select {
// 	 case c1 := <- ch1:
// 		//ch1からデータを読みだした時に発動する
// 	 case c2 := <- ch2:
// 		 //ch1からデータを読みだした時に発動する
// 	 case ch2 <- "c":
// 		//ch2にデータを書き込んだ時に発動する
// 	 default:
// 		//実行されなかった時に発動する
// 	 }
//  }

 func main(){
	 timeout := time.After(time.Second)
	 urls := []string{
		"http://example.com",
        "http://example.net",
        "http://example.org",
	 }
	 statusChan := getStatus(urls)


 LOOP:
 	for {
		 select {
		 case status := <-statusChan:
			fmt.Println(status)
		 case <- timeout:
			break LOOP
		 }
	 }
}

var empty struct{} //サイズがぜろの構造体

func getStatus(urls []string) <-chan string {
	statusChan := make(chan string, 3)
	limit := make(chan struct{}, 5)
	go func() {
		for _, url := range urls{
			select {
			case limit <- empty:
				go func (url string){
					res, err:= http.Get(url)
					if err != nil {
						log.Fatal(err)
					}
					statusChan <- res.Status
					<-limit
				}(url)
			}
		}
	}()
	return statusChan
}


// func getStatus(urls []string) <-chan string {
//     statusChan := make(chan string, 3)
//     // バッファを5に指定して生成
//     limit := make(chan struct{}, 5)
//     go func() {
//         for _, url := range urls {
//             select {
//             case limit <- empty:
//                 // limitに書き込みが可能な場合は取得処理を実施
//                 go func(url string) {
//                     // このゴルーチンは同時に5つしか起動しない
//                     res, err := http.Get(url)
//                     if err != nil {
//                         log.Fatal(err)
//                     }
//                     statusChan <- res.Status
//                     // 終わったら1つ読み出して空きを作る
//                     <-limit
//                 }(url)
//             }
//         }
//     }()
//     return statusChan
// }
