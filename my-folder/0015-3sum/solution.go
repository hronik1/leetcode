import "sort"
func threeSum(nums []int) [][]int {
    out := [][]int{}
    sort.Ints(nums)
    for i := 0; i < len(nums)-2; i++ {
        for j, k := i+1, len(nums)-1; j < k; {
            sum := nums[i] + nums[j] + nums[k]
            if sum == 0 {
                out = append(out, []int{nums[i], nums[j], nums[k]})
                j++
                for ; j < k && nums[j] == nums[j-1]; j++ { }
                k--
                for ; j < k  && nums[k] == nums[k+1]; k-- { }
            } else if sum < 0 {
                j++
                for ; j < k && nums[j] == nums[j-1]; j++ { }
            } else {
                k--
                for ; j < k  && nums[k] == nums[k+1]; k-- { }
            }
        }
        
        for ; i < len(nums)-1 && nums[i] == nums[i+1]; i++ {}
    }

    return out
}
