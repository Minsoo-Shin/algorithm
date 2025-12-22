package leetcode._xx_unique_substring_set

class Solution{
    fun longestUniqueSubstringSet(s: String): List<String> {
        val result = mutableListOf<String>()
        if (s.isEmpty()) return result

        // 1. 각 문자의 마지막 위치를 기록 (알파벳 소문자 26개 고정 크기 배열)
        val lastOccurrences = IntArray(26)
        for (i in s.indices) {
            lastOccurrences[s[i] - 'a'] = i
        }

        var start = 0
        var end = 0

        // 2. 문자열을 순회하며 구간 확정
        for (i in s.indices) {
            // 현재 문자의 마지막 위치와 기존 end 중 더 먼 곳을 선택
            end = maxOf(end, lastOccurrences[s[i] - 'a'])

            // 현재 인덱스가 구간의 끝(end)에 도달했다면 분리
            if (i == end) {
                result.add(s.substring(start, end + 1))
                start = i + 1
            }
        }

        return result
    }
}