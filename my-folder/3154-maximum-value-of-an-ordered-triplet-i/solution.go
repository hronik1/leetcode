func maximumTripletValue(nums []int) int64 {
    max := 0
    for i := 0; i < len(nums)-2; i += 1 {
        for j := i+1; j < len(nums)-1; j += 1 {
            for k := j+1; k < len(nums); k += 1 {
                val := (nums[i] - nums[j]) * nums[k]
                if val > max {
                    max = val
                }
            }
        }   
    }
    
    return int64(max)
}
