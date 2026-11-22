/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head.Next
	for fast != nil {
		slow = slow.Next

		if fast.Next == nil {
			break
		}
		fast = fast.Next.Next
	}

	return slow
}
