/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    var prev *ListNode
	curr := head

	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp

	}

	var newPrev *ListNode
	curr = prev

	if n == 1 {
		curr = curr.Next
	}

	for curr != nil {
		temp := curr.Next
		if n == 2 { 
			temp = curr.Next.Next
		}

		curr.Next = newPrev
		newPrev = curr
		curr = temp

		n--
	}

	return newPrev
}
