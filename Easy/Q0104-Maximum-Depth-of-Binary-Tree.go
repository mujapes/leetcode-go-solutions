/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil {return 0}
    return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

// Runtime: 0 ms, Beats 100.00%
// Memory: 6.31 MB, Beats 52.19%
