func runningSum(nums []int) []int {
    if len(nums) < 1 {
        return []int{}
    }
    
    sums := []int{nums[0]}
    for i := 1; i < len(nums); i++ {
        sums = append(sums, sums[i-1] + nums[i])
    }
    
    return sums
}
