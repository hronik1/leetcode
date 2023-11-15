func maxProfit(prices []int) int {
    profit := 0
    maxAfter := prices[len(prices) - 1]
    for i := len(prices) - 2; i >= 0; i -= 1 {
        curr := maxAfter - prices[i]
        if curr > profit {
            profit = curr
        }

        if prices[i] > maxAfter {
            maxAfter = prices[i]
        }
    }

    return profit
}
