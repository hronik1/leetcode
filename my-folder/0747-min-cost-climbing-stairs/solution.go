func minCostClimbingStairs(cost []int) int {
    memo := make([]int, len(cost)+1)
    memo[0] = 0
    memo[1] = 0
    for i := 2; i <= len(cost); i++ {
        memo[i] = -1
    }
    
    return helper(cost, memo, len(cost))
}

func helper(cost []int, memo []int, i int) int {
    if memo[i] != -1 {
        return memo[i]
    }
    
    helper(cost, memo, i-2)
    helper(cost, memo, i-1)
    
    cost2 := memo[i-2] + cost[i-2]
    cost1 := memo[i-1] + cost[i-1]
    if cost2 < cost1 {
        memo[i] = cost2
    } else {
        memo[i] = cost1
    }
    
    return memo[i]
    
}
