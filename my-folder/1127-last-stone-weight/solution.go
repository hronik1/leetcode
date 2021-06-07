import "container/heap"

func lastStoneWeight(stones []int) int {
    h := IntHeap{}
    for _, v := range stones {
        h = append(h, v)
    }
    heap.Init(&h)
    for len(h) > 1 {
        y := heap.Pop(&h).(int)
        x := heap.Pop(&h).(int)
        newWeight := y-x
        if newWeight > 0 {
            heap.Push(&h, newWeight)
        }
    }
    
    if len(h) == 1 {
        return h[0]
    }
    
    return 0
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
