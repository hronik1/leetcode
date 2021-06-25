/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    h := NodeHeap{}
    for _, node := range lists {
        if node != nil {
            h = append(h, node)
        }
    }
	heap.Init(&h)
    
    var head *ListNode
    var prev *ListNode
    for len(h) > 0 {
        cur := heap.Pop(&h).(*ListNode)
        if cur.Next != nil {
            heap.Push(&h, cur.Next)
        }
        
        if head == nil {
            head = cur
        } else {
            prev.Next = cur
        }
        
        prev = cur
        prev.Next = nil
    }
    
    return head
}

// An NodeHeap is a min-heap of ints.
type NodeHeap []*ListNode

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h NodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *NodeHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*ListNode))
}

func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
