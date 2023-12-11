func combinationSum3(k int, n int) [][]int {
    return combinationSumHelper(k, n, 9)
}

func combinationSumHelper(k int, n int, hi int) [][]int {
    out := [][]int{}
    if hi == 0 {
        return out
    }

    if k == 1 && hi == n {
        out = append(out, []int{hi})
        return out
    }

    if hi < n {
        included := combinationSumHelper(k-1, n-hi, hi-1)
        for _, v := range included {
            v = append(v, hi)
            out = append(out, v)
        }
    }

    excluded := combinationSumHelper(k, n, hi-1)
    out = append(out, excluded...)

    return out

}
