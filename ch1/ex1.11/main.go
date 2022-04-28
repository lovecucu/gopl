package main

// go run .\ch1\ex1.11\main.go
// http://www.baidu.com http://www.qq.com http://www.jd.com http://www.taobao.com http://www.sina.com http://www.google.com
// www.google.com一直无响应，导致程序卡住，直至http.Get超时

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for i, url := range os.Args[1:] {
		go fetch(url, ch, i)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string, index int) {
	start := time.Now()
	resp, err := http.Get(url)
	var f *os.File
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	filename := strings.Replace(url, "http://", "", -1) + fmt.Sprint(index) + ".html"
	f, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666) //打开文件
	if err != nil {
		ch <- fmt.Sprintf("while creating %s: %v", filename, err)
	}
	defer f.Close()
	nbytes, err := io.Copy(f, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
