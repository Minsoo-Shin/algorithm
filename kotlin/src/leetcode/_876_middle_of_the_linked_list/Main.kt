package leetcode._876_middle_of_the_linked_list

import leetcode.common.ListNode
import common.TestCase
import kotlin.reflect.KClass

fun main() {
    testSolution(Solution::class, Solution()::middleNode)
}

fun testSolution(testClass: KClass<*>, fn: (head: ListNode?) -> ListNode? ) {
    println("=== Testing ${testClass.simpleName} ===")

    val testCases = listOf(
        TestCase(
            name = "Example 1",
            input = buildLinkedList(listOf(1,2,3,4,5)),
            expected = buildLinkedList(listOf(3,6,5))
        ),
        TestCase(
            name = "Example 2",
            input = buildLinkedList(listOf(1,2,3,4,5,6)),
            expected = buildLinkedList(listOf(4,5,6))
        ),
    )

    testCases.forEach { case ->
        val p = case.input
        val result = fn(p)
        if (result == case.expected) {
            println("[✅ Pass] ${case.name}")
        } else {
            println("[❌ Fail] ${case.name}")
            println("   Expected: ${case.expected}")
            println("   Actual:   $result")
        }
    }
}

fun buildLinkedList(nodes: List<Int>): ListNode? {
    if (nodes.isEmpty()) return null

    val head = ListNode(nodes[0])
    var curr = head

    for (i in 1 until nodes.size) {
        curr.next = ListNode(nodes[i])
        curr = curr.next!!
    }
    return head
}
