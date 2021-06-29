func searchRange(nums []int, target int) []int {
    if len(nums) == 0 {
        return []int{-1,-1}
    }
    
    // find start index
    start := -1
    lo, hi := 0, len(nums)-1
    for lo <= hi {
        mid := lo + (hi-lo)/2
        
        if nums[mid] == target {
            start = mid
            hi = mid-1
        } else if nums[mid] < target {
            lo = mid+1
        } else {
            hi = mid-1
        }
    }
    
    if start == -1 {
        return []int{-1,-1}
    }
    
    end := start
    lo, hi = start, len(nums)-1
    for lo <= hi {
        mid := lo + (hi-lo)/2
        
        if nums[mid] == target {
            end = mid
            lo = mid+1
        } else if nums[mid] < target  {
            lo = mid+1
        } else {
            hi = mid-1
        }
    }
    
    
    return []int{start, end}
}
