func removeDuplicates(nums []int) int {
    if len(nums) < 1 {
        return 0
    }
    
    newI := 1
    lastSeen := nums[0]
    lastSeenCount := 1
    for i := 1; i < len(nums); i++ {
        if nums[i] != lastSeen {
            nums[newI] = nums[i]
            lastSeen = nums[i]
            lastSeenCount = 1
            newI++
        } else if nums[i] == lastSeen && lastSeenCount < 2 {
            nums[newI] = nums[i]
            lastSeenCount++
            newI++
        }
    }
    
    return newI
}
