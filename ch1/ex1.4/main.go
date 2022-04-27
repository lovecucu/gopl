package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

func main() {
	type item struct {
		num   int
		files map[string]struct{}
	}
	counts := make(map[string]*item)

	for _, file := range os.Args[1:] {
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		input := bufio.NewScanner(f)
		for input.Scan() {
			if row, ok := counts[input.Text()]; ok {
				row.num++
				if _, ok := row.files[file]; !ok {
					row.files[file] = struct{}{}
				}
			} else {
				row := item{
					num:   1,
					files: make(map[string]struct{}),
				}
				row.files[file] = struct{}{}
				counts[input.Text()] = &row
			}
		}
		f.Close()
	}

	for line, row := range counts {
		if row.num > 1 {
			fmt.Printf("%d\t%s\t%s\n", row.num, line, reflect.ValueOf(row.files).MapKeys())
		}
	}
}
