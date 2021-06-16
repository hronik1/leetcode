/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }
    
    if l2 == nil {
        return l1
    }
    
    var outHead *ListNode
    var prev *ListNode 
    for {
        if l1 == nil {
            prev.Next = l2
            break
        }
    
        if l2 == nil {
            prev.Next = l1
            break
        }
        
        var cur *ListNode
        if l2.Val < l1.Val {
            cur = l2
            l2 = l2.Next
        } else {
            cur = l1
            l1 = l1.Next
        }
        
        if outHead == nil {
            outHead = cur
        } else {
            prev.Next = cur
        }
        
        prev = cur
    }
    
    return outHead
}
