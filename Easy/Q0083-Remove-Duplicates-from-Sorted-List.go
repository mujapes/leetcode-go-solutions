/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {return head}
    prev := head
    cur := head
    var dup bool
    for cur.Next != nil {
        cur = cur.Next
        if cur.Val == prev.Val {
            dup = true
        } else {
            if dup {
                prev.Next = cur
                dup = false
            }
            prev = cur
        }
    }
    prev.Next = nil
    return head
}

//Runtime: 0 ms, Beats 100.00%
//Memory: 4.91 MB, Beats 25.65%
