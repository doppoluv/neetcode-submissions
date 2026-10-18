/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    if head == nil {
		return false
	}

	curr := head
	for curr.Next != nil {
		if curr.Next.Next == head {
			return true
		}

		temp := curr.Next
		curr.Next = head
		curr = temp
	}

	return false
}
