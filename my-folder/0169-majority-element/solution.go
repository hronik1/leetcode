func majorityElement(nums []int) int {
    if len(nums) < 3 {
        return nums[0]
    }
    
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] < nums[j]
    })
    
    curCount := 0
    curElement := nums[0]
    maxCount := 0
    maxElement := nums[0]
    for _, v := range nums {
        if v == curElement {
            curCount++
        } else {
            curCount = 1
            curElement = v
        }
        
        if curCount > maxCount {
            maxCount = curCount
            maxElement = curElement
        }
    }
    
    return maxElement
}
