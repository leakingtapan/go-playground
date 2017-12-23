package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"
)

type handler struct {
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
