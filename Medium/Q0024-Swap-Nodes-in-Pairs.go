/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {return head}
    prevNode := head
    node, head := head.Next, head.Next
    var tempNext *ListNode
    for prevNode != nil && node != nil {
        if node.Next != nil && node.Next.Next != nil {
            tempNext = node.Next
            prevNode.Next, node.Next = node.Next.Next, prevNode
            prevNode, node = tempNext, prevNode.Next
        } else {
            prevNode.Next, node.Next = node.Next, prevNode
            node = nil
        }
    }
    return head
}

// Runtime: 0 ms, Beats 100.00%
// Memory: 4.17 MB, Beats 2.74%
