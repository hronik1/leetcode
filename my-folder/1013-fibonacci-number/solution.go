func fib(n int) int {
    if n == 0 || n == 1 {
        return n
    }
    
    memo := []int{0, 1}
    
    return fibHelper(n, memo)
}

func fibHelper(n int, memo []int) int {
    if len(memo) > n {
        return memo[n]
    }
    
    f := fibHelper(n-1, memo) + fibHelper(n-2, memo)
    memo = append(memo, f)
    
    return f
}
