import "math/rand"
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
    head *ListNode
}


/** @param head The linked list's head.
        Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
    return Solution {
        head,
    }
}


/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
    cur := this.head
    out := this.head.Val
    for seen := 1; cur != nil; seen++ {
        if rand.Intn(seen) == 0 {
            out = cur.Val
        }
        cur = cur.Next
    }
    
    return out
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
