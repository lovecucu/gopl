package main

import (
	"bufio"
	"fmt"
	"gopl/ch2/tempconv"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			print(arg)
		}
		return
	}
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		print(input.Text())
	}
}

func print(s string) {
	t, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
}
