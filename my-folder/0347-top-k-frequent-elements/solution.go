import "container/heap"

func topKFrequent(nums []int, k int) []int {
    frequencies := map[int]int{}
    
    for _, v := range nums {
        frequencies[v]++
    }
    
    // init heap
    h := DoubleHeap{}
    heap.Init(&h)
    for num, frequency := range frequencies {
        if len(h) < k {
            heap.Push(&h, Double{Item: num, Count: frequency})
        } else {
            if frequency > h[0].Count {
                heap.Pop(&h)
                heap.Push(&h, Double{Item: num, Count: frequency})
            }
        }
    }
    
    out := []int{}
    for _, v := range h {
        out = append(out, v.Item)
    }
    
    return out
}

type Double struct {
    Item int
    Count int
}

// An IntHeap is a min-heap of ints.
type DoubleHeap []Double

func (h DoubleHeap) Len() int           { return len(h) }
func (h DoubleHeap) Less(i, j int) bool { return h[i].Count < h[j].Count }
func (h DoubleHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *DoubleHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Double))
}

func (h *DoubleHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
