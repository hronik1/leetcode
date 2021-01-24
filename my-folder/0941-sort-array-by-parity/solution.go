func sortArrayByParity(A []int) []int {
    odd := []int{}
    even := []int{}
    for _, v := range A {
        if v%2 == 0 {
            even = append(even, v)
        } else {
            odd = append(odd, v)
        }
    }
    
    even = append(even, odd...)
    return even
}
