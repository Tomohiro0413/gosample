package main

import (
	"log"
	"os"
)

func main() {
	//ファイルを作成
	file, err := os.Create("./file.txt")
	if err != nil {
		log.Fatal(err)
	}
	//プログラムが終わったらファイルを閉じる
	defer file.Close()

	//書き込むデータを[]byteで用意する
	message := []byte("hello world\n")

	//Write()を用いて書き込む
	_, err = file.WriteString("hello world\n")
	if err != nil {
		log.Fatal(err)
	}
}
