package main

import "strings"

func firstUniqChar(s string) int {
	checkMap := map[rune]struct{}{}

	for i, r := range s {
		if _, ok := checkMap[r]; ok {
			continue
		}
		nextIndex := strings.IndexRune(s[i+1:], r)

		if nextIndex == -1 {
			return i
		}

		checkMap[r] = struct{}{}
	}
	return -1
}
