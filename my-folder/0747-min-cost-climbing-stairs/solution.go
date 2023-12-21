func minCostClimbingStairs(cost []int) int {
    memo := make([]int, len(cost), len(cost))
    min := costClimbingStairs(cost, memo, 0)
    if skipped := costClimbingStairs(cost, memo, 1); skipped < min {
        min = skipped
    } 

    return min
}

func costClimbingStairs(cost []int, memo []int, fromStair int) int {
    if memo[fromStair] > 0 {
        return memo[fromStair]
    }

    costFromStair := cost[fromStair]
    if fromStair >= len(cost)-2 {
        memo[fromStair] = costFromStair
        return costFromStair
    }

    min := costClimbingStairs(cost, memo, fromStair+1)
    if skipped := costClimbingStairs(cost, memo, fromStair+2); skipped < min {
        min = skipped
    }

    costFromStair += min
    memo[fromStair] = costFromStair

    return costFromStair
}
