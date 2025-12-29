package leetcode._100_same_tree

class Solution{
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