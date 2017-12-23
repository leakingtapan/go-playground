package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"time"
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
	defer func() {
		os.Remove("/tmp/server.sock")
	}()
	h := &handler{}
	server := http.Server{
		Handler: h,
	}

	listerner, err := net.Listen("unix", "/tmp/server.sock")
	if err != nil {
		panic(err)
	}
	server.Serve(listerner)
}
