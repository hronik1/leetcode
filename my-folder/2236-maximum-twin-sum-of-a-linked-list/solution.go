/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
    vals := []int{}
    for cur := head; cur != nil; cur = cur.Next {
        vals = append(vals, cur.Val)
    }

    best := 0
    mid := len(vals)/2
    for i := 0; i < mid; i++ {
        cur := vals[i] + vals[len(vals)-1-i]
        if cur > best {
            best = cur
        }
    }

    return best
}
