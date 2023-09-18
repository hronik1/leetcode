func runningSum(nums []int) []int {
    if len(nums) == 0 {
        return []int{}
    }

    l := len(nums)
    out := make([]int, l)
    out[0] = nums[0]
	for i := 1; i < l; i++ {
		out[i] = nums[i] + out[i-1]
	}

    return out
}
