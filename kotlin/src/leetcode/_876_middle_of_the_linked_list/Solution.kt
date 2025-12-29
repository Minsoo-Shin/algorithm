package leetcode._876_middle_of_the_linked_list

import leetcode.common.ListNode

class Solution {
    fun middleNode(head: ListNode?): ListNode? {
        var slow = head
        var fast = head

        while (fast?.next != null) {
            slow = slow?.next
            fast = fast.next?.next
        }

        return slow
    }
}