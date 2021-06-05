func subsets(nums []int) [][]int {
    res := [][]int{}
    top := int(math.Pow(2, float64(len(nums))))

    masks := []int{}
    for dec := 1; dec < top; dec*=2 {
        masks = append(masks, dec)
    }
    
    for i := 0; i < top; i++ {
        v := []int{}
        for j, mask := range masks {
            if i & mask != 0 {
                v = append(v, nums[j])
            }
        }
        
        res = append(res, v)
    }
    
    return res
}
