/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    carry := 0
    var head *ListNode
    var prev *ListNode
    c1 := l1
    c2 := l2
    for {
        if c1 == nil && c2 == nil {
            break
        }
        
        v1, v2 := 0, 0
        if c1 != nil {
            v1 = c1.Val
            c1 = c1.Next
        }
        if c2 != nil {
            v2 = c2.Val
            c2 = c2.Next
        }
        
        sum := carry + v1 + v2
        carry = sum/10
        
        cur := &ListNode{
            Val: sum%10,
        }
        
        if prev != nil {
            prev.Next = cur
        } else {
            head = cur
        }
        
        prev = cur
    }
    
    if carry > 0 {
        cur := &ListNode{
            Val: carry,
        }
        
        prev.Next = cur
    }
    
    return head
}
