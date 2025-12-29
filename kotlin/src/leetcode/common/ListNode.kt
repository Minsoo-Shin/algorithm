package leetcode.common

class ListNode(var `val`: Int) {
    var next: ListNode? = null
    // 테스트 환경에서만 사용 (재귀)
    override fun equals(other: Any?): Boolean {
        if (this === other) return true
        if (javaClass != other?.javaClass) return false
        other as ListNode
        return `val` == other.`val` && next == other.next
    }
    // 테스트 환경에서만 사용 (재귀)
    override fun toString(): String {
        return "$`val`->$next"
    }
}