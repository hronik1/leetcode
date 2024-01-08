/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    var beforeHead, reverseHead, reverseTail *ListNode
    for i, cur := 1, head; i <= right; i, cur = i+1, cur.Next {
        if i == left-1 {
            beforeHead = cur
        } else if i == left {
            reverseHead = cur
        }
        
        if i == right {
            reverseTail = cur
        }
    }

    reverse(reverseHead, reverseTail)
    if beforeHead == nil {
        return reverseTail
    }

    beforeHead.Next = reverseTail
    return head
}

func reverse(head *ListNode, tail *ListNode) {
    newHead := head
    for head.Next != nil && newHead != tail {
        temp := head.Next
        head.Next = temp.Next
        temp.Next = newHead
        newHead = temp
    }
}
