/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
    var partitionHead *ListNode
    var partitionHeadPrev *ListNode 
    for cur := head; cur != nil; cur = cur.Next {
        if cur.Val >= x {
            partitionHead = cur
            break
        }

        partitionHeadPrev = cur
    }

    if partitionHead == nil {
        return head
    }

    newHead := head
    for cur := partitionHead; cur != nil && cur.Next != nil; {
        fmt.Printf("cur.Val:%d\n", cur.Val)
        if cur.Next.Val < x {
            next := cur.Next
            if partitionHeadPrev == nil {
                newHead = next
            } else {
                partitionHeadPrev.Next = next
            }
            
            cur.Next = next.Next
            next.Next = partitionHead
            partitionHeadPrev = next
        } else {
            cur = cur.Next
        }
    }

    return newHead
}
