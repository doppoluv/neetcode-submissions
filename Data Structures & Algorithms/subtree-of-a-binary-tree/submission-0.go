/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot != nil {
		return false
	}
	
    if isEqual(root, subRoot) || subRoot == nil {
		return true
	}

	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
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

	if !isEqual(p.Left, q.Left) || !isEqual(p.Right, q.Right) {
		return false
	}

	return true
}