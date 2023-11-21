/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    carry := 0
    head := &ListNode{}
    prev := head
    for l1 != nil || l2 != nil {
        sum := carry
        if l1 != nil {
            sum += l1.Val
            l1 = l1.Next
        }

        if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
        }

        cur := &ListNode{Val:sum%10}
        carry = sum/10
        prev.Next = cur
        prev = cur
    }

    if carry > 0 {
        cur := &ListNode{Val:carry}
        prev.Next = cur
    }

    return head.Next // treating head as Sentinel
}
