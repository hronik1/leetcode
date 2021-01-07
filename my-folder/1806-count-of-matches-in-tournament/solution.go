func numberOfMatches(n int) int {
    count := 0
    for n > 1 {
        skip := 0
        if n%2 == 1 {
            skip++
        }
        
        n/=2
        count += n
        n += skip
    }
    
    return count
}
