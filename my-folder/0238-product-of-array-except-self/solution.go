func productExceptSelf(nums []int) []int {
    leftProduct := make([]int, len(nums), len(nums))
    leftProduct[0] = 1
    for i := 1; i < len(nums); i += 1 {
        leftProduct[i] = leftProduct[i-1] * nums[i-1]
    }

    rightProduct := make([]int, len(nums), len(nums))
    rightProduct[len(nums) - 1] = 1
    for i := len(nums) - 2; i >= 0; i -= 1 {
        rightProduct[i] = rightProduct[i+1] * nums[i+1]
    }
    
    out := make([]int, len(nums), len(nums))
    for i := 0; i < len(nums); i += 1 {
        out[i] = leftProduct[i] * rightProduct[i]
    }
    
    return out
}


