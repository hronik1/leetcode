func climbStairs(n int) int {
    return climbStairsDP(n, map[int]int{})
}

func climbStairsDP(n int, memo map[int]int) int {
    if res, ok := memo[n]; ok {
        return res
    }
    
    if n < 3 {
        memo[n] = n
        return n
    }
    
    out := climbStairsDP(n-1, memo) + climbStairsDP(n-2, memo)
    memo[n] = out
    
    return out
}
