package main

import (
	"bytes"
	"context"
	"fmt"
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

	for i := 0; i < 100; i++ {
		now := time.Now().UnixNano() / 1000
		buf := bytes.NewReader([]byte(fmt.Sprintf("%d", now)))
		_, err := client.Post("http://unix/tmp/server.sock", "", buf)
		if err != nil {
			panic(err)
		}
	}
}
