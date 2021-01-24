func countVowelStrings(n int) int {
    return count(n, 0)
}

func count(n int, offset int) int {
    if n == 1 {
        return 5 - offset
    }
    
    c := 0 
    for i := offset; i < 5; i++ {
        c += count(n-1, i)
    }
    
    return c
}

