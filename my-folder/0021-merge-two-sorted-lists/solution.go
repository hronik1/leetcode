/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil {
        return list2
    } else if list2 == nil {
        return list1
    }
    
    head := list1
    tail := list1
    if list2.Val < list1.Val {
        head = list2
        tail = list2
        list2 = list2.Next
    } else {
        list1 = list1.Next
    }
    
    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            tail.Next = list1
            tail = list1
            list1 = list1.Next
        } else {
            tail.Next = list2
            tail = list2
            list2 = list2.Next
        }
    }
    
    if list1 != nil {
        tail.Next = list1
    } else if list2 != nil {
        tail.Next = list2
    }
    
    return head
}
