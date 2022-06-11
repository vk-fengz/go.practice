package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func saybodyHandler(w http.ResponseWriter, r *http.Request) {
	s, _ := ioutil.ReadAll(r.Body) // 把body 内容读入字符串 s
	getJson, _ := json.Marshal(s)
	fmt.Println(string(getJson))
}

func main() {
	http.HandleFunc("/", saybodyHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
