/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// Recursive
func inOrder(node *TreeNode, flipped bool) []int {
    if node == nil {return []int{101}}
    if flipped {return append(append(inOrder(node.Right, true), node.Val), inOrder(node.Left, true)...)}
    return append(append(inOrder(node.Left, false), node.Val), inOrder(node.Right, false)...)
}

func preOrder(node *TreeNode, flipped bool) []int {
    if node == nil {return []int{101}}
    if flipped {return append(append([]int{node.Val}, preOrder(node.Right, true)...), preOrder(node.Left, true)...)}
    return append(append([]int{node.Val}, preOrder(node.Left, false)...), preOrder(node.Right, false)...)
}

func isSymmetric(root *TreeNode) bool {
    // Combination of Pre-Order + In-Order or Post-Order + In-Order traversals
    // are required to build a unique definition of a binary tree
    leftInOrder, rightInOrder, leftPreOrder, rightPreOrder := inOrder(root.Left, false), inOrder(root.Right, true), preOrder(root.Left, false), preOrder(root.Right, true)
    if len(leftInOrder) != len(rightInOrder) {return false}
    if len(leftPreOrder) != len(rightPreOrder) {return false}
    for i, n := range leftInOrder {
        if n == 101 || rightInOrder[i] == 101 {
            if n == 101 && rightInOrder[i] == 101 {
                continue
            } else {return false}
        }
        if n != rightInOrder[i] {return false}
    }
    for i, n := range leftPreOrder {
        if n == 101 || rightPreOrder[i] == 101 {
            if n == 101 && rightPreOrder[i] == 101 {
                continue
            } else {return false}
        }
        if n != rightPreOrder[i] {return false}
    }
    return true
}

// Runtime: 1 ms, Beats 3.91%
// Memory: 6.92MB, Beats 3.03%

/*
// Iterative
func isSymmetric(root *TreeNode) bool {
    leftQueue, rightQueue := make([]*TreeNode, 0), make([]*TreeNode, 0)
    var leftNode, rightNode *TreeNode
    leftQueue, rightQueue = append(leftQueue, root.Left), append(rightQueue, root.Right)
queuesPopulated:
    for len(leftQueue) > 0 && len(rightQueue) > 0 {
        for leftQueue[0] == nil || rightQueue[0] == nil {
            if leftQueue[0] == nil && rightQueue[0] == nil {
                leftQueue, rightQueue = leftQueue[1:], rightQueue[1:]
                if len(leftQueue) == 0 || len(rightQueue) == 0 {break queuesPopulated}
            } else {return false}
        }
        leftNode, rightNode = leftQueue[0], rightQueue[0]
        leftQueue, rightQueue = leftQueue[1:], rightQueue[1:]
        if leftNode.Val != rightNode.Val {return false}
        leftQueue, rightQueue = append(leftQueue, leftNode.Left, leftNode.Right), append(rightQueue, rightNode.Right, rightNode.Left)
    }
    if len(leftQueue) != len(rightQueue) {return false}
    return true
}

// Runtime: 0 ms, Beats: 100.00%
// Memory: 4.76 MB, Beats 78.61%
*/
