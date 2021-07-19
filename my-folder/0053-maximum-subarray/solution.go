func maxSubArray(nums []int) int {
    curSum := 0
    bestSum := nums[0]
    for i := 0; i < len(nums); i++ {
        curSum += nums[i]
        if curSum > bestSum {
            bestSum = curSum
        }
        
        if curSum < 0 {
            curSum = 0
        }
    }
    
    return bestSum
}
