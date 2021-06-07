import "container/heap"

func findKthLargest(nums []int, k int) int {
    h := IntHeap{}
    heap.Init(&h)
    for _, v := range nums {
        if len(h) < k {
            heap.Push(&h, v)
        } else {
            if h[0] < v {
                _ = heap.Pop(&h)
                heap.Push(&h, v)
            }
        }
    }
    
    return h[0]
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
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
