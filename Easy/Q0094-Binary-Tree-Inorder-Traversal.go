/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func traverse(node *TreeNode, vals *[]int) {
    if node.Left != nil {traverse(node.Left, vals)}
    *vals = append(*vals, node.Val)
    if node.Right != nil {traverse(node.Right, vals)}
}

func inorderTraversal(root *TreeNode) []int {
    vals := make([]int, 0, 100)
    if root != nil {traverse(root, &vals)}
    return vals
}

//Runtime: 0 ms, Beats 100.00%
//Memory: 4.11 MB, Beats 30.46%
