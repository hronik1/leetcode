/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
    newHead := head.Next
    newNext := swapPairs(newHead.Next)
    newHead.Next = head
    head.Next = newNext
    
    return newHead
}
