import "container/heap"

func maximumScore(a int, b int, c int) int {
    h := IntHeap{a, b, c}
    heap.Init(&h)
    score := 0
    
    for len(h) > 1 {
        first := heap.Pop(&h).(int)
        second := heap.Pop(&h).(int)
        score++
        
        if first > 1 {
            heap.Push(&h, first-1)
        }
        
        if second > 1 {
            heap.Push(&h, second-1)
        }
    }
    
    return score
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
