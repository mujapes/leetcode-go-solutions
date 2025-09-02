/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	node1 := l1
	node2 := l2
	node3 := l3

	sum := node1.Val + node2.Val
	node3.Val = sum % 10
	if sum > 9 {
		sum = 1
	} else {
        sum = 0
    }

	for node1.Next != nil || node2.Next != nil {
        if node1.Next == nil {
    		node2 = node2.Next
    		sum += node2.Val
        } else if node2.Next == nil {
            node1 = node1.Next
    		sum += node1.Val
        } else {
            node1 = node1.Next
    		node2 = node2.Next
    		sum += node1.Val + node2.Val
        }
        node3.Next = &ListNode{Val: sum % 10}
    	node3 = node3.Next
    	if sum > 9 {
    		sum = 1
    	} else {
            sum = 0
        }
	}
	if sum == 1 {
		node3.Next = &ListNode{Val: 1}
	}
	return l3
}

//Runtime: 0 ms, Beats 100%
//Memory: 6.26 MB, Beats 59.62%
