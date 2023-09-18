/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    cur := head
    len := 0
    for cur != nil {
        cur = cur.Next
        len += 1
    }

    midI := len/2
    cur = head
    for i := 1; i <= midI; i++ {
        cur = cur.Next
    }

    return cur
}
