package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		now := time.Now().UnixNano() / 1000
		buf := bytes.NewReader([]byte(fmt.Sprintf("%d", now)))
		_, err := http.Post("http://localhost:8080", "", buf)
		if err != nil {
			panic(err)
		}
	}
}
