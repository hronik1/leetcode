/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    
    if head.Next == nil {
        return head
    }
    
    nextPair := swapPairs(head.Next.Next)
    newHead := head.Next
    newHead.Next = head
    head.Next = nextPair
    
    return newHead
}
