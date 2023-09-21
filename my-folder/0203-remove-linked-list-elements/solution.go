/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    if head == nil {
        return head
    }
    
    var newHead *ListNode 
    cur := head
    var prev *ListNode
    for cur != nil {
        if cur.Val == val {
            if prev != nil {
                prev.Next = cur.Next
            }
        } else {
            prev = cur
            if newHead == nil {
                newHead = cur
            }
        }
        
        cur = cur.Next
    }
    
    return newHead
}
