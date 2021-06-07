import "container/heap"

type KthLargest struct {
    H IntHeap
    K int
}


func Constructor(k int, nums []int) KthLargest {
    h := IntHeap{}
    heap.Init(&h)
    out := KthLargest{
        H: h,
        K: k,
    }
    
    for _, v := range nums {
        _ = out.Add(v)
    }
    
    return out
}


func (this *KthLargest) Add(val int) int {
    if len(this.H) < this.K {
        heap.Push(&this.H, val)
    } else {
        if this.H[0] < val {
            _ = heap.Pop(&this.H)
            heap.Push(&this.H, val)
        }
    }
    
    return this.H[0]
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
/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
