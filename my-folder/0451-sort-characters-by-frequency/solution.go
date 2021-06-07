import "container/heap"
import "strings"

func frequencySort(s string) string {
    frequencies := map[rune]int{}
    for _, r := range s {
        frequencies[r]++
    }
    
    h := &DoubleHeap{}
    for r, count := range frequencies {
        heap.Push(h, Double{Item:r, Count:count})
    }
    
    out := strings.Builder{}
    for len(*h) > 0 {
        item := heap.Pop(h).(Double)
        for i := 0; i < item.Count; i++ {
            _, _ = out.WriteRune(item.Item)
        }
    }
    
    return out.String()
}

type Double struct {
    Item rune
    Count int
}

// An IntHeap is a min-heap of ints.
type DoubleHeap []Double

func (h DoubleHeap) Len() int           { return len(h) }
func (h DoubleHeap) Less(i, j int) bool { return h[i].Count > h[j].Count }
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
