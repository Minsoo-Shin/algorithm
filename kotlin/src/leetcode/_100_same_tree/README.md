# 100. Same Tree

## 문제 설명
두 개의 이진 트리 `p`와 `q`의 루트가 주어졌을 때, 두 트리가 같은지 확인하는 함수를 작성하세요.

두 이진 트리가 구조적으로 동일하고 노드의 값이 같으면 같은 트리로 간주합니다.

## 예제

### Example 1
![ex1](https://assets.leetcode.com/uploads/2020/12/20/ex1.jpg)
- **Input:** `p = [1,2,3]`, `q = [1,2,3]`
- **Output:** `true`

### Example 2
![ex2](https://assets.leetcode.com/uploads/2020/12/20/ex2.jpg)
- **Input:** `p = [1,2]`, `q = [1,null,2]`
- **Output:** `false`

### Example 3
![ex3](https://assets.leetcode.com/uploads/2020/12/20/ex3.jpg)
- **Input:** `p = [1,2,1]`, `q = [1,1,2]`
- **Output:** `false`

## 제약사항
- 두 트리의 노드 개수는 `[0, 100]` 범위입니다.
- `-10^4 <= Node.val <= 10^4`

## 구현
```kotlin
class Solution {
    class TreeNode(var `val`: Int) {
        var left: TreeNode? = null
        var right: TreeNode? = null
    }

    fun isSameTree(p: TreeNode?, q: TreeNode?): Boolean {
        val pVal = p?.`val`
        val qVal = q?.`val`
        if (pVal != qVal) {
            return false
        }
        if (pVal == null && qVal == null) {
            return true
        }
        return isSameTree(p!!.left,  q!!.left) && isSameTree(p!!.right, q!!.right)
    }
}
```

## 풀이 전략 (DFS/Recursion)
1. **기저 사례 (Base Case):**
   - 두 노드가 모두 `null`이면 구조와 값이 같으므로 `true`를 반환합니다.
   - 두 노드 중 하나만 `null`이거나, 두 노드의 값이 다르면 `false`를 반환합니다.
2. **재귀 호출 (Recursive Step):**
   - 현재 노드의 값이 같다면, 왼쪽 자식끼리 비교(`isSameTree(p.left, q.left)`)하고 오른쪽 자식끼리 비교(`isSameTree(p.right, q.right)`)합니다.
   - 두 재귀 호출의 결과가 모두 `true`여야 전체 트리가 같은 것입니다.

## 복잡도
- **시간 복잡도:** $O(N)$ (여기서 $N$은 트리의 노드 수 중 작은 쪽. 모든 노드를 한 번씩 방문해야 함)
- **공간 복잡도:** $O(H)$ (여기서 $H$는 트리의 높이. 재귀 호출 스택의 깊이만큼 공간 필요. 최악의 경우 $O(N)$)
