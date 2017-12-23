package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	now := time.Now().UnixNano() / 1000
	fmt.Println(now, "\t", string(body))
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
