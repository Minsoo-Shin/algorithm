package valid_anagram

// O(n) : s를 다 map에 등록하여 t에서 빼는 방법으로 진행하고 마지막이 모두 값이 없다고 하면 일치하는 것
func isAnagram(s string, t string) bool {
	sRune, tRune := []rune(s), []rune(t)
	checkMap := make(map[rune]int)

	for _, v := range sRune {
		checkMap[v]++
	}

	for _, v := range tRune {
		qty, exists := checkMap[v]
		if !exists || qty == 0 {
			return false
		}
		checkMap[v]--
	}

	for _, v := range checkMap {
		if v != 0 {
			return false
		}
	}
	return true
}

func isAnagram2(s string, t string) bool {
	numSet := make([]int, 26)

	for _, ss := range s {
		numSet[int(ss-'a')]++
	}

	for _, tt := range t {
		numSet[int(tt-'a')]--
	}

	for _, val := range numSet {
		if val != 0 {
			return false
		}
	}
	return true
}
