package mergeTwoLists

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// Node 2개를 동시에 확인하고, 더 작은 노드를 mergedNode 추가한다.
	mergeNodeWithHead := New()
	if list1 == nil && list2 == nil {
		return nil
	}

	for list1 != nil || list2 != nil {
		if list1 == nil {
			mergeNodeWithHead.AddNode(list2.Val)
			list2 = list2.Next
			continue
		}
		if list2 == nil {
			mergeNodeWithHead.AddNode(list1.Val)
			list1 = list1.Next
			continue
		}
		if list1.Val < list2.Val {
			mergeNodeWithHead.AddNode(list1.Val)
			list1 = list1.Next
			continue
		}
		mergeNodeWithHead.AddNode(list2.Val)
		list2 = list2.Next
		continue
	}

	return mergeNodeWithHead.Head
}

type List struct {
	Head *ListNode
	Tail *ListNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func New() *List {
	return &List{Head: nil, Tail: nil}
}

// 제일 처음 노드( 리스트의 head)
func (l *List) First() *ListNode {
	return l.Head
}

func (l *List) AddNode(val int) *List {
	node := &ListNode{Val: val, Next: nil}
	if l.Head == nil {
		l.Head = node
	} else {
		l.Tail.Next = node
	}
	l.Tail = node
	return l
}
