import "math"
func minTimeToVisitAllPoints(points [][]int) int {
    d := 0
    for i := 1; i < len(points); i++ {
        d += dist(points[i-1], points[i])
    }
    
    return d
}

func dist(p0 []int, p1 []int) int {
    d0 := math.Abs(float64(p1[0]-p0[0]))
    d1 := math.Abs(float64(p1[1]-p0[1]))
    return int(math.Max(d0, d1))
    
}
