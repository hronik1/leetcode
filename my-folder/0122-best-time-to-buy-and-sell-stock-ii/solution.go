func maxProfit(prices []int) int {
    totalProfit := 0
    isLong := false
    boughtPrice := 0
    
    for i, price := range prices {
        if i < len(prices) - 1 {
            if isLong {
                if price > prices[i+1] {
                    totalProfit += (price - boughtPrice)
                    isLong = false
                }
            } else {
                if price < prices[i+1] {
                    boughtPrice = price
                    isLong = true
                }
            }
        } else {
            if isLong {
                totalProfit += (price - boughtPrice)
                isLong = false
            }
        }
    }
    
    return totalProfit
}
