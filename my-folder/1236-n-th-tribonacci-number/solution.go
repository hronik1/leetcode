func tribonacci(n int) int {
    if n < 2 {
        return n
    }
    
    if n == 2 {
        return 1
    }
    
    memo := map[int]int {
        0: 0,
        1: 1,
        2: 1,
    }
    
    return tribHelper(n, memo)
}

func tribHelper(n int, memo map[int]int) int {
    if v, ok := memo[n]; ok {
        return v
    }
    
    v := tribHelper(n-3, memo) + tribHelper(n-2, memo) + tribHelper(n-1, memo)
    memo[n] = v
    return v
}
