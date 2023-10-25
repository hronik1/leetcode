func findPeakElement(nums []int) int {   
    lo := 0
    hi := len(nums) - 1
    for lo < hi-2 {
        mid := lo + (hi-lo)/2
        if nums[mid-1] < nums[mid] {
            if nums[mid] > nums[mid+1] {
                return mid
            } else {
                lo = mid + 1
            }
        } else {
            hi = mid - 1
        } 
    }
    
    peak := nums[lo]
    peakI := lo
    for i := lo; i <= hi; i += 1 {
        v := nums[i]
        if v > peak {
            peak = v
            peakI = i
        }
    }
    
    return peakI
}
