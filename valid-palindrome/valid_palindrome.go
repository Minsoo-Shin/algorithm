package valid_palindrome

import (
	"regexp"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	// [^\p{L}\p{N} ]+
	re, _ := regexp.Compile("[^a-zA-Z\\d]+")
	key := strings.ToLower(re.ReplaceAllString(s, ""))

	rune := []rune(key)
	l := len(rune)

	if l == 0 {
		return true
	}
	for i := 0; i < int(l/2)+1; i++ {
		if rune[i] != rune[l-1-i] {
			return false
		}
	}

	return true
}

func isPalindrome2(s string) bool {
	f := func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return -1
		}
		return unicode.ToLower(r)
	}

	strings := strings.Map(f, s)

	i, j := 0, len([]rune(strings))-1
	for i < j {
		if strings[i] != strings[j] {
			return false
		}

		i += 1
		j -= 1
	}
	return true
}
