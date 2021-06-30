func intersection(nums1 []int, nums2 []int) []int {
    vals1 := toMap(nums1)
    vals2 := toMap(nums2)
    
    out := []int{}
    for k, _ := range vals1 {
        if _, ok := vals2[k]; ok {
            out = append(out, k)
        }
    }
    
    return out
}

func toMap(nums []int) map[int]bool {
    out := map[int]bool{}
    for _, v := range nums {
        out[v] = true
    }
    
    return out
}
