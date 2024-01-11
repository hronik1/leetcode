func climbStairs(n int) int {
    memo := map[int]int {
        1: 1,
        2: 2,
    }

    return climbStairsMemo(n, memo)
}

func climbStairsMemo(n int, memo map[int]int) int {
    if v, ok := memo[n]; ok {
        return v
    }

    v := climbStairsMemo(n-2, memo) + climbStairsMemo(n-1, memo)
    memo[n] = v
    
    return v
}
