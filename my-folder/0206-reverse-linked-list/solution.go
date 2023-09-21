/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    
    originalHead := head
    newHead := head
    for originalHead.Next != nil {
        prevHead := newHead
        newHead = originalHead.Next
        originalHead.Next = newHead.Next
        newHead.Next = prevHead
    }
    
    return newHead
}
