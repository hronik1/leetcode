import "cmp"
import "slices"

func merge(intervals [][]int) [][]int {
    ret := [][]int{}
    if len(intervals) == 0 {
        return ret
    }

    slices.SortFunc(intervals, func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	})

    
    curStart := intervals[0][0]
    curEnd := intervals[0][1]
    for i := 1; i < len(intervals); i += 1 {
        if intervals[i][0] <= curEnd {
            if intervals[i][1] > curEnd {
                curEnd = intervals[i][1]
            }

        } else {
            ret = append(ret, []int{curStart, curEnd})
            curStart = intervals[i][0]
            curEnd = intervals[i][1]
        }
    }

    ret = append(ret, []int{curStart, curEnd})

    return ret
}
