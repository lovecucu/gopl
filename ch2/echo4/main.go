package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
	fmt.Println("aaa") // 用于标识是否有换行

	// go run .\ch2\echo4\main.go -n -s / a bc def
	// 输出：a/bc/defaaaa
}
