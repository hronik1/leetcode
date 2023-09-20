func findMaxConsecutiveOnes(nums []int) int {
    maxCount := 0
    curCount := 0
    for _, v := range nums {
        if v == 1 {
            curCount += 1
            if curCount > maxCount {
                maxCount = curCount
            }
        } else {
            curCount = 0
        }
    }
    
    return maxCount
}
