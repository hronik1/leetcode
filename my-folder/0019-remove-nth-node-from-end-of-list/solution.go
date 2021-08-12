/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    l := 0
    node := head
    for node != nil {
        node = node.Next
        l++
    }
    
    iToRemove := l-n
    if iToRemove == 0 {
        out := head.Next
        head.Next = nil
        return out
    }
    
    node = head
    var prev *ListNode
    for i := 0; i <= iToRemove; i++ {
        if i == iToRemove {
            prev.Next = node.Next
            node.Next = nil
        } else {
            prev = node
            node = node.Next
        }
    }
    
    return head
}
