func singleNumber(nums []int) int {
    d := map[int]bool{}
    for _, v := range nums {
        if _, ok := d[v]; ok {
            delete(d, v)
        } else {
            d[v] = true
        }
    }
    
    for k, _ := range d {
        return k
    }
    
    return -1
}
