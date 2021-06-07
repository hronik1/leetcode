import "container/heap"
import "math"

func kClosest(points [][]int, k int) [][]int {
    h := PointHeap{}
    heap.Init(&h)
    
    for _, point := range points {
        dist := math.Pow(math.Pow(float64(point[0]), 2.0) + math.Pow(float64(point[1]), 2.0), 0.5)
        if len(h) == k {
            if dist > h[0].Distance {
                continue
            }
            
            _ = heap.Pop(&h)
        }
        
        heap.Push(&h, Point{Point:point, Distance:dist})
    }
    
    out := [][]int{}
    for _, v := range h {
        out = append(out, v.Point)
    }
    
    return out
    
}

type Point struct {
    Point []int
    Distance float64
}

// An IntHeap is a min-heap of ints.
type PointHeap []Point

func (h PointHeap) Len() int           { return len(h) }
func (h PointHeap) Less(i, j int) bool { return h[i].Distance > h[j].Distance }
func (h PointHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PointHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Point))
}

func (h *PointHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
