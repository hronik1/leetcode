func sortArray(nums []int) []int {
    return sortHelper(nums, 0, len(nums)-1)
}

func sortHelper(nums[]int, lo int, hi int) []int {
    if hi < lo {
        return []int{}
    }
    
    if lo == hi {
        return []int{nums[lo]}
    }
    
    mid := lo + (hi-lo)/2
    loRes := sortHelper(nums, lo, mid)
    hiRes := sortHelper(nums, mid+1, hi)
    
    res := []int{}
    i := 0
    j := 0
    for i < len(loRes) && j < len(hiRes) {
        if loRes[i] < hiRes[j] {
            res = append(res, loRes[i])
            i += 1
        } else {
            res = append(res, hiRes[j])
            j += 1
        }
    }
    
    //append rest of loRes
    for ;i < len(loRes); i += 1 {
        res = append(res, loRes[i])

    }
    //append rest of hiRes
    for ; j < len(hiRes); j += 1 {
        res = append(res, hiRes[j])
    }
    
    return res
}
