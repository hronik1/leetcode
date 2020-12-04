func kthFactor(n int, k int) int {
    factorCount := 0
    for i := 1; i <= n; i++ {
        if n%i == 0 {
            factorCount++
        }
        
        if factorCount == k {
            return i
        }
    }
    
    return -1
}
