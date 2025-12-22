# Longest Unique Substring Set - Apple

## 문제 설명
문자열 `s`가 주어집니다. 문자열 `s`를 최대한 많은 개수의 substring으로 나누는데, 각 substring에 들어있는 문자들은 다른 substring에는 존재하지 않아야 합니다. 즉 각각의 substring들은 각각 유일한 문자들로만 이루어진 문자열이어야 합니다.

최대한 많은 substring으로 나눈 문자열들을 리턴하세요.

## 예제
### Example 1
- **Input:** `s = "abcbedfed"`
- **Output:** `["a", "bcb", "edfed"]`
- **설명:** 3개의 문자열 `a`, `bcb`, `edfed`로 나누어질 수 있고 각 문자열은 유일한 문자들로만 이루어져있습니다.
    - `a` = [a]
    - `bcb` = [b, c]
    - `edfed` = [e, d, f]

### Example 2
- **Input:** `s = "abcdefg"`
- **Output:** `["a", "b", "c", "d", "e", "f", "g"]`

### Example 3
- **Input:** `s = "abmowodfsxadejihgepczpc"`
- **Output:** `["abmowodfsxad", "ejihge", "pczpc"]`
- **설명:**
    - `abmowodfsxad` = [a, b, m, o, w, d, f, s, x]
    - `ejihge` = [e, j, i, h, g]
    - `pczpc` = [p, c, z]

## 제약사항
- 1 <= `s.length` <= 2^30
- 1 <= `result.size()` <= 26
- `s`는 알파벳 소문자로만 이루어져있습니다.

## 구현
```kotlin
fun longestUniqueSubstringSet(s: String): List<String> {
    // implementation
}
```

## 풀이 전략 (Greedy)
1. **마지막 위치 기록:** 먼저 문자열을 한 번 순회하며 각 알파벳이 마지막으로 등장하는 인덱스를 저장합니다.
2. **구간 확장:** 문자열을 다시 순회하면서 현재 구간에 포함된 문자들의 '마지막 위치' 중 가장 먼 곳을 `end`로 갱신합니다.
3. **구간 확정:** 현재 인덱스 `i`가 `end`와 같아지면, 현재까지의 구간을 하나의 부분 문자열로 확정하고 다음 구간을 시작합니다.

## 복잡도
- **시간 복잡도:** $O(N)$ (문자열을 두 번 순회)
- **공간 복잡도:** $O(1)$ (알파벳 26개를 저장하기 위한 고정 크기 배열 사용)
