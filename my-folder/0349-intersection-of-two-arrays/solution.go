func intersection(nums1 []int, nums2 []int) []int {
    first := map[int]bool{}
    for _, v := range nums1 {
        if _, ok := first[v]; !ok {
            first[v] = true
        }
    }
    
    intersection := map[int]bool{}
    for _, v := range nums2 {
        if _, ok := first[v]; ok {
            intersection[v] = true
        }
    }
    
    res := make([]int, 0, len(intersection))
    for k, _ := range intersection {
        res = append(res, k)
    }
    
    return res
}
