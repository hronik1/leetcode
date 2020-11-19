func search(nums []int, target int) int {
    lo, hi := 0, len(nums)-1
    for hi >= lo + 4 {
        mid := lo + (hi-lo)/2
        if nums[mid] == target {
            return mid
        }
        
        sortedLo, sortedHi := lo, mid-1
        otherLo, otherHi := mid+1, hi
        if nums[sortedLo] > nums[sortedHi] {
            sortedLo, sortedHi = mid+1, hi
            otherLo, otherHi = lo, mid-1
        }
        
        if target >= nums[sortedLo] && target <= nums[sortedHi] {
            lo, hi = sortedLo, sortedHi
        } else {
            lo, hi = otherLo, otherHi
        }
    }
    
    for ;lo <= hi; lo++ {
        if nums[lo] == target {
            return lo
        }
    }
    
    return -1
}
