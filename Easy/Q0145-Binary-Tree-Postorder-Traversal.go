/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    if root == nil {return []int{}}
    return append(append(postorderTraversal(root.Left), postorderTraversal(root.Right)...), root.Val)
}

// Runtime: 0 ms, Beats 100.00%
// Memory: 4.13 MB, Beats 29.87%
