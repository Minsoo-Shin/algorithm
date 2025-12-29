# 876. Middle of the Linked List

## 문제 설명
단일 연결 리스트(Singly Linked List)의 `head`가 주어졌을 때, 연결 리스트의 중간 노드를 반환하세요.

만약 중간 노드가 두 개라면, 두 번째 중간 노드를 반환하세요.

## 예제

### Example 1
- **Input:** `head = [1,2,3,4,5]`
- **Output:** `[3,4,5]`
- **설명:** 중간 노드는 값 3을 가진 노드입니다. (직렬화된 형태: [3,4,5])
  반환된 노드는 (값 3) -> (값 4) -> (값 5) 순서로 연결되어 있습니다.

### Example 2
- **Input:** `head = [1,2,3,4,5,6]`
- **Output:** `[4,5,6]`
- **설명:** 리스트의 노드가 6개이므로 중간 노드는 3과 4 두 개입니다. 이 경우 두 번째 중간 노드인 4를 반환합니다.

## 제약사항
- 리스트의 노드 개수는 `[1, 100]` 범위입니다.
- `1 <= Node.val <= 100`

## 구현
```kotlin
class Solution {
    class ListNode(var `val`: Int) {
        var next: ListNode? = null
    }

    fun middleNode(head: ListNode?): ListNode? {
        var slow = head
        var fast = head

        // fast가 null이거나 fast.next가 null일 때까지 진행
        while (fast?.next != null) {
            slow = slow?.next
            fast = fast.next?.next
        }

        // 이 방식은 루프가 끝났을 때 slow가 자동으로
        // 홀수일 땐 정중앙, 짝수일 땐 두 번째 중간을 가리킵니다.
        return slow
    }
}
```

## 풀이 전략 (Two Pointers / Floyd's Cycle Finding Algorithm)
1. **두 개의 포인터 사용:** `slow`와 `fast` 두 개의 포인터를 `head`에서 시작합니다.
2. **이동 속도 차이:** `slow`는 한 번에 한 칸씩, `fast`는 한 번에 두 칸씩 이동합니다.
3. **종료 조건:** `fast`가 리스트의 끝(`null`)에 도달하거나, `fast`의 다음 노드가 없을 때 반복을 종료합니다.
4. **결과:** `fast`가 끝에 도달했을 때, `slow`는 정확히 리스트의 중간 지점에 위치하게 됩니다.
   - 노드 개수가 홀수일 때: `fast`가 마지막 노드에 도착하고, `slow`는 정중앙에 위치.
   - 노드 개수가 짝수일 때: `fast`가 `null`에 도착하고, `slow`는 두 번째 중간 노드에 위치.

## 복잡도
- **시간 복잡도:** $O(N)$ (리스트를 한 번만 순회하면 되므로 선형 시간 소요)
- **공간 복잡도:** $O(1)$ (추가적인 공간 없이 포인터 두 개만 사용)
