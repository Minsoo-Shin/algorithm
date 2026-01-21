package leetcode._208_implement_trie

import leetcode.common.ListNode
import common.TestCase
import kotlin.reflect.KClass

fun main() {
    val newTrie = Trie()
    newTrie.insert("apple")
    newTrie.insert("and")

    println(newTrie.search("apple"))
    println(newTrie.search("and"))
    println(newTrie.search("app"))
    println(newTrie.startsWith("app"))
}