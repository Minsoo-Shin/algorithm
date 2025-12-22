// Main.kt
package leetcode._56_merge_intervals

import kotlin.reflect.KClass

fun main() {
    testSolution(SolutionGreedy::class, SolutionGreedy()::merge)
}

fun testSolution(testClass: KClass<*>, mergeFn: (Array<IntArray>) -> Array<IntArray>) {
    println("=== Testing ${testClass.simpleName} ===")

    val testCases = listOf(
        TestCase(
            name = "Example 1",
            input = arrayOf(intArrayOf(1, 3), intArrayOf(2, 6), intArrayOf(8, 10), intArrayOf(15, 18)),
            expected = arrayOf(intArrayOf(1, 6), intArrayOf(8, 10), intArrayOf(15, 18))
        ),
        TestCase(
            name = "Example 2",
            input = arrayOf(intArrayOf(1, 4), intArrayOf(4, 5)),
            expected = arrayOf(intArrayOf(1, 5))
        ),
        TestCase(
            name = "Example 3",
            input = arrayOf(intArrayOf(4, 7), intArrayOf(1, 4)),
            expected = arrayOf(intArrayOf(1, 7))
        )
    )

    testCases.forEach { case ->
        val result = mergeFn(case.input)
        if (result.contentDeepEquals(case.expected)) {
            println("[✅ Pass] ${case.name}")
        } else {
            println("[❌ Fail] ${case.name}")
            println("   Expected: ${case.expected.contentDeepToString()}")
            println("   Actual:   ${result.contentDeepToString()}")
        }
    }
}