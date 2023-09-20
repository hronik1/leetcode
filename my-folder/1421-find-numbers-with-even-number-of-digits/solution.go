func findNumbers(nums []int) int {
    count := 0
    for _, v := range nums {
        if hasEvenDigitCount(v) {
            count += 1
        }
    }
    
    return count
}

func hasEvenDigitCount(v int) bool {
    digitCount := 0
    for v > 0 {
        v /= 10
        digitCount += 1
    }
    
    return digitCount%2 == 0
}
