/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// Better name: heightIfBalanced
func depthIfSimilar(node *TreeNode, depth int) int {
    if node == nil {return depth}
    depth++
    leftDepth := depthIfSimilar(node.Left, depth)
    rightDepth := depthIfSimilar(node.Right, depth)
    if leftDepth == -1 || rightDepth == -1 || leftDepth-rightDepth > 1 || rightDepth-leftDepth > 1 {return -1}
    return max(leftDepth, rightDepth)
}

func isBalanced(root *TreeNode) bool {
    if depthIfSimilar(root, 0) == -1 {return false}
    return true
}

// Runtime: 0 ms, Beats 100.00%
// Memory: 7.36 MB, Beats 56.57%
