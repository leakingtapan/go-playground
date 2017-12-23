package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

func main() {
	client := http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", "/tmp/server.sock")
			},
		},
	}

	r, w := io.Pipe()
	go func() {
		for i := 0; i < 100; i++ {
			now := time.Now().UnixNano() / 1000
			w.Write([]byte(fmt.Sprintf("%d\n", now)))
		}
		w.Close()
	}()

	_, err := client.Post("http://unix/tmp/server.sock", "", r)
	if err != nil {
		panic(err)
	}
}
