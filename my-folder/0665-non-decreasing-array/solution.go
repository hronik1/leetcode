import "fmt"

func checkPossibility(nums []int) bool {
    count := 0
    
    for i := 0; i < len(nums)-1; i++ {
        if nums[i] > nums[i+1] {
            count++
            if count >= 2 {
                return false
            }
            
            if i == 0 {
                nums[i] = nums[i+1]
            } else if nums[i-1] <= nums[i+1] {
                nums[i] = nums[i-1]
            } else {
                nums[i+1] = nums[i]
            }
            
            
        }
    }

    return true
}

