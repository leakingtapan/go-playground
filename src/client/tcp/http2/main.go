// http2 client with tcp socket
package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/http2"
)

func main() {
	// Disable security check for self signed cert
	// use http2 pkg to set to client
	tr := &http2.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	for i := 0; i < 100; i++ {
		now := time.Now().UnixNano() / 1000
		buf := bytes.NewReader([]byte(fmt.Sprintf("%d", now)))
		_, err := client.Post("https://localhost:8080", "", buf)
		if err != nil {
			panic(err)
		}
	}
}
