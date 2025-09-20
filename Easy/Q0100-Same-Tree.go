/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {return true}
    if p == nil || q == nil {return false}
    if p.Val != q.Val {return false}
    return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

//Runtime: 0 ms, Beats 100.00%
//Memory: 4.28 MB, Beats 11.33%
