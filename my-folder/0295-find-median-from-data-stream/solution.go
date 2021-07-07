// An IntHeap is a min-heap of ints.
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// An IntHeap is a min-heap of ints.
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MedianFinder struct {
    topHalf *MinHeap
    bottomHalf *MaxHeap
}


/** initialize your data structure here. */
func Constructor() MedianFinder {
    topHalf := &MinHeap{}
    heap.Init(topHalf)
    bottomHalf := &MaxHeap{}
    heap.Init(bottomHalf)
    
    return MedianFinder {
        topHalf: topHalf,
        bottomHalf: bottomHalf,
    }
}


func (this *MedianFinder) AddNum(num int)  {
    if len(*this.topHalf) == 0 && len(*this.bottomHalf) == 0 {
        heap.Push(this.bottomHalf, num)
        return
    }
     
    if len(*this.topHalf) > len(*this.bottomHalf) {
        if num > (*this.topHalf)[0] {
            heap.Push(this.bottomHalf, (*this.topHalf)[0])
            heap.Pop(this.topHalf)
            heap.Push(this.topHalf, num)
        } else {
            heap.Push(this.bottomHalf, num)
        }
    } else if len(*this.bottomHalf) > len(*this.topHalf) {
        if  num < (*this.bottomHalf)[0] {
            heap.Push(this.topHalf, (*this.bottomHalf)[0])
            heap.Pop(this.bottomHalf)
            heap.Push(this.bottomHalf, num)
        } else {
            heap.Push(this.topHalf, num)
        }
    } else {
        if num <= (*this.bottomHalf)[0] {
            heap.Push(this.bottomHalf, num)
        } else {
            heap.Push(this.topHalf, num)
        }  
    }
}


func (this *MedianFinder) FindMedian() float64 {
    if len(*this.topHalf) > len(*this.bottomHalf) {
        return float64((*this.topHalf)[0])
    } else if len(*this.bottomHalf) > len(*this.topHalf) {
        return float64((*this.bottomHalf)[0])
    }
    
    return float64((*this.topHalf)[0] + (*this.bottomHalf)[0])/2.0
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
