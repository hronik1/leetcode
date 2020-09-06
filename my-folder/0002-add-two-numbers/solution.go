/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    cur1 := l1
    cur2 := l2
    cur := &ListNode{}
    head := cur
    prev := &ListNode{}
    remainder := 0
    
    for cur1 != nil || cur2 != nil {
        val := remainder
        if cur1 != nil {
            val += cur1.Val
            cur1 = cur1.Next
        }
        if cur2 != nil {
            val += cur2.Val
            cur2 = cur2.Next
        }
        
        remainder = val/10
        cur.Val = val%10
        cur.Next = &ListNode{}
        prev = cur
        cur = cur.Next
    }
    
    if remainder > 0 {
        cur.Val = remainder
    } else {
        prev.Next = nil
    }
    
    return head
}
