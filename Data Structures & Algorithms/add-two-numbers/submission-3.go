/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	curr1 := l1
	curr2 := l2
	last := curr1
	extra := 0

	for curr1 != nil || curr2 != nil {
		if curr1 == nil {
			curr1 = &ListNode{Val: 0}
			last.Next = curr1 
		}

		if curr2 == nil {
			curr2 = &ListNode{Val: 0}
		}

		curr1.Val = curr1.Val + curr2.Val + extra
		if curr1.Val >= 10 {
			extra = 1
			curr1.Val -= 10
		} else {
			extra = 0
		}

		last = curr1
		curr1 = curr1.Next
		curr2 = curr2.Next
	}

	if extra == 1 {
		last.Next = &ListNode{Val: 1}
	}

	return l1
}
