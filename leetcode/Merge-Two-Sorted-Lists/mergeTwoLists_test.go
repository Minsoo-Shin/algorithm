package mergeTwoLists

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListNode_AddNode(t *testing.T) {
	tests := []struct {
		list1    *ListNode
		list2    *ListNode
		expected *ListNode
	}{
		{
			list1: New().
				AddNode(1).
				AddNode(2).
				AddNode(3).
				AddNode(4).
				Head,
			list2: New().
				AddNode(2).
				AddNode(3).
				AddNode(5).
				Head,
			expected: New().
				AddNode(1).
				AddNode(2).
				AddNode(2).
				AddNode(3).
				AddNode(3).
				AddNode(4).
				AddNode(5).
				Head,
		},
		{
			list1: New().
				Head,
			list2: New().
				AddNode(3).
				AddNode(5).
				Head,
			expected: New().
				AddNode(3).
				AddNode(5).
				Head,
		},
		{
			list1: New().
				AddNode(4).
				Head,
			list2: New().
				AddNode(3).
				AddNode(5).
				Head,
			expected: New().
				AddNode(3).
				AddNode(4).
				AddNode(5).
				Head,
		},
	}
	for _, v := range tests {
		assert.Equal(t, v.expected, mergeTwoLists(v.list1, v.list2))
		assert.Equal(t, v.expected, mergeTwoLists2(v.list1, v.list2))
	}
}

func TestCreateNode(t *testing.T) {
	list := New().AddNode(1).AddNode(2).AddNode(3)

	for {
		if list.Head != nil {
			fmt.Println(list.Head.Val)
		}
		if list.Head.Next != nil {
			list.Head = list.Head.Next
		} else {
			break
		}
	}
}
