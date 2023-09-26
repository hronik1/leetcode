/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    seen := map[*ListNode]struct{}{}
    
    for curA := headA; curA != nil; curA = curA.Next {
        seen[curA] = struct{}{}
    }

    for curB := headB; curB!= nil; curB = curB.Next {
        if _, ok := seen[curB]; ok {
            return curB
        }
    }

    return nil
}
