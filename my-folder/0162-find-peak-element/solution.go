func findPeakElement(nums []int) int {
    lo, hi := 0, len(nums)-1
    for lo < hi-1 {
        mid := lo + (hi-lo)/2
        
        if mid == 0 || nums[mid] > nums[mid-1] {
            lo = mid
        } else {
            hi = mid
        }
    }
    
    max := lo
    for i := lo; i <= hi; i++ {
        if nums[i] > nums[max] {
            max = i
        }
    }
    
    return max
}
