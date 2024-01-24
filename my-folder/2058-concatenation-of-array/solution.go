func getConcatenation(nums []int) []int {
    out := []int{}
    for i := 0; i < 2; i++ {
        for _, num := range nums {
            out = append(out, num)
        }
    }

    return out
}
