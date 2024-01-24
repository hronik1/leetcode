func buildArray(nums []int) []int {
    out := []int{}
    for i := 0; i < len(nums); i++ {
        out = append(out, nums[nums[i]])
    }

    return out
}
