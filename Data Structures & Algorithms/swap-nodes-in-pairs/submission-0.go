/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    curr := head
    var prev *ListNode
    start := curr.Next
    for curr != nil && curr.Next != nil {
        if prev != nil {
            prev.Next = curr.Next
        }

        temp := curr.Next.Next
        curr.Next.Next = curr
        curr.Next = temp

        prev = curr
        curr = temp
    }

    return start
}