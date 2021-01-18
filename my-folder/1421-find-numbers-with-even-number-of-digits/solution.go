func findNumbers(nums []int) int {
    count := 0
    for _, num := range nums {
        if hasEvenDigits(num) {
            count++
        }
    }
    
    return count
}

func hasEvenDigits(num int) bool {
    count := 0
    for num > 0 {
        count++
        num /= 10
    }
    
    return count%2 == 0
}
