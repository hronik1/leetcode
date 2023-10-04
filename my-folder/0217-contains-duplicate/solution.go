func containsDuplicate(nums []int) bool {
    d := map[int]bool{}
    for _, v := range nums {
        if _, ok := d[v]; ok {
            return true
        }
        
        d[v] = true
    }
    
    return false
}
