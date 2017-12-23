// server accept request in a stream
package main

import (
	"bufio"
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	sc := bufio.NewScanner(r.Body)

	for sc.Scan() {
		body := sc.Text()
		now := time.Now().UnixNano() / 1000
		fmt.Println(now, "\t", string(body))
	}
	if err := sc.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
