/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {
    traversal := []int{}
    if root == nil {
        return traversal
    }

    getTraversal(root, &traversal)

    return traversal
}

func getTraversal(node *TreeNode, t *[]int) {
    if node.Left != nil {
        getTraversal(node.Left, t)
    }
    *t = append(*t, node.Val)
    if node.Right != nil {
        getTraversal(node.Right, t)
    }
}