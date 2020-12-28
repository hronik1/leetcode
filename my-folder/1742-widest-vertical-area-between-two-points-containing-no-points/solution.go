import "sort"

func maxWidthOfVerticalArea(points [][]int) int {
    sort.Slice(points, func(i, j int) bool {
        return points[i][0] < points[j][0]
    })
    
    maxWidth := 0
    prev := points[0][0]
    for i := 1; i < len(points); i++ {
        if currWidth := points[i][0] - prev; currWidth > maxWidth {
            maxWidth = currWidth
        }
        
        prev = points[i][0]
    }
    
    return maxWidth
}
