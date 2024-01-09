import "fmt"
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil {
        return nil
    }

    length := 0
    for cur := head; cur != nil; cur = cur.Next {
        length++
    }

    rotateAmount := k%length
    if rotateAmount == 0 {
        return head
    }

    newHeadI := length - rotateAmount
    newTailI := newHeadI - 1
    var tail *ListNode
    var newHead *ListNode
    var newTail *ListNode
    for i, cur := 0, head; i < length; i, cur = i+1, cur.Next {
        
        if i == newTailI {
            newTail = cur
        }
        if i == newHeadI {
            newHead = cur
        }
        if i == length-1 {
            tail = cur
        }
    }


    newTail.Next = nil
    tail.Next = head

    return newHead 
}
