func permute(nums []int) [][]int {
    res := [][]int{}
    if len(nums) > 0  {
        res = helper(len(nums), nums, res)
    }

    return res
}

func helper(k int, nums []int, res [][]int) [][]int {
    if k == 0 {
        res = appendPerm(nums, res)
        return res
    }
    
    for i := 0; i < k; i++ {
        res = helper(k-1, nums, res)
        
        if k%2 == 1 {
            swap(nums, i, k-1)
        } else {
            swap(nums, 0, k-1)
        }
    }
    
    return res
}

func appendPerm(nums []int, res [][]int) [][]int{
    cp := make([]int, len(nums))
    _ = copy(cp, nums)
    res = append(res, cp)

    return res
}

func swap(nums []int, i int, j int) {
    tmp := nums[i]
    nums[i] = nums[j]
    nums[j] = tmp
}
