import "fmt"

func insert(intervals [][]int, newInterval []int) [][]int {
    before_i := -1
    for i := 0;  i < len(intervals); i++ {
        if intervals[i][1] < newInterval[0] {
            before_i = i
        } else {
            break
        }
    }
    
    after_i := len(intervals)
    for i := len(intervals) - 1; i >= 0; i-- {
        if intervals[i][0] > newInterval[1] {
            after_i = i
        } else {
            break
        }
    }
    
    output := [][]int{}
    for i := 0; i <= before_i; i++ {
        output = append(output, intervals[i])
    }
    
    // append new stuff
    output = append(output, merge(intervals, newInterval, before_i+1, after_i-1))
                    
    for i:= after_i; i < len(intervals); i++ {
        output = append(output, intervals[i])
    }
    
    return output
}
            
func doesOverlap(interval1 []int, interval2 []int) bool{
    return !(interval1[1] < interval2[0] || interval1[0] > interval2[1])
}

func merge(intervals [][]int, newInterval []int, start_i int, end_i int) []int{
    merged := []int{newInterval[0], newInterval[1]}
    if start_i >= 0 && start_i < len(intervals) && intervals[start_i][0] < newInterval[0]{
        merged[0] = intervals[start_i][0]
    }
    
    if end_i >= 0 && end_i < len(intervals) && intervals[end_i][1] > newInterval[1]{
        merged[1] = intervals[end_i][1]
    }
    return merged
}

