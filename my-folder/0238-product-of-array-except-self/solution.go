func productExceptSelf(nums []int) []int {
    leftProd := make([]int, len(nums))
    rightProd := make([]int, len(nums))
    leftProd[0] = 1
    rightProd[len(nums)-1] = 1
    for i := 1; i < len(nums); i++ {
        leftProd[i] = leftProd[i-1]*nums[i-1]
        rightProd[len(nums)-1-i] = rightProd[len(nums)-i]*nums[len(nums)-i]
    }
    
    out := []int{}
    for i := 0; i < len(nums); i++ {
        out = append(out, leftProd[i]*rightProd[i])
    }
    
    return out
}
