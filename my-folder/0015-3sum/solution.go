import "sort"
func threeSum(nums []int) [][]int {
    out := [][]int{}
    seen := map[[3]int]bool{}
    if len(nums) < 3 {
        return out
    }
    
    sort.Ints(nums)
    for i := 0; i < len(nums)-2; i++ {
        for j, k := i+1, len(nums)-1; j < k; {
            sum := nums[i] + nums[j] + nums[k]
            if sum == 0 {
                seenKey := [3]int{nums[i],nums[j],nums[k]}
                if _, ok := seen[seenKey]; !ok {
                    out = append(out, []int{nums[i],nums[j],nums[k]})
                    seen[seenKey] = true
                }
                
                k--
                j++
            } else if sum > 0 {
                k--
            } else {
                j++
            }
        }
    }
    
    
    return out
}
