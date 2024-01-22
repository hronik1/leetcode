func findErrorNums(nums []int) []int {
    counts := make([]int, len(nums), len(nums))
    for _, v := range nums {
        counts[v-1]++
    }

    out := make([]int, 2, 2)
    for i, count := range counts {
        if count == 2 {
            out[0] = i+1
        }
        if count == 0 {
            out[1] = i+1
        }
    }
    return out
}
