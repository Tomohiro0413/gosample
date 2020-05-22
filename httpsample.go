package main

import (
    "fmt"
    "net/http"

)

func IndexHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"hello world")
}

func main() {
	http.HandlerFunc("/", IndexHandler)
	http.ListenAndServe(":3000",nil)
}