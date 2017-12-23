// http stream
package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	r, w := io.Pipe()
	go func() {
		for i := 0; i < 100; i++ {
			now := time.Now().UnixNano() / 1000
			w.Write([]byte(fmt.Sprintf("%d\n", now)))
		}
		w.Close()
	}()
	_, err := http.Post("http://localhost:8080", "", r)
	if err != nil {
		panic(err)
	}
}
