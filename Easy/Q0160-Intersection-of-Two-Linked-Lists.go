/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    visited := make(map[*ListNode]struct{})
    for headA != nil || headB != nil {
        if headA != nil {
            if _, ok := visited[headA]; ok {return headA}
            visited[headA] = struct{}{}
            headA = headA.Next
        }
        if headB != nil {
            if _, ok := visited[headB]; ok {return headB}
            visited[headB] = struct{}{}
            headB = headB.Next
        }
    }
    return nil
}

// Runtime: 29 ms, Beats 50.97%
// Memory: 9.25 MB, Beats 13.27%
