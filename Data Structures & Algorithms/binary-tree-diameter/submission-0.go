/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil {
		return 0
	}

	currDiam := maxDepth(root.Left) + maxDepth(root.Right) 
	maxLeftDiam := diameterOfBinaryTree(root.Left)
	maxRightDiam := diameterOfBinaryTree(root.Right)

	return int(math.Max(math.Max(float64(maxLeftDiam), float64(maxRightDiam)), float64(currDiam)))
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + int(math.Max(float64(maxDepth(root.Left)), float64(maxDepth(root.Right))))
}