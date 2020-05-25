package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
)

type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var t = template.Must(template.ParseFiles("index.txt"))

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}
func PersonHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method == "POST" {
		//リクエストボディをjsonに変換
		var person Person
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&person)
		if err != nil {
			log.Fatal(err)
		}
		//ファイル名を[id].txtとする
		filename := fmt.Sprintf("%d.txt", person.ID)
		file, err := os.Create(filename) //ファイルを作成
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		//ファイルにNameを書き込む
		_, err = file.WriteString(person.Name)
		if err != nil {
			log.Fatal(err)
		}
		//レスポンスとしてステータスコード201を送信
		w.WriteHeader(http.StatusCreated)

	} else if r.Method == "GET" {
		//パラメータを取得する
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			log.Fatal(err)
		}
		filename := fmt.Sprintf("%d.txt", id)
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}
		person := Person{
			ID:   id,
			Name: string(b),
		}
		//レスポンスにエンコードしたHTMLを書き込む
		t.Execute(w, person)
	}
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/persons", PersonHandler)
	http.ListenAndServe(":3000", nil)
}
