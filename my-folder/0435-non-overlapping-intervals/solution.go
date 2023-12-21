import "fmt"
import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(i int, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    removalCount := 0
    end := intervals[0][1]
    for i := 1; i < len(intervals); i++ {
        v := intervals[i]
        if v[0] < end {
            removalCount++
            if v[1] < end {
                end = v[1]
            }
        } else {
            end = v[1]
        }
    }

    return removalCount
}
