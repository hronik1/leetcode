func countPrimes(n int) int {
    if n <= 2 {
        return 0
    }
    count := 1
    primes := []int{2}
    for i := 3; i < n; i++ {
        divisible := false
        for j := 0; j < len(primes) && primes[j]*primes[j] <= i; j++ {
            if i%primes[j] == 0 {
                divisible = true
                break
            }
        }
        
        if !divisible {
            count++
            primes = append(primes,i)
        }
    }
    
    return count
}
