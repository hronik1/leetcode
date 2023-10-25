func findMin(nums []int) int {
    lo := 0
    hi := len(nums) - 1
    for lo <= hi - 2 { // guarantees unique lo, mid, and hi indices, which makes boundary conditions easier
        mid := lo + (hi-lo)/2
        if nums[mid-1] > nums[mid] && nums[mid+1] > nums[mid] {
            return nums[mid]
        } else if nums[hi] > nums[mid] {
            hi = mid - 1
        } else {
            lo = mid + 1
        }
    }
    
    min := nums[lo]
    for i := lo; i <= hi; i += 1 {
        if nums[i] < min {
            min = nums[i]
        }
    }
    
    return min
}
