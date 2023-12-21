func tribonacci(n int) int {
    memo := make([]int, n+1, n+1)
    return nth(n, memo)
}

func nth(n int, memo []int) int {
    if n == 0 {
        return 0
    } else if n < 3 {
        return 1
    }

    if memo[n] > 0 {
        return memo[n]
    }

    sum := nth(n-3, memo) + nth(n-2, memo) + nth(n-1, memo)
    memo[n] = sum
    
    return sum
}
