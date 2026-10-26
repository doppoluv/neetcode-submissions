/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
    if root == nil {
		return 0
	}

	depth := 1
	if maxDepth(root.Left) >= maxDepth(root.Right) {
		depth += maxDepth(root.Left)
	} else {
		depth += maxDepth(root.Right)
	}

	return depth
}
