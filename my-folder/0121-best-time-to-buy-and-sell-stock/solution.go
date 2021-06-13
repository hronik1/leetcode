func maxProfit(prices []int) int {
    minPricesBefore := []int{prices[0]}
    for i := 1; i < len(prices); i++ {
        m := minPricesBefore[i-1]
        if prices[i] < minPricesBefore[i-1] {
            m = prices[i]
        }
        
        minPricesBefore = append(minPricesBefore, m)
    }
    
    best := 0
    for i := 0; i < len(prices); i++ {
        cur := prices[i] - minPricesBefore[i]
        if cur > best {
            best = cur
        }
    }
    
    return best
}
