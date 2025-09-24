/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
    if root == nil {return 0}
    if !(root.Left == nil && root.Right == nil) {
        if root.Left == nil {
            return 1 + minDepth(root.Right)
        } else if root.Right == nil {
            return 1 + minDepth(root.Left)
        }
    }
    return 1 + min(minDepth(root.Left), minDepth(root.Right))
}

// Runtime: 4 ms, Beats 46.78%
// Memory: 22.11 MB, Beats 88.50%
