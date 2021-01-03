import "sort"
func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
    out := []bool{}
    for i := 0; i < len(l); i++ {
        out = append(out, helper(nums, l[i], r[i]))
    }
    
    return out
}

func helper(nums []int, lo int, hi int) bool {
    if hi - lo < 2 {
        return true
    }

    s := nums[lo:hi+1]
    scopy := make([]int, len(s))
    copy(scopy, s)
    sort.Slice(scopy, func(i int, j int) bool {
        return scopy[i] < scopy[j]
    })

    diff := scopy[1] - scopy[0]
    for j := 2; j < len(scopy); j++ {
        if scopy[j] - scopy[j-1] != diff {
            return false
        }
    }
    
    return true
}
