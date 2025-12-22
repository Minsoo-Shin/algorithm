package leetcode._xx_unique_substring_set

import common.TestCase
import kotlin.reflect.KClass

fun main() {
    testSolution(Solution::class, Solution()::longestUniqueSubstringSet)
}

fun testSolution(testClass: KClass<*>, fn: (String) -> List<String>) {
    println("=== Testing ${testClass.simpleName} ===")

    val testCases = listOf(
        TestCase(
            name = "Example 1",
            input = "abcbedfed",
            expected = listOf("a", "bcb", "edfed")
        ),
        TestCase(
            name = "Example 2",
            input = "abcdefg",
            expected = listOf("a", "b", "c", "d", "e", "f", "g")
        ),
        TestCase(
            name = "Example 3",
            input = "abmowodfsxadejihgepczpc",
            expected = listOf("abmowodfsxad", "ejihge", "pczpc")
        )
    )

    testCases.forEach { case ->
        val result = fn(case.input)
        if (result == case.expected) {
            println("[✅ Pass] ${case.name}")
        } else {
            println("[❌ Fail] ${case.name}")
            println("   Expected: ${case.expected}")
            println("   Actual:   $result")
        }
    }
}