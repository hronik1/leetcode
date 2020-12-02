func fib(N int) int {
    if N < 2 {
        return N
    }
    
    memo := map[int]int {
        0: 0,
        1: 1,
    }
    
    return fibMemo(N, memo)
}

func fibMemo(N int, memo map[int]int) int {
    if v, ok := memo[N]; ok {
        return v
    }
    
    v := fibMemo(N-1, memo) + fibMemo(N-2, memo)
    memo[N] = v
    return v
}


