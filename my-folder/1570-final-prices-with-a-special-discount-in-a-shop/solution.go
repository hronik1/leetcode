func finalPrices(prices []int) []int {
    l := len(prices)
    out := []int{}
  
    for i := 0; i < l; i++ {
        out = append(out, prices[i])
        for j := i+1; j < l; j++ {
            if prices[j] <= prices[i] {
                out[i] -= prices[j]
                break
            }
        }
    }
    
    return out
}
