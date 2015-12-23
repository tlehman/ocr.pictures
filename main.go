package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadFile("index.html")
	fmt.Fprintf(w, string(buf))
}

func recognizeHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/recognize", recognizeHandler)

	http.ListenAndServe(":8080", nil)
}
