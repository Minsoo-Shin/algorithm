package leetcode._100_same_tree

import common.TestCase
import leetcode._100_same_tree.Solution.TreeNode
import kotlin.reflect.KClass
import java.util.ArrayDeque

fun main() {
    testSolution(Solution::class, Solution()::isSameTree)
}

fun testSolution(testClass: KClass<*>, fn: (p: TreeNode?, q: TreeNode?) -> Boolean ) {
    println("=== Testing ${testClass.simpleName} ===")

    val testCases = listOf(
        TestCase(
            name = "Example 1",
            input = Pair(
                buildTree(listOf(1, 2, 3)),
                buildTree(listOf(1, 2, 3))
            ),
            expected = true
        ),
        TestCase(
            name = "Example 2",
            input = Pair(
                buildTree(listOf(1, 2)),
                buildTree(listOf(1, null, 2))
            ),
            expected = false
        ),
        TestCase(
            name = "Example 3",
            input = Pair(
                buildTree(listOf(1, 2, 1)),
                buildTree(listOf(1, 1, 2))
            ),
            expected = false
        )
    )

    testCases.forEach { case ->
        val (p, q) = case.input
        val result = fn(p, q)
        if (result == case.expected) {
            println("[✅ Pass] ${case.name}")
        } else {
            println("[❌ Fail] ${case.name}")
            println("   Expected: ${case.expected}")
            println("   Actual:   $result")
        }
    }
}

fun buildTree(nodes: List<Int?>): TreeNode? {
    if (nodes.isEmpty() || nodes[0] == null) return null

    val root = TreeNode(nodes[0]!!)
    val queue = ArrayDeque<TreeNode>()

    queue.add(root)

    var i = 1
    while (i < nodes.size) {
        val curr = queue.removeFirst()

        if (i < nodes.size) {
            val leftVal = nodes[i++]

            if (leftVal != null) {
                curr.left = TreeNode(leftVal)
                queue.add(curr.left!!)
            }
        }

        if (i < nodes.size) {
            val rightVal = nodes[i++]

            if (rightVal != null) {
                curr.right = TreeNode(rightVal)
                queue.add(curr.right!!)
            }
        }
    }
    return root
}
