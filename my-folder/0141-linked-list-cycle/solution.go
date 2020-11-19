/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return false
    }
    
    slow, fast := head.Next, head.Next.Next
    for slow != nil && fast != nil{
        if slow == fast {
            return true
        }
        
        slow = slow.Next
        if fast.Next != nil {
            fast = fast.Next.Next
        } else {
            fast = nil
        }
    }
    
    return false
}
