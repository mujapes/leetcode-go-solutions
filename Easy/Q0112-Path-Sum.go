/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {return false}
    if root.Left == nil && root.Right == nil {return targetSum == root.Val}
    return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

// Runtime: 0 ms, Beats 100.00%
// Memory: 6.65 MB, Beats 85.29%
