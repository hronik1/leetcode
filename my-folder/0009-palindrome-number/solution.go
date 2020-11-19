func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    
    digits := []int{}
    for r := x; r > 0; r/=10 {
        digits = append([]int{r%10}, digits...)
    }
    
    for i := 0; i < len(digits)/2; i++ {
        if digits[i] != digits[len(digits)-1-i] {
            return false
        }
    }
    
    return true
}
