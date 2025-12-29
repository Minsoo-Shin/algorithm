package leetcode._100_same_tree

class Solution{
    class TreeNode(var `val`: Int) {
        var left: TreeNode? = null
        var right: TreeNode? = null
    }

    fun isSameTree(p: TreeNode?, q: TreeNode?): Boolean {
        if (p == null && q == null) return true
        if (p == null || q == null || q.`val` != p.`val`) return false
        return isSameTree(p.left,  q.left) && isSameTree(p.right, q.right)
    }
}