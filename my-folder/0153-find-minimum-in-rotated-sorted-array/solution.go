func findMin(nums []int) int {
    lo, hi := 0, len(nums)-1
    
    for hi >= lo+4 {
        if nums[lo] <= nums[hi] {
            return nums[lo]
        }
        
        mid := lo + (hi-lo)/2
        if nums[mid] < nums[mid-1] {
            return nums[mid]
        }
        
        if nums[lo] < nums[mid-1] {
            lo = mid+1
        } else {
            hi = mid-1
        }
    }
    
    lowest := nums[lo]
    for ; lo <= hi; lo++ {
        if nums[lo] < lowest {
            lowest = nums[lo]
        }
    }
    
    return lowest
}
