import "math"
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
    // scan into slice
    var reversed []int
    cur := head
    for (cur != nil) {
        reversed = append([]int{cur.Val}, reversed...)
        cur = cur.Next
    }
    
    total := float64(0)
    for i, bit := range reversed {
        total += float64(bit) * math.Pow(float64(2), float64(i))
    }
    
    return int(total)
}
