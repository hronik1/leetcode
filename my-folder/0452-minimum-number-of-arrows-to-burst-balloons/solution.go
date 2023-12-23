import "sort"

func findMinArrowShots(points [][]int) int {
    sort.Slice(points, func (i, j int) bool { return points[i][0] < points[j][0] })

    arrowCount := 1
    minEnd := points[0][1]
    for i := 1; i < len(points); i++ {
        v := points[i]
        if v[0] <= minEnd {
            if v[1] < minEnd {
                minEnd = v[1]
            }
        } else {
            arrowCount++
            minEnd = v[1] 
        }
    }

    return arrowCount    
}
