/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    node := head
    cnt := 0
    for node != nil {
        cnt++
        node = node.Next
    }
    cnt -= n
    if cnt == 0 {return head.Next}
    node = head
    for _ = range cnt-1 {
        node = node.Next
    }
    node.Next = node.Next.Next
    return head
}

// Runtime: 0 ms, Beays 100.00%
// Memory: 4.16 MB, Beats 45.82%
