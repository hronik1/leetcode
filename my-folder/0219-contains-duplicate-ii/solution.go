func containsNearbyDuplicate(nums []int, k int) bool {
    indices := map[int][]int{}
    for i, num := range nums {
        indices[num] = append(indices[num], i)
    }
    
    for _, index := range indices {
        for i := 0; i < len(index)-1; i += 1 {
            if index[i+1] - index[i] <= k {
                return true
            }
        }
    }
    
    return false
}
