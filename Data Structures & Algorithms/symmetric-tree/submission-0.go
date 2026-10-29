/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }

    return isEqual(root.Left, root.Right)
}

func isEqual(p *TreeNode, q *TreeNode) bool {
    if p == nil || q == nil {
        if p == nil && q == nil {
            return true
        }

        return false
    }

    if p.Val != q.Val {
        return false
    }

    if !isEqual(p.Left, q.Right) || !isEqual(q.Left, p.Right) {
        return false
    }

    return true
}