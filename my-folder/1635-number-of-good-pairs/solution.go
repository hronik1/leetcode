func numIdenticalPairs(nums []int) int {
    counts := make(map[int]int)
    for _, v := range nums {
        counts[v]++
    }
    
    out := 0
    for _, v := range counts {
        out += v*(v-1)/2
    }
    
    return out
}
