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

func greatestDiv(v1 int, v2 int) int {
    res := 1

    if v1 < v2 {
        for i := 1; i <= int(math.Sqrt(float64(v1))); i++ {
            if v1 % i == 0 {
                if i > res && v2 % i == 0 {
                    res = i
                }
                if (v1/i) > res && v2 % (v1/i) == 0 {
                    res = (v1/i)
                }
            }
        }
    } else {
        for i := 1; i <= int(math.Sqrt(float64(v2))); i++ {
            if v2 % i == 0 {
                if i > res && v1 % i == 0 {
                    res = i
                }
                if (v2/i) > res && v1 % (v2/i) == 0 {
                    res = (v2/i)
                }
            }
        }
    }

    return res
}