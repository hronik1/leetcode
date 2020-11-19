func plusOne(digits []int) []int {
    out := []int{}
    carry := 0
    l := len(digits)-1
    for i:=l; i >= 0; i-- {
        s := digits[i] + carry
        if i == l {
            s += 1
        }
        
        carry = s/10
        out = append([]int{s%10}, out...)
    }
    
    if carry == 1 {
        out = append([]int{carry}, out...)    
    }
    
    return out
}
