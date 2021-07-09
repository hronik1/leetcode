func smallerNumbersThanCurrent(nums []int) []int {
    if len(nums) == 0 {
        return []int{}
    }
    
    cp := make([]int, len(nums))
    for i, v := range nums {
        cp[i] = v
    }
    
    sort.Slice(cp, func(i, j int) bool {
        return cp[i] > cp[j]
    })
    
    remainingCounts := map[int]int{}
    for i := 0; i < len(cp); i++ {
        for i+1 < len(nums) && cp[i] == cp[i+1] {
            i++
        }
        
        remainingCounts[cp[i]] = len(nums)-1-i
    }
    
    out := make([]int, len(nums))
    for i, v := range nums {
        out[i] = remainingCounts[v]
    }
    
    return out
}
