/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
    len := listLen(head)
    if len == 1 {
        return nil
    }

    prev := head
    cur := head
    mid := len/2
    for i := 0; i < mid; i++ {
        prev = cur
        cur = cur.Next
    }

    prev.Next = cur.Next

    return head
}

func listLen(head *ListNode) int {
    len := 0
    for ; head != nil; len, head = len+1, head.Next {}
    return len
}
