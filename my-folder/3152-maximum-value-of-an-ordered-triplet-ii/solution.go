func maximumTripletValue(nums []int) int64 {
    res := 0
    maxDiff := 0
    maxV := 0

    for _, v := range nums {
        curRes := maxDiff * v
        if curRes > res {
            res = curRes
        }

        diff := maxV - v
        if diff > maxDiff {
            maxDiff = diff
        }

        if v > maxV {
            maxV = v
        }
    }

    return int64(res)
}
