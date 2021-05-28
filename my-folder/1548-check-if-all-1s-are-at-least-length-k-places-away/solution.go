func kLengthApart(nums []int, k int) bool {
    space := -1
    for _, v := range nums {
        if space == -1 {
            if v == 1 {
                space = 0
            }
        } else {
            if v == 0 {
                space++
            } else if v == 1 {
                if space < k {
                    return false
                }
                
                space = 0
            }
        }
    }
    
    return true
}
