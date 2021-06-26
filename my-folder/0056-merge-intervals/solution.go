func merge(intervals [][]int) [][]int {
    if len(intervals) == 0 {
        return [][]int{}
    }
    // sort intervals by start time, if most over laps, merge other wise add new entry
    sort.SliceStable(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    
    out := [][]int{intervals[0]}
    for i := 1; i < len(intervals); i++ {
        prev := out[len(out)-1]
        cur := intervals[i]
        if doesOverlap(prev, cur) {
            out[len(out)-1] = merg(prev, cur)
        } else {
            out = append(out, cur)
        }
    }
    
    return out
}

func doesOverlap(first []int, second []int) bool {
    return first[1] >= second[0]
}

func merg(first []int, second []int) []int {
    out := first
    if second[1] > out[1] {
        out[1] = second[1]
    }
    
    return out
}
