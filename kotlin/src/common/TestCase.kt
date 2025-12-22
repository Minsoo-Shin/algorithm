package common

data class TestCase<K,V>(
    val name: String,
    val input: K,
    val expected: V
)