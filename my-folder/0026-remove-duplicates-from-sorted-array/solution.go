func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    uniqueCount := 1
    prevNum := nums[0]
    for _, num := range nums {
        if num != prevNum {
            prevNum = num
            nums[uniqueCount] = num
            uniqueCount++
        }        
    }
    
    return uniqueCount
}
