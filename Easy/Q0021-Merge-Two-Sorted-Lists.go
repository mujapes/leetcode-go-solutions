/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    var mergedList *ListNode
    firstNode := mergedList
    for {
        if list1 != nil {
            if mergedList != nil {
                mergedList.Next = &ListNode{}
                mergedList = mergedList.Next
            } else {
                mergedList = &ListNode{}
                firstNode = mergedList
            }
            if list2 != nil {
                if list1.Val<=list2.Val {
                    mergedList.Val = list1.Val
                    list1 = list1.Next
                } else {
                    mergedList.Val = list2.Val
                    list2 = list2.Next
                }
            } else {
                mergedList.Val = list1.Val
                list1 = list1.Next
            }
        } else if list2 != nil {
            if mergedList != nil {
                mergedList.Next = &ListNode{}
                mergedList = mergedList.Next
            } else {
                mergedList = &ListNode{}
                firstNode = mergedList
            }
            mergedList.Val = list2.Val
            list2 = list2.Next
        } else {
            return firstNode
        }
    }
}

//Runtime: 0 ms, Beats 100.00%
//Memory: 4.47 MB, Beats 39.84%
