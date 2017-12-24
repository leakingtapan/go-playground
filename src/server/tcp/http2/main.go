package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/net/http2"
)

type handler struct {
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	now := time.Now().UnixNano() / 1000
	fmt.Println(now, "\t", string(body))
	w.WriteHeader(http.StatusOK)
}

func main() {
	h := &handler{}
	server := http.Server{
		Addr:    ":8080",
		Handler: h,
	}
	http2.ConfigureServer(&server, &http2.Server{})
	server.ListenAndServeTLS("./domain.crt", "./domain.key")
}
