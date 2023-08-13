package main

/*
skip문자열로 후보군을 추리고, 인덱스에 맞게 뽑아준다.
후보군 배열 길이의 나머지로 재정의를 하면서 알파벳을 가져온다.
*/

import (
	"strings"
)

func toCharRune(i int) rune {
	return rune('a' - 1 + i)
}

func excludeSkip(candidate []rune, skip string) []rune {
	for i := 1; i < 27; i++ {
		if strings.IndexRune(skip, toCharRune(i)) != -1 {
			continue
		}
		candidate = append(candidate, toCharRune(i))
	}

	return candidate
}

func solution(s string, skip string, index int) string {
	var candidate []rune

	c := excludeSkip(candidate, skip)

	ret := ""

	for _, v := range s {
		for i, vv := range c {
			if vv == v {
				ret += string(c[(i+index)%len(c)])
				continue
			}
		}
	}

	return ret
}
