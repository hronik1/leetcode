func shuffle(nums []int, n int) []int {
    double := n+n
    out := make([]int, double)
    for i := 0; i < n; i++ {
        newi := i+i
        out[newi] = nums[i]
        out[newi+1] = nums[i+n]
    }
    
    return out
}
