package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"testing"
)

func TestServer(t *testing.T) {
	var wg sync.WaitGroup
	cycle := 1000 // cycle超过一定值时请求会发不出去，和本机的tcp的backlog设置有关
	wg.Add(cycle)
	for i := 0; i < cycle; i++ {
		go func() {
			http.Get("http://localhost:8000")
			wg.Done()
		}()
	}
	wg.Wait()
	resp, err := http.Get("http://localhost:8000/count")
	if err != nil {
		fmt.Println("request err ", err)
	}
	_, err = io.Copy(os.Stdout, resp.Body)
}
