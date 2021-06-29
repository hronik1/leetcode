func search(nums []int, target int) int {
    if len(nums) == 0 {
        return -1
    }
    
    lo, hi := 0, len(nums)-1
    // want to ensure distinct [lo..mid-1,mid,mid+1..hi]
    for hi >= lo+4 {
        mid := lo + (hi-lo)/2
        
        if nums[mid] == target {
            return mid
        }
        
        if nums[lo] < nums[mid-1] {
            if nums[lo] <= target && target <= nums[mid-1] {
                hi = mid-1
            } else {
                lo = mid+1
            }
        } else {
            if nums[mid+1] <= target && target <= nums[hi] {
                lo = mid+1
            } else {
                hi = mid-1
            }
        }
    }
    
    for i := lo; i <= hi; i++ {
        if nums[i] == target {
            return i
        }
    }
    
    return -1
}
