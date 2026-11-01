/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
    curr := head
    for curr.Next != nil {
        temp := curr.Next
        div := greatestDiv(curr.Val, curr.Next.Val)
        curr.Next = &ListNode{Val: div, Next: temp}
        curr = temp
    }

    return head
}

func greatestDiv(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }

    return a
}