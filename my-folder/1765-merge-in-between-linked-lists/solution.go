/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
    list2Tail := list2
    for list2Tail.Next != nil {
        list2Tail = list2Tail.Next
    }
    
    var beforeRemoved *ListNode
    var afterRemoved *ListNode
    curr := list1
    for i:=0;;i++ {
        if i == a-1 {
            beforeRemoved = curr
        } else if i == b+1 {
            afterRemoved = curr
            break
        }
        
        curr = curr.Next
    }
    
    beforeRemoved.Next = list2
    list2Tail.Next = afterRemoved
    
    return list1
}
