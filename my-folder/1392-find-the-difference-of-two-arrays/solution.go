func findDifference(nums1 []int, nums2 []int) [][]int {
    set1 := settify(nums1)
    set2 := settify(nums2)

    onlyIn1 := onlyInFirst(set1, set2)
    onlyIn2 := onlyInFirst(set2, set1)

    return [][]int{onlyIn1, onlyIn2}
}

func settify(nums []int) map[int]bool {
    m := map[int]bool{}
    for _, v := range nums {
        m[v] = true
    }

    return m
}

func onlyInFirst(set1 map[int]bool, set2 map[int]bool) []int {
    out := []int{}
    for k, _ := range set1 {
        if _, ok := set2[k]; !ok {
            out = append(out, k)
        }
    }

    return out
}
