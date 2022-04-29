package main

import "fmt"

func main() {
	fmt.Println(isAnagram("abb", "bba"))
	fmt.Println(isAnagram("aab", "bba"))
	fmt.Println(isAnagram("aa", "bba"))
}

func isAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	cachea := make(map[rune]int, len(a))
	for _, v := range a {
		cachea[v]++
	}

	for _, v := range b {
		if _, ok := cachea[v]; !ok {
			return false
		}

		cachea[v]--
		if cachea[v] < 0 {
			return false
		}
	}
	return true
}
