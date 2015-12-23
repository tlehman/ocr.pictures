package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const googleVisionApiUrl = "https://vision.googleapis.com/v1alpha1/images:annotate?key="

func indexHandler(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadFile("index.html")
	fmt.Fprintf(w, string(buf))
}

func recognizeHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	url := googleVisionApiUrl + os.Getenv("GVAPIKEY")
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("request error: %v", err)
	}
	resbody, _ := ioutil.ReadAll(res.Body)
	fmt.Fprintf(w, "%s", resbody)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/recognize", recognizeHandler)

	http.ListenAndServe(":8080", nil)
}
