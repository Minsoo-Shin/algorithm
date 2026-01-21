package leetcode._208_implement_trie

import kotlin.text.iterator

class Trie() {

    class TrieNode(
        val children: MutableMap<Char, TrieNode> = mutableMapOf(),
        var isEndOfWord: Boolean = false,
    )

    val root = TrieNode()

    fun insert(word: String) {
        var currentNode = root

        for (char in word) {
            currentNode = currentNode.children.getOrPut(char) { TrieNode() }
        }

        currentNode.isEndOfWord = true
    }

    fun search(word: String): Boolean {

        var currentNode = root

        for (char in word) {
            currentNode.children[char]?.let{
                currentNode = it
            } ?: return false
        }

        return currentNode.isEndOfWord
    }

    fun startsWith(prefix: String): Boolean {

        var currentNode = root

        for (char in prefix) {
            currentNode.children[char]?.let{
                currentNode = it
            } ?: return false
        }

        return true
    }
}