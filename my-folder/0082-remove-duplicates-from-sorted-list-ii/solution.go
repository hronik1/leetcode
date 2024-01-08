/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    var newHead *ListNode
    var prev *ListNode
    for cur := head; cur != nil; cur = cur.Next {
        if cur.Next != nil && cur.Val == cur.Next.Val {
            for ; cur.Next != nil && cur.Val == cur.Next.Val; cur = cur.Next {}
        } else {
            if newHead == nil {
                newHead = cur
            } else {
                prev.Next = cur
            }

            prev = cur
        }
    }

    if prev != nil {
        prev.Next = nil
    }

    return newHead
}
