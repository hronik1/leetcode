func singleNumber(nums []int) int {
    var bitmask uint64
    for _, num := range nums {
        bitmask = bitmask ^ uint64(num)
    }
    
    return int(bitmask)
}
