package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	countMap := make(map[string]int)

	for _, field := range strings.Fields(s) {
		countMap[field] += 1
	}

	return countMap
}

func main() {
	wc.Test(WordCount)
}
