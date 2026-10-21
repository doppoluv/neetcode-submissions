/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    slow, fast := head, head
    for fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    
    right := slow.Next
    slow.Next = nil
    left := head

    left = sortList(left)
    right = sortList(right)

    if right.Val < left.Val {
        head = right
        right = right.Next
    } else {
        head = left
        left = left.Next
    }

    curr := head
    for left != nil && right != nil {
        if right.Val < left.Val {
            curr.Next = right
            right = right.Next
        } else {
            curr.Next = left
            left = left.Next
        }

        curr = curr.Next
    }

    if left != nil {
        curr.Next = left
    } else {
        curr.Next = right
    }

    return head
}