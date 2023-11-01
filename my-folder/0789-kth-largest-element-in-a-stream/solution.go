import "container/heap"

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type KthLargest struct {
    h *IntHeap
    k int
}


func Constructor(k int, nums []int) KthLargest {
    h := &IntHeap{}
    heap.Init(h) 
    
    for _, v := range nums {
        if (*h).Len() == k {
            if v > (*h)[0] {
                heap.Pop(h)
                heap.Push(h, v)
            }
        } else {
            heap.Push(h, v)
        }
    }
    
    return KthLargest{h: h, k: k}
}


func (this *KthLargest) Add(val int) int {
    if (*this.h).Len() == this.k {
        if val > (*this.h)[0] {
                heap.Pop(this.h)
                heap.Push(this.h, val)
            }
        } else {
            heap.Push(this.h, val)
        }
    
    return (*this.h)[0]
}


/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
